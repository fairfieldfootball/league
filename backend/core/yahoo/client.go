package yahoo

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/fairfieldfootball/league/backend/api"
	"github.com/fairfieldfootball/league/backend/common"
)

type Client interface {
	Do(url string, opts YahooRequestOptions) (*http.Response, error)
	Auth(grantType, code string) (Auth, error)
}

func NewClient(config common.Config) Client {
	return &client{
		base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{
			config.Yahoo.ClientID,
			config.Yahoo.ClientSecret,
		}, ":"))),
		new(http.Client),
	}
}

type client struct {
	basicAuth  string
	httpClient *http.Client
}

const (
	TokensActionGet     = "authorization_code"
	TokensActionRefresh = "refresh_token"
)

func (c *client) Auth(grantType, code string) (Auth, error) {
	data := url.Values{}
	data.Set("grant_type", grantType)
	data.Set("redirect_uri", "oob")
	if grantType == TokensActionGet {
		data.Set("code", code)
	} else {
		data.Set("refresh_token", code)
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.login.yahoo.com/oauth2/get_token", strings.NewReader(data.Encode()))
	if err != nil {
		return Auth{}, err
	}

	req.Header.Add(api.HeaderAuthorization, "Basic "+c.basicAuth)
	req.Header.Add(api.HeaderContentType, api.ContentTypeFormURLEncoded)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Auth{}, err
	}

	if res.StatusCode != http.StatusOK {
		return Auth{}, parseYahooError(res, "failed to get yahoo tokens")
	}

	defer res.Body.Close()

	var tokens Auth
	if json.NewDecoder(res.Body).Decode(&tokens); err != nil {
		return Auth{}, err
	}
	return tokens, nil
}

type YahooRequestOptions struct {
	Method      string
	ContentType string
	Body        io.Reader
	AccessToken string
}

func (c *client) Do(url string, opts YahooRequestOptions) (*http.Response, error) {
	method := opts.Method
	if method == "" {
		method = http.MethodGet
	}

	req, err := http.NewRequest(method, url, opts.Body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(api.HeaderAuthorization, "Bearer "+opts.AccessToken)
	if opts.ContentType != "" {
		req.Header.Add(api.HeaderContentType, opts.ContentType)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusUnauthorized {
		return nil, common.NewErr("must authenticate with yahoo", common.ErrCodeInvalidAuth)
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, parseYahooError(res, fmt.Sprintf("failed to %s %s", method, url))
	}

	return res, nil
}

type yahooErrorXml struct {
	Description string `json:"description" xml:"description"`
}

func parseYahooError(res *http.Response, errMsg string) error {
	defer res.Body.Close()

	contentType := res.Header.Get(api.HeaderContentType)

	data := map[string]interface{}{
		"res_status":       res.Status,
		"res_content_type": contentType,
	}

	switch {
	case strings.HasPrefix(contentType, api.ContentTypeXML):
		if body, err := io.ReadAll(res.Body); err != nil {
			data["read_err"] = err.Error()
		} else {
			data["raw_body"] = string(body)

			var out yahooErrorXml
			if err := xml.NewDecoder(res.Body).Decode(&out); err != nil {
				data["decode_err"] = err.Error()
			} else {
				data["cause"] = out
			}
		}
	case strings.HasPrefix(contentType, api.ContentTypeJSON):
		var out map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
			data["decode_err"] = err.Error()
		} else {
			data["cause"] = out
		}
	default:
		if body, err := io.ReadAll(res.Body); err != nil {
			data["read_err"] = err.Error()
		} else {
			data["raw_body"] = string(body)
		}
	}

	return common.NewErr(errMsg, common.ErrCodeServer, common.ErrData(data))
}
