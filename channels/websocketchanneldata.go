package channels

type websocketMsgData struct {
	msgType int
	msg     []byte
	err     error
}
