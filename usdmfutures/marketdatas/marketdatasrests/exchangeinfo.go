package marketdatasrests

import (
	"context"

	"github.com/Lornzo/gobinance/commons"
)

func NewExchangeInfo(baseURL string) *ExchangeInfo {
	var exchangeInfo *ExchangeInfo = &ExchangeInfo{
		Base: commons.NewBinanceRestfulBase(),
	}
	exchangeInfo.Base.SetBaseURL(baseURL)
	return exchangeInfo
}

type ExchangeInfo struct {
	Base commons.RestfulBase
}

func (e *ExchangeInfo) GetAPIPathes() []string {
	return []string{"fapi", "v1", "exchangeInfo"}
}

func (e *ExchangeInfo) GetWeights() int {
	return 1
}

func (e *ExchangeInfo) DoRequest(ctx context.Context) (ExchangeInfoResponse, error) {

	var (
		resp ExchangeInfoResponse
		err  error
	)

	e.initApiUrl()

	if _, err = e.Base.GET(ctx, &resp); err != nil {
		return ExchangeInfoResponse{}, err
	}

	return resp, nil

}

func (e *ExchangeInfo) initApiUrl() {
	e.Base.SetPathes(e.GetAPIPathes()...)
}
