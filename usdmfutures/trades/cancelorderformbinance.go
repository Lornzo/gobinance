package trades

type CancelOrderFormBinance struct {
	Symbol              string `json:"symbol"`
	OrderID             int64  `json:"orderId"`
	OriginClientOrderID string `json:"origClientOrderId"`
	RecvWindow          int64  `json:"recvWindow"`
	Timestamp           int64  `json:"timestamp"`
}

func (c *CancelOrderFormBinance) CheckRequired() error {
	return nil
}

func (c *CancelOrderFormBinance) GetSymbol() string {
	return c.Symbol
}

func (c *CancelOrderFormBinance) SetSymbol(symbol string) {
	c.Symbol = symbol
}

func (c *CancelOrderFormBinance) GetOrderID() int64 {
	return c.OrderID
}

func (c *CancelOrderFormBinance) SetOrderID(orderID int64) {
	c.OrderID = orderID
}

func (c *CancelOrderFormBinance) GetOriginClientOrderID() string {
	return c.OriginClientOrderID
}

func (c *CancelOrderFormBinance) SetOriginClientOrderID(originClientOrderID string) {
	c.OriginClientOrderID = originClientOrderID
}

func (c *CancelOrderFormBinance) GetRecvWindow() int64 {
	return c.RecvWindow
}

func (c *CancelOrderFormBinance) SetRecvWindow(recvWindow int64) {
	c.RecvWindow = recvWindow
}

func (c *CancelOrderFormBinance) GetTimestamp() int64 {
	return c.Timestamp
}

func (c *CancelOrderFormBinance) SetTimestamp(timestamp int64) {
	c.Timestamp = timestamp
}
