package gobinanceenums

import "fmt"

type NewOrderRespType string

func (n *NewOrderRespType) GetLabelName() string {
	return "newOrderRespType"
}

func (n *NewOrderRespType) GetValue() string {
	return string(*n)
}

func (n *NewOrderRespType) GetNewOrderRespType() string {
	return n.GetValue()
}

func (n *NewOrderRespType) SetNewOrderRespType(newOrderRespType string) {
	*n = NewOrderRespType(newOrderRespType)
}

func (n *NewOrderRespType) SetNewOrderRespTypeACK() {
	*n = NewOrderRespType(NEW_ORDER_RESP_TYPE_ACK)
}

func (n *NewOrderRespType) SetNewOrderRespTypeRESULT() {
	*n = NewOrderRespType(NEW_ORDER_RESP_TYPE_RESULT)
}

func (n *NewOrderRespType) GetQueryString() string {
	return n.GetLabelName() + "=" + n.GetValue()
}

func (n *NewOrderRespType) CheckENUM() error {

	switch n.GetValue() {
	case NEW_ORDER_RESP_TYPE_ACK, NEW_ORDER_RESP_TYPE_RESULT:
		return nil
	default:
		return fmt.Errorf("invalid %s: %s", n.GetLabelName(), n.GetValue())
	}

}
