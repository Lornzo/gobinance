package accountsrests

import (
	"context"
	"fmt"

	"github.com/Lornzo/gobinance/commons"
)

func NewBalance(baseUrl string, apiKey string, apiSecret string) *Balance {
	var balanceAPI *Balance = &Balance{
		Base: commons.NewBinanceRestfulBase(),
	}
	balanceAPI.Base.SetBaseURL(baseUrl)
	balanceAPI.Base.SetAPIKeyAndSecret(apiKey, apiSecret)
	return balanceAPI
}

type Balance struct {
	Base       commons.RestfulBase
	RecvWindow int64
	Timestamp  int64
}

func (b *Balance) SetRecvWindow(recvWindow int64) {
	b.RecvWindow = recvWindow
}

func (b *Balance) SetTimestamp(ts int64) {
	b.Timestamp = ts
}

func (b *Balance) GetWeights() int {
	return 5
}

func (b *Balance) GetAPIPathes() []string {
	var pathes []string = []string{"fapi", "v2", "balance"}
	return pathes
}

func (b *Balance) DoRequest(ctx context.Context) (BalanceResponse, error) {

	var (
		resp BalanceResponse
		err  error
	)

	b.initApiUrl()

	if _, err = b.Base.GET(ctx, &resp); err != nil {
		return BalanceResponse{}, err
	}

	return resp, nil
}

func (b *Balance) initApiUrl() {

	b.Base.SetPathes(b.GetAPIPathes()...)

	if b.RecvWindow > 0 {
		b.Base.SetQuery("recvWindow", fmt.Sprint(b.RecvWindow))
	}

	if b.Timestamp > 0 {
		b.Base.SetQuery("timestamp", fmt.Sprint(b.Timestamp))
	}

	b.Base.UseSignature(true)

}
