package binancewebsockets

import (
	"context"
	"errors"

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

func (w *websocketWithSubscribers) ReadMessage(ctx context.Context) (int, []byte, error) {

	type message struct {
		msgType int
		msg     []byte
		err     error
	}

	var (
		msgChan  chan message = make(chan message)
		msgType  int
		msgBytes []byte
		msgErr   error
	)

	go func() {
		var msg message
		msg.msgType, msg.msg, msg.err = w.Conn.ReadMessage()
		msgChan <- msg
	}()

	select {
	case <-ctx.Done():
		msgType = -2
		msgErr = errors.New("read message timeout")
	case msg := <-msgChan:
		msgType = msg.msgType
		msgBytes = msg.msg
		msgErr = msg.err
	}

	go w.subscribers.update(msgType, msgBytes, msgErr)

	return msgType, msgBytes, msgErr
}

func (w *websocketWithSubscribers) Run(ctx context.Context, handlers ...MessageHander) {
	for {
		msgType, msgBytes, msgErr := w.ReadMessage(ctx)

		if msgType == -1 {
			w.updateHandlers(msgType, msgBytes, msgErr, handlers...)
			break
		}

		go w.updateHandlers(msgType, msgBytes, msgErr, handlers...)
	}
}

func (w *websocketWithSubscribers) updateHandlers(msgType int, msg []byte, err error, handlers ...MessageHander) {
	for _, handler := range handlers {
		handler(msgType, msg, err)
	}
}
