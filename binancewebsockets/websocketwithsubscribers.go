package binancewebsockets

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/Lornzo/gobinance/channels"
	"github.com/Lornzo/gobinance/internal/loggers"
	"github.com/Lornzo/gobinance/threadsafetypes"
	"github.com/gorilla/websocket"
)

type websocketWithSubscribers struct {
	logger loggers.Logger
	*websocket.Conn
	channel     channels.WebsocketMsgChannel
	subscribers subscribers
	isRuning    threadsafetypes.Bool
}

func (w *websocketWithSubscribers) GetName() string {
	return "websocketWithSubscribers"
}

func (w *websocketWithSubscribers) SetLogger(logger Logger) {
	w.logger.SetConcreteLogger(logger)
}

func (w *websocketWithSubscribers) Subscribe(subscriber Subscriber) error {

	var (
		action string = "Subscribe"
		err    error
	)

	w.logger.Debuf("%s-%s: subscriber id:%s", w.GetName(), action, subscriber.GetID())

	if err = w.subscribers.add(subscriber); err != nil {
		w.logger.Errorf("%s-%s err1: add subscriber %s error : %w", w.GetName(), action, subscriber.GetID(), err)
		return err
	}

	return nil
}

func (w *websocketWithSubscribers) UnSubscribe(subscriber Subscriber) error {

	var (
		err    error
		action string = "UnSubscribe"
	)

	w.logger.Debuf("%s-%s: subscriber id:%s", w.GetName(), action, subscriber.GetID())

	if err = w.subscribers.remove(subscriber); err != nil {
		w.logger.Errorf("%s-%s err1: remove subscriber %s error : %w", w.GetName(), action, subscriber.GetID(), err)
	}

	return nil
}

func (w *websocketWithSubscribers) MakeRequestByIntIndex(ctx context.Context, request Request) (Response, error) {

	var (
		action    string = "MakeRequestByIntIndex"
		channelID int    = w.channel.GetIntID()
	)
	w.logger.Debuf("%s-%s: channelID:%d", w.GetName(), action, channelID)

	return w.MakeRequest(ctx, channelID, request)
}

func (w *websocketWithSubscribers) MakeRequestByUUIDIndex(ctx context.Context, request Request) (Response, error) {

	var (
		action    string = "MakeRequestByUUIDIndex"
		channelID string = w.channel.GetID()
	)

	w.logger.Debuf("%s-%s: channelID:%s", w.GetName(), action, channelID)

	return w.MakeRequest(ctx, channelID, request)
}

func (w *websocketWithSubscribers) MakeRequest(ctx context.Context, requestID interface{}, request Request) (Response, error) {

	var (
		err error
		req struct {
			ID     interface{} `json:"id"`
			Method string      `json:"method"`
			Params interface{} `json:"params"`
		}
		resp chan Response = make(chan Response)
	)

	defer close(resp)

	req.ID = requestID
	req.Method = request.GetMethod()
	req.Params = request.GetParams()

	if err = w.channel.CreateChannel(requestID); err != nil {
		return Response{}, err
	}

	defer w.channel.CloseChannel(requestID)

	go func() {

		msgType, msg, err := w.channel.ReadChannel(requestID)
		resp <- Response{
			RequestID: requestID,
			Type:      msgType,
			Msg:       msg,
			Err:       err,
		}

	}()

	if err = w.WriteJSON(req); err != nil {
		return Response{}, err
	}

	select {
	case <-ctx.Done():
		return Response{}, errors.New("request timeout")
	case data := <-resp:
		return data, nil
	}

}

func (w *websocketWithSubscribers) ReadMessage(ctx context.Context) (int, []byte, error) {

	type message struct {
		msgType int
		msg     []byte
		err     error
	}

	var (
		action   string       = "ReadMessage"
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

	if msgType == -1 {
		msgErr = errors.New("websocket close")
	}

	if msgErr != nil {
		w.logger.Errorf("%s-%s websocket error: msgType:%d, msg:%s, msgErr:%w", w.GetName(), action, msgType, string(msgBytes), msgErr)
	}

	go w.subscribers.update(msgType, msgBytes, msgErr)
	go w.idHandler(msgType, msgBytes, msgErr)

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

		go w.updateHandlers(msgType, msgBytes, msgErr, handlers...)

		if msgType == -1 {
			break
		}

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

func (w *websocketWithSubscribers) idHandler(msgType int, msg []byte, msgErr error) {

	if msgType == -1 {
		w.channel.WriteChannelAll(msgType, msg, msgErr)
	}

	var (
		channelID    interface{}
		channelExist bool
		msgMap       map[string]interface{}
		err          error
	)

	if err = json.Unmarshal(msg, &msgMap); err != nil {
		return
	}

	if channelID, channelExist = msgMap["id"]; !channelExist {
		return
	}

	w.channel.WriteChannel(channelID, msgType, msg, msgErr)

}
