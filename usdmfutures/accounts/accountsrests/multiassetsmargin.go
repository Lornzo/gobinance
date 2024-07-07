package accountsrests

import (
	"context"
	"fmt"

	"github.com/Lornzo/gobinance/commons"
)

func NewMultiAssetsMargin(baseUrl string, apiKey string, apiSecret string) *MultiAssetsMargin {
	var multiAssetsMarginAPI *MultiAssetsMargin = &MultiAssetsMargin{
		Base: commons.NewBinanceRestfulBase(),
	}
	multiAssetsMarginAPI.Base.SetBaseURL(baseUrl)
	multiAssetsMarginAPI.Base.SetAPIKeyAndSecret(apiKey, apiSecret)
	return multiAssetsMarginAPI
}

// 查询用户目前在 所有symbol 合约上的联合保证金模式。
type MultiAssetsMargin struct {
	Base       commons.RestfulBase
	RecvWindow int64
	Timestamp  int64
}

func (m *MultiAssetsMargin) GetAPIPathes() []string {
	return []string{"fapi", "v1", "multiAssetsMargin"}
}

func (m *MultiAssetsMargin) GetWeights() int {
	return 30
}

func (m *MultiAssetsMargin) SetRecvWindow(recvWindow int64) {
	m.RecvWindow = recvWindow
}

func (m *MultiAssetsMargin) SetTimestamp(ts int64) {
	m.Timestamp = ts
}

func (m *MultiAssetsMargin) DoRequest(ctx context.Context) (MultiAssetsMarginResponse, error) {
	var (
		resp MultiAssetsMarginResponse
		err  error
	)

	m.initApiUrl()

	if _, err = m.Base.GET(ctx, &resp); err != nil {
		return MultiAssetsMarginResponse{}, err
	}

	return resp, nil
}

func (m *MultiAssetsMargin) initApiUrl() {
	m.Base.SetPathes(m.GetAPIPathes()...)

	if m.RecvWindow > 0 {
		m.Base.SetQuery("recvWindow", fmt.Sprint(m.RecvWindow))
	}

	if m.Timestamp > 0 {
		m.Base.SetQuery("timestamp", fmt.Sprint(m.Timestamp))
	}

	m.Base.UseSignature(true)

}
