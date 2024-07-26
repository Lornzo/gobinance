package tradeswebsockets

import "github.com/shopspring/decimal"

type PositionV2 struct {
	Symbol                 string          `json:"symbol"`
	PositionSide           string          `json:"positionSide"` // 持仓方向
	PositionAmt            decimal.Decimal `json:"positionAmt"`  // 持仓数量
	EntryPrice             decimal.Decimal `json:"entryPrice"`
	BreakEventPrice        decimal.Decimal `json:"breakEvenPrice"`
	MarkPrice              decimal.Decimal `json:"markPrice"`
	UnRealizedProfit       decimal.Decimal `json:"unRealizedProfit"` // 持仓未实现盈亏
	LiquidationPrice       decimal.Decimal `json:"liquidationPrice"`
	IsolatedMargin         decimal.Decimal `json:"isolatedMargin"`
	Notional               decimal.Decimal `json:"notional"`
	MarginAsset            string          `json:"marginAsset"`
	IsolatedWallet         decimal.Decimal `json:"isolatedWallet"`
	InitialMargin          decimal.Decimal `json:"initialMargin"`          // 初始保证金
	MaintMargin            decimal.Decimal `json:"maintMargin"`            // 维持保证金
	PositionInitialMargin  decimal.Decimal `json:"positionInitialMargin"`  // 仓位初始保证金
	OpenOrderInitialMargin decimal.Decimal `json:"openOrderInitialMargin"` // 订单初始保证金
	ADL                    int             `json:"adl"`
	BidNotional            decimal.Decimal `json:"bidNotional"`
	AskNotional            decimal.Decimal `json:"askNotional"`
	UpdateTimestampMilli   int64           `json:"updateTime"` // 更新时间
}

type PositionsV2 []PositionV2
