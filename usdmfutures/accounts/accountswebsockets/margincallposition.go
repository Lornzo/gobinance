package accountswebsockets

import "github.com/shopspring/decimal"

type MarginCallPosition struct {
	Symbol                  string          `json:"s"`
	PositionSide            string          `json:"ps"` // 持倉方向
	PositionAmount          decimal.Decimal `json:"pa"` // 持倉倉位
	MarginType              string          `json:"mt"` // 持倉模式
	IsolatedMargin          decimal.Decimal `json:"iw"` // 逐倉保證金
	MarkPrice               decimal.Decimal `json:"mp"` // 標記價格
	UnrealizedProfitAndLoss decimal.Decimal `json:"up"` // 未实现损益
	MustMargin              decimal.Decimal `json:"mm"` // 持仓需要的维持保证金
}
