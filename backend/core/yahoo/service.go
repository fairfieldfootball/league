package yahoo

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"

	"github.com/fairfieldfootball/league/backend/common"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	client            Client
	refreshTokenStore RefreshTokenStore

	session Session
}

func NewService(client Client, refreshTokenStore RefreshTokenStore) Service {
	return Service{
		client:            client,
		refreshTokenStore: refreshTokenStore,
	}
}

func (s Service) WithSession(session Session) Service {
	s.session = session
	return s
}

func (s Service) EndSession() (Session, bool) {
	session := s.session
	s.session = Session{}
	return session, session.AccessTokenExpiry > 0
}

func (s *Service) FFC(ctx context.Context) (FFC, error) {
	res, _, err := s.doRequest(
		ctx,
		fmt.Sprintf(
			"https://fantasysports.yahooapis.com/fantasy/v2/users;use_login=1/games;game_keys=%s/leagues;league_ids=%s;out=settings,standings/teams",
			strings.Join(FFCGameKeys, ","),
			strings.Join(FFCLeagueIDs, ","),
		),
		YahooRequestOptions{AccessToken: s.session.AccessToken},
	)
	if err != nil {
		return FFC{}, err
	}

	defer res.Body.Close()

	var ffc FFC
	if err := xml.NewDecoder(res.Body).Decode(&ffc); err != nil {
		return FFC{}, err
	}
	return ffc, nil
}

func (s *Service) Authenticate(ctx context.Context, userID primitive.ObjectID, code string) (Auth, error) {
	auth, err := s.client.Auth(TokensActionGet, code)
	if err != nil {
		return Auth{}, err
	}

	if err := s.refreshTokenStore.Upsert(ctx, userID, auth.RefreshToken); err != nil {
		return Auth{}, err
	}

	s.session = Session{
		AccessToken:       auth.AccessToken,
		AccessTokenExpiry: auth.ExpiresIn,
		UserID:            userID,
	}

	return auth, nil
}

func (s *Service) ClearSession(ctx context.Context, userID primitive.ObjectID) error {
	s.EndSession()
	return s.refreshTokenStore.Remove(ctx, userID)
}

func (s *Service) doRequest(ctx context.Context, url string, opts YahooRequestOptions) (*http.Response, Auth, error) {
	res, err := s.client.Do(url, opts)
	if err == nil {
		return res, Auth{}, nil
	}

	e, ok := err.(common.ErrCodeProvider)
	if !ok || e.Code() != common.ErrCodeInvalidAuth {
		return nil, Auth{}, err
	}

	refreshToken, err := s.refreshTokenStore.FindByUserID(ctx, s.session.UserID)
	if err != nil {
		return nil, Auth{}, err
	}

	auth, err := s.client.Auth(TokensActionRefresh, refreshToken.Code)
	if err != nil {
		return nil, Auth{}, err
	}

	if err := s.refreshTokenStore.Upsert(ctx, s.session.UserID, auth.RefreshToken); err != nil {
		return nil, Auth{}, err
	}

	s.session.AccessToken = auth.AccessToken
	s.session.AccessTokenExpiry = auth.ExpiresIn

	opts.AccessToken = auth.AccessToken
	res, err = s.client.Do(url, opts)
	return res, auth, err
}
