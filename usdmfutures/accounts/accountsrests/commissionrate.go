package accountsrests

import (
	"context"
	"fmt"

	"github.com/Lornzo/gobinance/commons"
)

// 用戶手續費率
type CommissionRate struct {
	Base       commons.RestfulBase
	Symbol     string
	RecvWindow int64
	Timestamp  int64
}

func (c *CommissionRate) SetSymbol(symbol string) {
	c.Symbol = symbol
}

func (c *CommissionRate) SetRecvWindow(recvWindow int64) {
	c.RecvWindow = recvWindow
}

func (c *CommissionRate) SetTimestamp(ts int64) {
	c.Timestamp = ts
}

func (c *CommissionRate) GetWeights() int {
	return 20
}

func (c *CommissionRate) GetAPIPathes() []string {
	var pathes []string = []string{"fapi", "v1", "commissionRate"}
	return pathes
}

func (c *CommissionRate) DoRequest(ctx context.Context) (CommissionRateResponse, error) {

	var (
		resp CommissionRateResponse
		err  error
	)

	c.initApiUrl()

	if _, err = c.Base.GET(ctx, &resp); err != nil {
		return CommissionRateResponse{}, err
	}

	return resp, nil
}

func (c *CommissionRate) initApiUrl() {

	c.Base.SetPathes(c.GetAPIPathes()...)

	if c.Symbol != "" {
		c.Base.SetQuery("symbol", c.Symbol)
	}

	if c.RecvWindow > 0 {
		c.Base.SetQuery("recvWindow", fmt.Sprint(c.RecvWindow))
	}

	if c.Timestamp > 0 {
		c.Base.SetQuery("timestamp", fmt.Sprint(c.Timestamp))
	}

	c.Base.UseSignature(true)
}
