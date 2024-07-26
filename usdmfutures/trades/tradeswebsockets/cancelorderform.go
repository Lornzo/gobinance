package tradeswebsockets

type CancelOrderForm interface {
	CheckRequired() error
	GetSymbol() string
	GetOrderID() int64
	GetOriginClientOrderID() string
	GetRecvWindow() int64
	GetTimestamp() int64
}
