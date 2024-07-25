package binancewebsockets

type Request interface {
	GetMethod() string
	GetParams() interface{}
}
