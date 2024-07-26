package tradeswebsockets

import "errors"

type modifyOrderFormBuilder struct {
	order ModifyOrderForm
}

func (m modifyOrderFormBuilder) check() error {
	if m.order == nil {
		return errors.New("need is nil")
	}
	return nil
}

func (m modifyOrderFormBuilder) buildMap() (map[string]interface{}, error) {

	var (
		err    error
		dstMap map[string]interface{} = make(map[string]interface{})
	)

	if err = m.check(); err != nil {
		return nil, err
	}

	if orderID := m.order.GetOrderID(); orderID != 0 {
		dstMap["orderId"] = orderID
	}

	if origClientOrderID := m.order.GetOriginClientOrderID(); origClientOrderID != "" {
		dstMap["origClientOrderId"] = origClientOrderID
	}

	if symbol := m.order.GetSymbol(); symbol != "" {
		dstMap["symbol"] = symbol
	}

	if side := m.order.GetSide(); side != "" {
		dstMap["side"] = side
	}

	if quantity := m.order.GetQuantity(); !quantity.IsZero() {
		dstMap["quantity"] = quantity.String()
	}

	if price := m.order.GetPrice(); !price.IsZero() {
		dstMap["price"] = price.String()
	}

	if recvWindow := m.order.GetRecvWindow(); recvWindow != 0 {
		dstMap["recvWindow"] = recvWindow
	}

	if timestamp := m.order.GetTimestamp(); timestamp != 0 {
		dstMap["timestamp"] = timestamp
	}

	return dstMap, nil
}
