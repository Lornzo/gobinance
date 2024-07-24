package threadsafetypes

import "sync"

type Account struct {
	apiKey    string
	apiSecret string
	lock      sync.RWMutex
}

func (a *Account) Set(apiKey string, apiSecret string) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.apiKey = apiKey
	a.apiSecret = apiSecret
}

func (a *Account) Get() (string, string) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.apiKey, a.apiSecret
}

func (a *Account) GetAPIKey() string {
	apiKey, _ := a.Get()
	return apiKey
}

func (a *Account) GetAPISecret() string {
	_, apiSecret := a.Get()
	return apiSecret
}
