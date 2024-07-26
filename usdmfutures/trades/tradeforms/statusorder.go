package tradeforms

type StatusOrder struct {
	Symbol              string
	OrderID             int64
	OriginClientOrderID string
	RecvWindow          int64
	Timestamp           int64
}

func (s *StatusOrder) CheckRequired() error {
	return nil
}

func (s *StatusOrder) GetSymbol() string {
	return s.Symbol
}

func (s *StatusOrder) GetOrderID() int64 {
	return s.OrderID
}

func (s *StatusOrder) GetOriginClientOrderID() string {
	return s.OriginClientOrderID
}

func (s *StatusOrder) GetRecvWindow() int64 {
	return s.RecvWindow
}

func (s *StatusOrder) GetTimestamp() int64 {
	return s.Timestamp
}
