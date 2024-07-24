package accountsrests

type Restful struct {
	BaseURL string
	Account AccountUser
}

func (r Restful) ListenKeyCreate() *ListenKeyCreate {
	return NewListenKeyCreate(r.BaseURL, r.Account.GetAPIKey())
}

func (r Restful) ListenKeyDelete() *ListenKeyDelete {
	return NewListenKeyDelete(r.BaseURL, r.Account.GetAPIKey())
}

func (r Restful) ListenKeyUpdate() *ListenKeyUpdate {
	return NewListenKeyUpdate(r.BaseURL, r.Account.GetAPIKey())
}
