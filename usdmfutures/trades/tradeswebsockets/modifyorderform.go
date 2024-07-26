package tradeswebsockets

import "github.com/shopspring/decimal"

type ModifyOrderForm interface {
	GetOrderID() int64
	GetOriginClientOrderID() string
	GetSymbol() string
	GetSide() string
	GetQuantity() decimal.Decimal
	GetPrice() decimal.Decimal
	GetRecvWindow() int64
	GetTimestamp() int64
}
