package binancewebsockets

type MessageHander func(msgType int, msg []byte, err error)
