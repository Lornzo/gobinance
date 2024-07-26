package tradeforms

import (
	"github.com/Lornzo/gobinance/gobinanceenums"
	"github.com/shopspring/decimal"
)

type PlaceOrder struct {
	gobinanceenums.Side
	gobinanceenums.PositionSide
	gobinanceenums.OrderType
	gobinanceenums.ReduceOnly
	gobinanceenums.ClosePosition
	gobinanceenums.TimeInForce
	gobinanceenums.WorkingType
	gobinanceenums.PriceProtect
	gobinanceenums.NewOrderRespType
	Symbol           string
	Quantity         decimal.Decimal
	Price            decimal.Decimal
	NewClientOrderID string
	StopPrice        decimal.Decimal
	ActivationPrice  decimal.Decimal
	CallbackRate     decimal.Decimal
	RecvWindow       int64
	Timestamp        int64
}

func (p *PlaceOrder) CheckRequired() error {
	return nil
}

func (p *PlaceOrder) GetSymbol() string {
	return p.Symbol
}

func (p *PlaceOrder) GetQuantity() decimal.Decimal {
	return p.Quantity
}

func (p *PlaceOrder) GetPrice() decimal.Decimal {
	return p.Price
}

func (p *PlaceOrder) GetNewClientOrderID() string {
	return p.NewClientOrderID
}

func (p *PlaceOrder) GetStopPrice() decimal.Decimal {
	return p.StopPrice
}

func (p *PlaceOrder) GetActivationPrice() decimal.Decimal {
	return p.ActivationPrice
}

func (p *PlaceOrder) GetCallbackRate() decimal.Decimal {
	return p.CallbackRate
}

func (p *PlaceOrder) GetRecvWindow() int64 {
	return p.RecvWindow
}

func (p *PlaceOrder) GetTimestamp() int64 {
	return p.Timestamp
}
