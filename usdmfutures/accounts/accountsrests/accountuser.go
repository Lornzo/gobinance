package accountsrests

type AccountUser interface {
	GetAPIKey() string
	GetAPISecret() string
}
