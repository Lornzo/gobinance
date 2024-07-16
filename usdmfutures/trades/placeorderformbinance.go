package trades

import (
	"github.com/Lornzo/gobinance/gobinanceenums"
	"github.com/shopspring/decimal"
)

type PlaceOrderFormBinance struct {
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

func (p *PlaceOrderFormBinance) CheckRequired() error {
	return nil
}

func (p *PlaceOrderFormBinance) GetSymbol() string {
	return p.Symbol
}

func (p *PlaceOrderFormBinance) GetQuantity() decimal.Decimal {
	return p.Quantity
}

func (p *PlaceOrderFormBinance) GetPrice() decimal.Decimal {
	return p.Price
}

func (p *PlaceOrderFormBinance) GetNewClientOrderID() string {
	return p.NewClientOrderID
}

func (p *PlaceOrderFormBinance) GetStopPrice() decimal.Decimal {
	return p.StopPrice
}

func (p *PlaceOrderFormBinance) GetActivationPrice() decimal.Decimal {
	return p.ActivationPrice
}

func (p *PlaceOrderFormBinance) GetCallbackRate() decimal.Decimal {
	return p.CallbackRate
}

func (p *PlaceOrderFormBinance) GetRecvWindow() int64 {
	return p.RecvWindow
}

func (p *PlaceOrderFormBinance) GetTimestamp() int64 {
	return p.Timestamp
}
