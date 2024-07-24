package accountsrests

import (
	"context"

	"github.com/Lornzo/gobinance/commons"
)

func NewListenKeyCreate(baseURL string, apiKey string) *ListenKeyCreate {
	var createListenKeyAPI *ListenKeyCreate = &ListenKeyCreate{
		Base: commons.NewBinanceRestfulBase(),
	}
	createListenKeyAPI.Base.SetBaseURL(baseURL)
	createListenKeyAPI.Base.SetAPIKeyAndSecret(apiKey, "")
	return createListenKeyAPI
}

type ListenKeyCreate struct {
	Base commons.RestfulBase
}

func (l *ListenKeyCreate) GetAPIPathes() []string {
	return []string{"fapi", "v1", "listenKey"}
}

func (l *ListenKeyCreate) GetWeights() int {
	return 1
}

func (l *ListenKeyCreate) DoRequest(ctx context.Context) (ListenKeyResponse, error) {

	var (
		resp ListenKeyResponse
		err  error
	)

	l.initApiUrl()

	if _, err = l.Base.POST(ctx, nil, &resp); err != nil {
		return ListenKeyResponse{}, err
	}

	return resp, nil

}

func (l *ListenKeyCreate) initApiUrl() {
	l.Base.SetPathes(l.GetAPIPathes()...)
	l.Base.UseApiKeyHeader(true)
}
