package channels

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type WebsocketMsgChannel struct {
	channels    map[string]chan websocketMsgData
	lock        sync.RWMutex
	counter     int
	counterLock sync.Mutex
}

func (w *WebsocketMsgChannel) GetID() string {
	return uuid.NewString()
}

func (w *WebsocketMsgChannel) GetIntID() int {
	w.counterLock.Lock()
	defer w.counterLock.Unlock()
	w.counter++
	return w.counter
}

func (w *WebsocketMsgChannel) Close() error {

	w.lock.Lock()
	defer w.lock.Unlock()

	if w.channels == nil {
		return nil
	}

	for channelName, channel := range w.channels {
		close(channel)
		delete(w.channels, channelName)
	}

	return nil
}

func (w *WebsocketMsgChannel) CreateChannel(channelID interface{}) error {

	w.lock.Lock()

	defer w.lock.Unlock()

	if w.channels == nil {
		w.channels = make(map[string]chan websocketMsgData)
	}

	var id string = fmt.Sprint(channelID)

	if _, exist := w.channels[id]; exist {
		return fmt.Errorf("channel already exist: %s", id)
	}

	w.channels[id] = make(chan websocketMsgData)

	return nil
}

func (w *WebsocketMsgChannel) CloseChannel(channelID interface{}) error {

	w.lock.Lock()

	defer w.lock.Unlock()

	var id string = fmt.Sprint(channelID)

	if w.channels == nil {
		return fmt.Errorf("channel not exist: %s", id)
	}

	if channel, exist := w.channels[id]; exist {
		close(channel)
		delete(w.channels, id)
		return nil
	}

	return fmt.Errorf("channel not exist: %s", id)

}

func (w *WebsocketMsgChannel) getChannel(channelID interface{}) (chan websocketMsgData, error) {

	w.lock.RLock()

	defer w.lock.RUnlock()

	var id string = fmt.Sprint(channelID)

	if w.channels == nil {
		return nil, fmt.Errorf("channel not exist: %s", id)
	}

	if channel, exist := w.channels[id]; exist {
		return channel, nil
	}

	return nil, fmt.Errorf("channel not exist: %s", id)

}

func (w *WebsocketMsgChannel) getChannelAll() []chan websocketMsgData {

	w.lock.RLock()
	defer w.lock.RUnlock()

	var channels []chan websocketMsgData

	if w.channels == nil {
		return channels
	}

	for _, channel := range w.channels {
		channels = append(channels, channel)
	}

	return channels

}

func (w *WebsocketMsgChannel) WriteChannelAll(msgType int, msg []byte, msgErr error) {

	var channels []chan websocketMsgData = w.getChannelAll()

	for _, channel := range channels {
		go func(dstChannel chan websocketMsgData) {
			dstChannel <- websocketMsgData{
				msgType: msgType,
				msg:     msg,
				err:     msgErr,
			}
		}(channel)
	}
}

func (w *WebsocketMsgChannel) WriteChannel(channelID interface{}, msgType int, msg []byte, msgErr error) error {

	var (
		channel chan websocketMsgData
		err     error
	)

	if channel, err = w.getChannel(channelID); err != nil {
		return err
	}

	channel <- websocketMsgData{
		msgType: msgType,
		msg:     msg,
		err:     msgErr,
	}

	return nil
}

func (w *WebsocketMsgChannel) ReadChannel(channelID interface{}) (int, []byte, error) {

	var (
		channel chan websocketMsgData
		err     error
		data    websocketMsgData
	)

	if channel, err = w.getChannel(channelID); err != nil {
		return 0, nil, err
	}

	data = <-channel

	return data.msgType, data.msg, data.err
}

func (w *WebsocketMsgChannel) ReadChannelJSON(channelID interface{}, data interface{}) (int, error) {

	var (
		msgType int
		msg     []byte
		err     error
	)

	if msgType, msg, err = w.ReadChannel(channelID); err != nil {
		return msgType, err
	}

	if err = json.Unmarshal(msg, data); err != nil {
		return msgType, err
	}

	return msgType, nil

}
