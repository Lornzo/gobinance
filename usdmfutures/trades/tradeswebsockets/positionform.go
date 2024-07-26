package tradeswebsockets

type PositionForm interface {
	GetSymbol() string
	GetRecvWindow() int64
	GetTimestamp() int64
}
