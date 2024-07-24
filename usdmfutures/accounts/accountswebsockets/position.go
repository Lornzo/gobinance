package accountswebsockets

import "github.com/shopspring/decimal"

type Position struct {
	Symbol                          string          `json:"s"`
	PositionAmount                  decimal.Decimal `json:"pa"`  // 倉位
	EntryPrice                      decimal.Decimal `json:"ep"`  // 入倉價格
	BreakEvenPrice                  decimal.Decimal `json:"bep"` // 盈亏平衡价
	CumulativeRealizedProfitAndLoss decimal.Decimal `json:"cr"`  // 累计实现损益
	UnrealizedProfitAndLoss         decimal.Decimal `json:"up"`  // 未实现损益
	MarginType                      string          `json:"mt"`  // 保证金模式
	IsolatedMargin                  decimal.Decimal `json:"iw"`  // 逐仓仓位保证金
	PositionSide                    string          `json:"ps"`  // 持倉方向
}
