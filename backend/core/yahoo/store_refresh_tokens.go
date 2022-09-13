package yahoo

import (
	"context"
	"fmt"

	"github.com/fairfieldfootball/league/backend/common"
	"github.com/fairfieldfootball/league/backend/core/mongodb"
	"github.com/fairfieldfootball/league/backend/core/namespaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RefreshTokenStore interface {
	FindByUserID(ctx context.Context, userID primitive.ObjectID) (RefreshToken, error)
	Upsert(ctx context.Context, userID primitive.ObjectID, code string) error
	Remove(ctx context.Context, userID primitive.ObjectID) error
}

func NewRefreshTokenStore(client *mongo.Client) (RefreshTokenStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), common.TimeoutServerOp)
	defer cancel()

	coll, err := mongodb.NewColl(ctx, client, namespaces.DBAuth, namespaces.CollYahooRefreshTokens, mongodb.Index{
		Unique: true,
		Key: mongodb.NewIndexKey(
			mongodb.IndexField{namespaces.FieldUserID, 1}),
	})
	if err != nil {
		return nil, err
	}

	return &refreshTokenStore{coll}, nil
}

type refreshTokenStore struct {
	coll *mongo.Collection
}

func (s *refreshTokenStore) FindByUserID(ctx context.Context, userID primitive.ObjectID) (RefreshToken, error) {
	var refreshToken RefreshToken
	if err := s.coll.FindOne(ctx, bson.D{{namespaces.FieldUserID, userID}}).Decode(&refreshToken); err != nil {
		if err == mongo.ErrNoDocuments {
			return RefreshToken{}, common.NewErr("cannot find refresh token", common.ErrCodeNotFound)
		}
		return RefreshToken{}, common.WrapErr(fmt.Errorf("failed to find refresh token: %s", err), common.ErrCodeServer)
	}
	return refreshToken, nil
}

func (s *refreshTokenStore) Upsert(ctx context.Context, userID primitive.ObjectID, code string) error {
	if _, err := s.coll.UpdateOne(
		ctx,
		bson.D{{namespaces.FieldUserID, userID}},
		bson.D{{"$set", bson.D{{namespaces.FieldCode, code}}}},
		options.Update().SetUpsert(true),
	); err != nil {
		return common.WrapErr(fmt.Errorf("failed to upsert refresh token: %s", err), common.ErrCodeServer)
	}
	return nil
}

func (s *refreshTokenStore) Remove(ctx context.Context, userID primitive.ObjectID) error {
	if _, err := s.coll.DeleteOne(ctx, bson.D{{namespaces.FieldUserID, userID}}); err != nil {
		return common.WrapErr(fmt.Errorf("failed to delete refresh token: %s", err), common.ErrCodeServer)
	}
	return nil
}
