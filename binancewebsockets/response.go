package binancewebsockets

type Response struct {
	RequestID interface{}
	Type      int
	Msg       []byte
	Err       error
}
