package tradeswebsockets

import "errors"

type cancelOrderFormBuilder struct {
	order CancelOrderForm
}

func (c cancelOrderFormBuilder) check() error {

	if c.order == nil {
		return errors.New("CancelOrderForm is nil")
	}

	return nil
}

func (c cancelOrderFormBuilder) buildMap() (map[string]interface{}, error) {

	var (
		err    error
		dstMap map[string]interface{} = make(map[string]interface{})
	)

	if err = c.check(); err != nil {
		return nil, err
	}

	if symbol := c.order.GetSymbol(); symbol != "" {
		dstMap["symbol"] = symbol
	}

	if orderID := c.order.GetOrderID(); orderID != 0 {
		dstMap["orderId"] = orderID
	}

	if originClientOrderID := c.order.GetOriginClientOrderID(); originClientOrderID != "" {
		dstMap["origClientOrderId"] = originClientOrderID
	}

	if recvWindow := c.order.GetRecvWindow(); recvWindow != 0 {
		dstMap["recvWindow"] = recvWindow
	}

	if timestamp := c.order.GetTimestamp(); timestamp != 0 {
		dstMap["timestamp"] = timestamp
	}

	return dstMap, nil
}
