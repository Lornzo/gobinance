package accountswebsockets

type Account interface {
	GetAPIKey() string
	GetAPISecret() string
}
