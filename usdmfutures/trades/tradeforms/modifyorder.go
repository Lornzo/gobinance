package tradeforms

import (
	"github.com/Lornzo/gobinance/gobinanceenums"
	"github.com/shopspring/decimal"
)

type ModifyOrder struct {
	gobinanceenums.Side
	OrderID             int64
	OriginClientOrderID string
	Symbol              string
	Quantity            decimal.Decimal
	Price               decimal.Decimal
	RecvWindow          int64
	Timestamp           int64
}

func (p *ModifyOrder) GetOrderID() int64 {
	return p.OrderID
}

func (p *ModifyOrder) GetOriginClientOrderID() string {
	return p.OriginClientOrderID
}

func (p *ModifyOrder) GetSymbol() string {
	return p.Symbol
}

func (p *ModifyOrder) GetQuantity() decimal.Decimal {
	return p.Quantity
}

func (p *ModifyOrder) GetPrice() decimal.Decimal {
	return p.Price
}

func (p *ModifyOrder) GetRecvWindow() int64 {
	return p.RecvWindow
}

func (p *ModifyOrder) GetTimestamp() int64 {
	return p.Timestamp
}
