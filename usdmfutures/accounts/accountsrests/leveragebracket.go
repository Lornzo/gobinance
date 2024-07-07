package accountsrests

import (
	"context"
	"fmt"

	"github.com/Lornzo/gobinance/commons"
)

func NewLeverageBracket(baseUrl string, apiKey string, apiSecret string) *LeverageBracket {
	var leverageBracketAPI *LeverageBracket = &LeverageBracket{
		Base: commons.NewBinanceRestfulBase(),
	}
	leverageBracketAPI.Base.SetBaseURL(baseUrl)
	leverageBracketAPI.Base.SetAPIKeyAndSecret(apiKey, apiSecret)
	return leverageBracketAPI
}

// 查询账户特定交易对的杠杆分层标准
type LeverageBracket struct {
	Base       commons.RestfulBase
	Symbol     string
	RecvWindow int64
	Timestamp  int64
}

func (l *LeverageBracket) GetAPIPathes() []string {
	return []string{"fapi", "v1", "leverageBracket"}
}

func (l *LeverageBracket) GetWeights() int {
	return 1
}

func (l *LeverageBracket) SetSymbol(symbol string) {
	l.Symbol = symbol
}

func (l *LeverageBracket) SetRecvWindow(recvWindow int64) {
	l.RecvWindow = recvWindow
}

func (l *LeverageBracket) SetTimestamp(ts int64) {
	l.Timestamp = ts
}

func (l *LeverageBracket) DoRequest(ctx context.Context) (LeverageBracketResponse, error) {
	var (
		resp LeverageBracketResponse
		err  error
	)

	l.initApiUrl()

	if _, err = l.Base.GET(ctx, &resp); err != nil {
		return LeverageBracketResponse{}, err
	}

	return resp, nil
}

func (l *LeverageBracket) initApiUrl() {
	l.Base.SetPathes(l.GetAPIPathes()...)

	if l.Symbol != "" {
		l.Base.SetQuery("symbol", l.Symbol)
	}

	if l.RecvWindow > 0 {
		l.Base.SetQuery("recvWindow", fmt.Sprint(l.RecvWindow))
	}

	if l.Timestamp > 0 {
		l.Base.SetQuery("timestamp", fmt.Sprint(l.Timestamp))
	}

	l.Base.UseSignature(true)

}
