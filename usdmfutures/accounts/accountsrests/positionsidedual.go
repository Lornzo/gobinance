package accountsrests

import (
	"context"
	"fmt"

	"github.com/Lornzo/gobinance/commons"
)

func NewPositionSideDual(baseUrl string, apiKey string, apiSecret string) *PositionSideDual {
	var positionSideDualAPI *PositionSideDual = &PositionSideDual{
		Base: commons.NewBinanceRestfulBase(),
	}
	positionSideDualAPI.Base.SetBaseURL(baseUrl)
	positionSideDualAPI.Base.SetAPIKeyAndSecret(apiKey, apiSecret)
	return positionSideDualAPI

}

// 查询用户目前在 所有symbol 合约上的持仓模式：双向持仓或单向持仓。
type PositionSideDual struct {
	Base       commons.RestfulBase
	RecvWindow int64
	Timestamp  int64
}

func (p *PositionSideDual) GetAPIPathes() []string {
	return []string{"fapi", "v1", "positionSide", "dual"}
}

func (p *PositionSideDual) GetWeights() int {
	return 30
}

func (p *PositionSideDual) SetRecvWindow(recvWindow int64) {
	p.RecvWindow = recvWindow
}

func (p *PositionSideDual) SetTimestamp(ts int64) {
	p.Timestamp = ts
}

func (p *PositionSideDual) DoRequest(ctx context.Context) (PositionSideDualResponse, error) {
	var (
		resp PositionSideDualResponse
		err  error
	)

	p.initApiUrl()

	if _, err = p.Base.GET(ctx, &resp); err != nil {
		return PositionSideDualResponse{}, err
	}

	return resp, nil
}

func (p *PositionSideDual) initApiUrl() {
	p.Base.SetPathes(p.GetAPIPathes()...)

	if p.RecvWindow > 0 {
		p.Base.SetQuery("recvWindow", fmt.Sprint(p.RecvWindow))
	}

	if p.Timestamp > 0 {
		p.Base.SetQuery("timestamp", fmt.Sprint(p.Timestamp))
	}

	p.Base.UseSignature(true)
}
