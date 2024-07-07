package accountspushs

import (
	"context"

	"github.com/Lornzo/gobinance/commons"
)

func NewListenKey(baseURL string, apiKey string) *ListenKey {
	var listenKeyAPI *ListenKey = &ListenKey{
		Base: commons.NewBinanceRestfulBase(),
	}
	listenKeyAPI.Base.SetBaseURL(baseURL)
	listenKeyAPI.Base.SetAPIKeyAndSecret(apiKey, "")
	return listenKeyAPI
}

type ListenKey struct {
	Base commons.RestfulBase
}

func (l *ListenKey) GetAPIPathes() []string {
	return []string{"fapi", "v1", "listenKey"}
}

func (l *ListenKey) GetWeights() int {
	return 1
}

func (l *ListenKey) Create(ctx context.Context) (ListenKeyResponse, error) {
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

func (l *ListenKey) Update(ctx context.Context) (ListenKeyResponse, error) {

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

func (l *ListenKey) Delete(ctx context.Context) error {

	var err error

	l.initApiUrl()

	if _, err = l.Base.DELETE(ctx, nil); err != nil {
		return err
	}

	return nil
}

func (l *ListenKey) initApiUrl() {
	l.Base.SetPathes(l.GetAPIPathes()...)
	l.Base.UseApiKeyHeader(true)
}
