package tradeswebsockets

type StatusOrderForm interface {
	GetSymbol() string
	GetOrderID() int64
	GetOriginClientOrderID() string
	GetRecvWindow() int64
	GetTimestamp() int64
}
