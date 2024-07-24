package accountswebsockets

import "github.com/shopspring/decimal"

type AccountPosition struct {
	EntryPrice       decimal.Decimal `json:"entryPrice"`
	MarginType       string          `json:"marginType"`
	IsAutoAddMargin  bool            `json:"isAutoAddMargin"`
	IsolatedMargin   decimal.Decimal `json:"isolatedMargin"`
	Leverage         int64           `json:"leverage"`
	LiquidationPrice decimal.Decimal `json:"liquidationPrice"`
	MarkPrice        decimal.Decimal `json:"markPrice"`
	MaxQuantity      decimal.Decimal `json:"maxQty"`
	PositionAmount   decimal.Decimal `json:"positionAmt"`
	Symbol           string          `json:"symbol"`
	UnRealizedProfit decimal.Decimal `json:"unRealizedProfit"`
	PositionSide     string          `json:"positionSide"`
}

type AccountPositions []AccountPosition
