package accountsrests

import (
	"context"
	"fmt"

	"github.com/Lornzo/gobinance/commons"
)

func NewRateLimitOrder(baseUrl string, apiKey string, apiSecret string) *RateLimitOrder {
	var rateLimitOrderAPI *RateLimitOrder = &RateLimitOrder{
		Base: commons.NewBinanceRestfulBase(),
	}
	rateLimitOrderAPI.Base.SetBaseURL(baseUrl)
	rateLimitOrderAPI.Base.SetAPIKeyAndSecret(apiKey, apiSecret)
	return rateLimitOrderAPI
}

// 查詢用戶下單限頻
type RateLimitOrder struct {
	Base       commons.RestfulBase
	RecvWindow int64
	Timestamp  int64
}

func (r *RateLimitOrder) GetAPIPathes() []string {
	return []string{"fapi", "v1", "rateLimit", "order"}
}

func (r *RateLimitOrder) GetWeights() int {
	return 1
}

func (r *RateLimitOrder) SetRecvWindow(recvWindow int64) {
	r.RecvWindow = recvWindow
}

func (r *RateLimitOrder) SetTimestamp(ts int64) {
	r.Timestamp = ts
}

func (r *RateLimitOrder) DoRequest(ctx context.Context) (RateLimitOrderResponse, error) {
	var (
		resp RateLimitOrderResponse
		err  error
	)

	r.initApiUrl()

	if _, err = r.Base.GET(ctx, &resp); err != nil {
		return RateLimitOrderResponse{}, err
	}

	return resp, nil
}

func (r *RateLimitOrder) initApiUrl() {
	r.Base.SetPathes(r.GetAPIPathes()...)

	if r.RecvWindow > 0 {
		r.Base.SetQuery("recvWindow", fmt.Sprint(r.RecvWindow))
	}

	if r.Timestamp > 0 {
		r.Base.SetQuery("timestamp", fmt.Sprint(r.Timestamp))
	}

	r.Base.UseSignature(true)
}
