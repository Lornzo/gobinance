package tradeforms

type CancelOrder struct {
	Symbol              string `json:"symbol"`
	OrderID             int64  `json:"orderId"`
	OriginClientOrderID string `json:"origClientOrderId"`
	RecvWindow          int64  `json:"recvWindow"`
	Timestamp           int64  `json:"timestamp"`
}

func (c *CancelOrder) CheckRequired() error {
	return nil
}

func (c *CancelOrder) GetSymbol() string {
	return c.Symbol
}

func (c *CancelOrder) SetSymbol(symbol string) {
	c.Symbol = symbol
}

func (c *CancelOrder) GetOrderID() int64 {
	return c.OrderID
}

func (c *CancelOrder) SetOrderID(orderID int64) {
	c.OrderID = orderID
}

func (c *CancelOrder) GetOriginClientOrderID() string {
	return c.OriginClientOrderID
}

func (c *CancelOrder) SetOriginClientOrderID(originClientOrderID string) {
	c.OriginClientOrderID = originClientOrderID
}

func (c *CancelOrder) GetRecvWindow() int64 {
	return c.RecvWindow
}

func (c *CancelOrder) SetRecvWindow(recvWindow int64) {
	c.RecvWindow = recvWindow
}

func (c *CancelOrder) GetTimestamp() int64 {
	return c.Timestamp
}

func (c *CancelOrder) SetTimestamp(timestamp int64) {
	c.Timestamp = timestamp
}
