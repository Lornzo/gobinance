package tradeswebsockets

import "github.com/shopspring/decimal"

type Position struct {
	Symbol               string          `json:"symbol"`
	PositionAmt          decimal.Decimal `json:"positionAmt"`
	EntryPrice           decimal.Decimal `json:"entryPrice"`
	BreakEventPrice      decimal.Decimal `json:"breakEvenPrice"`
	MarkPrice            decimal.Decimal `json:"markPrice"`
	UnRealizedProfit     decimal.Decimal `json:"unRealizedProfit"`
	LiquidationPrice     decimal.Decimal `json:"liquidationPrice"`
	Leverage             decimal.Decimal `json:"leverage"`
	MaxNotionalValue     decimal.Decimal `json:"maxNotionalValue"`
	MarginType           string          `json:"marginType"`
	IsolatedMargin       decimal.Decimal `json:"isolatedMargin"`
	IsAutoAddMargin      string          `json:"isAutoAddMargin"`
	PositionSide         string          `json:"positionSide"`
	Notional             decimal.Decimal `json:"notional"`
	IsolatedWallet       decimal.Decimal `json:"isolatedWallet"`
	UpdateTimestampMilli int64           `json:"updateTime"`
}

type Positions []Position
