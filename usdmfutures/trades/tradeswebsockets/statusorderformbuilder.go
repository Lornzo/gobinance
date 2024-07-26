package tradeswebsockets

type statusOrderFormBuilder struct {
	order StatusOrderForm
}

func (s statusOrderFormBuilder) buildMap() (map[string]interface{}, error) {
	var (
		dstMap map[string]interface{} = make(map[string]interface{})
	)

	if symbol := s.order.GetSymbol(); symbol != "" {
		dstMap["symbol"] = symbol
	}

	if orderID := s.order.GetOrderID(); orderID != 0 {
		dstMap["orderId"] = orderID
	}

	if origClientOrderID := s.order.GetOriginClientOrderID(); origClientOrderID != "" {
		dstMap["origClientOrderId"] = origClientOrderID
	}

	if recvWindow := s.order.GetRecvWindow(); recvWindow != 0 {
		dstMap["recvWindow"] = recvWindow
	}

	if timestamp := s.order.GetTimestamp(); timestamp != 0 {
		dstMap["timestamp"] = timestamp
	}

	return dstMap, nil
}
