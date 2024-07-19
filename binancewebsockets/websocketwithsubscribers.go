package binancewebsockets

import (
	"github.com/gorilla/websocket"
)

type websocketWithSubscribers struct {
	*websocket.Conn
	subscribers subscribers
}

func (w *websocketWithSubscribers) Subscribe(subscriber Subscriber) error {
	return w.subscribers.add(subscriber)
}

func (w *websocketWithSubscribers) UnSubscribe(subscriber Subscriber) error {
	return w.subscribers.remove(subscriber)
}

func (w *websocketWithSubscribers) ReadMessage() (int, []byte, error) {

	var (
		msgType int
		msg     []byte
		err     error
	)

	msgType, msg, err = w.Conn.ReadMessage()

	go w.subscribers.update(msgType, msg, err)

	return msgType, msg, err
}
