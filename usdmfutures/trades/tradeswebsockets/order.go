package tradeswebsockets

import "github.com/shopspring/decimal"

type Order struct {
	OrderID                 int64           `json:"orderId,omitempty"`
	Symbol                  string          `json:"symbol"`
	Status                  string          `json:"status"`
	ClientOrderID           string          `json:"clientOrderId,omitempty"`
	Price                   decimal.Decimal `json:"price"`
	AveragePrice            decimal.Decimal `json:"avgPrice,omitempty"`
	OriginQuantity          decimal.Decimal `json:"origQty"`
	ExecutedQuantity        decimal.Decimal `json:"executedQty"`
	CumulativeQuantity      decimal.Decimal `json:"cumQty"`
	CumulativeQuote         decimal.Decimal `json:"cumQuote"`
	TimeInForce             string          `json:"timeInForce"`
	Type                    string          `json:"type"`
	ReduceOnly              bool            `json:"reduceOnly"`
	ClosePosition           bool            `json:"closePosition"`
	Side                    string          `json:"side"`
	PositionSide            string          `json:"positionSide"`
	StopPrice               decimal.Decimal `json:"stopPrice"`
	WorkingType             string          `json:"workingType"`
	PriceProtect            bool            `json:"priceProtect"`
	OriginType              string          `json:"origType"`
	PriceMatch              string          `json:"priceMatch"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	GoodTillDateMilli       int64           `json:"goodTillDate"`
	UpdateTimeMilli         int64           `json:"updateTime"`
}
