package binancewebsockets

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/Lornzo/gobinance/threadsafetypes"
	"github.com/gorilla/websocket"
)

type websocketWithSubscribers struct {
	*websocket.Conn
	subscribers subscribers
	isRuning    threadsafetypes.Bool
	isDebug     threadsafetypes.Bool
}

func (w *websocketWithSubscribers) Debug() {
	w.isDebug.Set(true)
	log.Println("websocket debug mode on")
}

func (w *websocketWithSubscribers) Subscribe(subscriber Subscriber) error {

	var err error = w.subscribers.add(subscriber)

	if err != nil {
		return err
	}

	return nil

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

	if w.isRuning.Get() {
		return
	}

	w.isRuning.Set(true)
	defer w.isRuning.Set(false)

	for {
		msgType, msgBytes, msgErr := w.ReadMessage(ctx)

		if msgType == -1 {
			msgErr = errors.New("websocket close")
			w.updateHandlers(msgType, msgBytes, msgErr, handlers...)
			break
		}

		go w.updateHandlers(msgType, msgBytes, msgErr, handlers...)
	}
}

func (w *websocketWithSubscribers) RunNewThread(ctx context.Context, handlers ...MessageHander) {

	if w.isRuning.Get() {
		return
	}

	var channel chan bool = make(chan bool)

	go func() {

		var once sync.Once

		for {

			once.Do(func() {
				channel <- true
			})

			msgType, msgBytes, msgErr := w.ReadMessage(ctx)

			if msgType == -1 {
				msgErr = errors.New("websocket close")
				w.updateHandlers(msgType, msgBytes, msgErr, handlers...)
				break
			}

			go w.updateHandlers(msgType, msgBytes, msgErr, handlers...)
		}

	}()

	w.isRuning.Set(<-channel)

}

func (w *websocketWithSubscribers) WriteJSON(v interface{}) error {

	if !w.isRuning.Get() {
		return errors.New("websocket is not running")
	}

	return w.Conn.WriteJSON(v)
}

func (w *websocketWithSubscribers) updateHandlers(msgType int, msg []byte, err error, handlers ...MessageHander) {
	for _, handler := range handlers {
		handler(msgType, msg, err)
	}
}
