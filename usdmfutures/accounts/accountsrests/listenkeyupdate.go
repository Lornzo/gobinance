package accountsrests

import (
	"context"

	"github.com/Lornzo/gobinance/commons"
)

func NewListenKeyUpdate(baseURL string, apiKey string) *ListenKeyUpdate {
	var dstAPI *ListenKeyUpdate = &ListenKeyUpdate{
		Base: commons.NewBinanceRestfulBase(),
	}
	dstAPI.Base.SetBaseURL(baseURL)
	dstAPI.Base.SetAPIKeyAndSecret(apiKey, "")
	return dstAPI
}

type ListenKeyUpdate struct {
	Base commons.RestfulBase
}

func (l *ListenKeyUpdate) GetAPIPathes() []string {
	return []string{"fapi", "v1", "listenKey"}
}

func (l *ListenKeyUpdate) GetWeights() int {
	return 1
}

func (l *ListenKeyUpdate) DoRequest(ctx context.Context) (ListenKeyResponse, error) {

	var (
		resp ListenKeyResponse
		err  error
	)

	l.initApiUrl()

	if _, err = l.Base.PUT(ctx, nil, &resp); err != nil {
		return ListenKeyResponse{}, err
	}

	return resp, nil

}

func (l *ListenKeyUpdate) initApiUrl() {
	l.Base.SetPathes(l.GetAPIPathes()...)
	l.Base.UseApiKeyHeader(true)
}
