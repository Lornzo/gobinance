package tradeforms

type Position struct {
	Symbol     string
	RecvWindow int64
	Timestamp  int64
}

func (p *Position) GetSymbol() string {
	return p.Symbol
}

func (p *Position) GetRecvWindow() int64 {
	return p.RecvWindow
}

func (p *Position) GetTimestamp() int64 {
	return p.Timestamp
}
