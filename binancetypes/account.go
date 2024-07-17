package binancetypes

type Account struct {
	APIKey    string
	APISecret string
}

func (a Account) GetAPIKey() string {
	return a.APIKey
}

func (a Account) GetAPISecret() string {
	return a.APISecret
}
