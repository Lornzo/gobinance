package gobinanceenums

import "fmt"

type OrderType string

func (o *OrderType) GetLabelName() string {
	return "type"
}

func (o *OrderType) GetValue() string {
	return string(*o)
}

func (o *OrderType) GetType() string {
	return o.GetValue()
}

func (o *OrderType) SetType(orderType string) {
	*o = OrderType(orderType)
}

func (o *OrderType) SetTypeLIMIT() {
	*o = OrderType(ORDER_TYPE_LIMIT)
}

func (o *OrderType) SetTypeMARKET() {
	*o = OrderType(ORDER_TYPE_MARKET)
}

func (o *OrderType) SetTypeSTOP() {
	*o = OrderType(ORDER_TYPE_STOP)
}

func (o *OrderType) SetTypeSTOPMARKET() {
	*o = OrderType(ORDER_TYPE_STOP_MARKET)
}

func (o *OrderType) SetTypeTAKEPROFIT() {
	*o = OrderType(ORDER_TYPE_TAKE_PROFIT)
}

func (o *OrderType) SetTypeTAKEPROFITMARKET() {
	*o = OrderType(ORDER_TYPE_TAKE_PROFIT_MARKET)
}

func (o *OrderType) SetTypeTRAILINGSTOPMARKET() {
	*o = OrderType(ORDER_TYPE_TRAILING_STOP_MARKET)
}

func (o *OrderType) GetQueryString() string {
	return fmt.Sprint(o.GetLabelName(), "=", o.GetValue())
}

func (o *OrderType) CheckENUM() error {

	switch o.GetValue() {
	case ORDER_TYPE_LIMIT, ORDER_TYPE_MARKET, ORDER_TYPE_STOP, ORDER_TYPE_STOP_MARKET, ORDER_TYPE_TAKE_PROFIT, ORDER_TYPE_TAKE_PROFIT_MARKET, ORDER_TYPE_TRAILING_STOP_MARKET:
		return nil
	default:
		return fmt.Errorf("invalid %s: %s", o.GetLabelName(), o.GetValue())
	}

}
