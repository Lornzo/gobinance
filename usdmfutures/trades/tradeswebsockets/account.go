package tradeswebsockets

type Account interface {
	GetAPIKey() string
	GetAPISecret() string
}
