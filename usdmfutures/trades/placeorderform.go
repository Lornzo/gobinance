package trades

import "github.com/shopspring/decimal"

type PlaceOrderForm interface {
	CheckRequired() error
	GetSymbol() string
	GetSide() string
	GetPositionSide() string
	GetType() string
	GetReduceOnly() string
	GetQuantity() decimal.Decimal
	GetPrice() decimal.Decimal
	GetNewClientOrderID() string
	GetStopPrice() decimal.Decimal
	GetClosePosition() string
	GetActivationPrice() decimal.Decimal
	GetCallbackRate() decimal.Decimal
	GetTimeInForce() string
	GetWorkingType() string
	GetPriceProtect() string
	GetNewOrderRespType() string
	GetRecvWindow() int64
	GetTimestamp() int64
}
