package binancewebsockets

type Subscriber interface {
	GetID() string
	Update(msgType int, msg []byte, err error)
}
