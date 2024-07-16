package trades

type Account interface {
	GetAPIKey() string
	GetAPISecret() string
	GetSignatureWithQueries(queries ...string) (string, error)
}
