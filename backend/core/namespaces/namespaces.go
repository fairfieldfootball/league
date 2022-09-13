package namespaces

var (
	DBApp    = "ffc_app"
	CollBets = "bets"

	DBAuth                 = "ffc_auth"
	CollRefreshTokens      = "refresh_tokens"
	CollPasswords          = "passwords"
	CollUsers              = "users"
	CollYahooRefreshTokens = "yahoo_refresh_tokens"
)

type Namespace struct {
	Database   *string
	Collection *string
}

var (
	Registry = []Namespace{
		{&DBApp, &CollBets},
		{&DBAuth, &CollRefreshTokens},
		{&DBAuth, &CollPasswords},
		{&DBAuth, &CollUsers},
		{&DBAuth, &CollYahooRefreshTokens},
	}
)

const (
	FieldID       = "_id"
	FieldName     = "name"
	FieldSessions = "sessions"
	FieldUserID   = "user_id"

	FieldConsumed = "consumed"
	FieldSub      = "sub"

	FieldUsername       = "username"
	FieldSalt           = "salt"
	FieldHashedPassword = "hashed_password"

	FieldCode = "code"

	FieldFirstName = "first_name"
	FieldLastName  = "last_name"
)
