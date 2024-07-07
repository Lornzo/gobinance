package accountsrests

import (
	"context"
	"fmt"

	"github.com/Lornzo/gobinance/commons"
)

func NewAccount(baseUrl string, apiKey string, apiSecret string) *Account {
	var balanceAPI *Account = &Account{
		Base: commons.NewBinanceRestfulBase(),
	}
	balanceAPI.Base.SetBaseURL(baseUrl)
	balanceAPI.Base.SetAPIKeyAndSecret(apiKey, apiSecret)
	return balanceAPI
}

type Account struct {
	Base       commons.RestfulBase
	RecvWindow int64
	Timestamp  int64
}

func (a *Account) SetRecvWindow(recvWindow int64) {
	a.RecvWindow = recvWindow
}

func (a *Account) SetTimestamp(ts int64) {
	a.Timestamp = ts
}

func (a *Account) GetAPIPathes() []string {
	var pathes []string = []string{"fapi", "v2", "account"}
	return pathes
}

func (a *Account) GetWeights() int {
	return 5
}

func (a *Account) DoRequest(ctx context.Context) (AccountResponse, error) {

	var (
		resp AccountResponse
		err  error
	)

	a.initApiUrl()

	if _, err = a.Base.GET(ctx, &resp); err != nil {
		return AccountResponse{}, err
	}

	return resp, nil
}

func (a *Account) initApiUrl() {

	a.Base.SetPathes(a.GetAPIPathes()...)

	if a.RecvWindow > 0 {
		a.Base.SetQuery("recvWindow", fmt.Sprint(a.RecvWindow))
	}

	if a.Timestamp > 0 {
		a.Base.SetQuery("timestamp", fmt.Sprint(a.Timestamp))
	}

	a.Base.UseSignature(true)

}
