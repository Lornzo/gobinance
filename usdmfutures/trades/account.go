package trades

type Account interface {
	GetAPIKey() string
	GetAPISecret() string
}
