package yahoo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	XOAuthGuid   string `json:"xoauth_yahoo_guid"`
}

type RefreshToken struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `bson:"user_id"`
	Code   string             `bson:"code"`
}

type Session struct {
	AccessToken       string
	AccessTokenExpiry int
	UserID            primitive.ObjectID
}
