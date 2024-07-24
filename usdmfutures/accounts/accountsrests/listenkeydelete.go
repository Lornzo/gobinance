package accountsrests

import (
	"context"

	"github.com/Lornzo/gobinance/commons"
)

func NewListenKeyDelete(baseURL string, apiKey string) *ListenKeyDelete {
	var dstAPI *ListenKeyDelete = &ListenKeyDelete{
		Base: commons.NewBinanceRestfulBase(),
	}
	dstAPI.Base.SetBaseURL(baseURL)
	dstAPI.Base.SetAPIKeyAndSecret(apiKey, "")
	return dstAPI
}

type ListenKeyDelete struct {
	Base commons.RestfulBase
}

func (l *ListenKeyDelete) GetAPIPathes() []string {
	return []string{"fapi", "v1", "listenKey"}
}

func (l *ListenKeyDelete) GetWeights() int {
	return 1
}

func (l *ListenKeyDelete) DoRequest(ctx context.Context) error {

	var (
		err error
	)

	l.initApiUrl()

	if _, err = l.Base.DELETE(ctx, nil); err != nil {
		return err
	}

	return nil

}

func (l *ListenKeyDelete) initApiUrl() {
	l.Base.SetPathes(l.GetAPIPathes()...)
	l.Base.UseApiKeyHeader(true)
}
