package channels

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type BytesChannel struct {
	channels    map[string]chan bytesChannelData
	lock        sync.RWMutex
	counter     int
	counterLock sync.Mutex
}

func (b *BytesChannel) GetID() string {
	return uuid.NewString()
}

func (b *BytesChannel) GetIntID() int {
	b.counterLock.Lock()
	defer b.counterLock.Unlock()
	b.counter++
	return b.counter
}

func (b *BytesChannel) Close() error {

	b.lock.Lock()
	defer b.lock.Unlock()

	if b.channels == nil {
		return nil
	}

	for channelName, channel := range b.channels {
		close(channel)
		delete(b.channels, channelName)
	}

	return nil
}

func (b *BytesChannel) CreateChannel(channelID interface{}) error {

	b.lock.Lock()

	defer b.lock.Unlock()

	if b.channels == nil {
		b.channels = make(map[string]chan bytesChannelData)
	}

	var id string = fmt.Sprint(channelID)

	if _, exist := b.channels[id]; exist {
		return fmt.Errorf("channel already exist: %s", id)
	}
	b.channels[id] = make(chan bytesChannelData)

	return nil
}

func (b *BytesChannel) CloseChannel(channelID interface{}) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	var id string = fmt.Sprint(channelID)

	if b.channels == nil {
		return fmt.Errorf("channel not exist: %s", id)
	}

	if channel, exist := b.channels[id]; exist {
		close(channel)
		delete(b.channels, id)
		return nil
	}
	return fmt.Errorf("channel not exist: %s", id)

}

func (b *BytesChannel) getChannel(channelID interface{}) (chan bytesChannelData, error) {
	b.lock.RLock()
	defer b.lock.RUnlock()

	var id string = fmt.Sprint(channelID)

	if b.channels == nil {
		return nil, fmt.Errorf("channel not exist: %s", id)
	}

	if channel, exist := b.channels[id]; exist {
		return channel, nil
	}
	return nil, fmt.Errorf("channel not exist: %s", id)
}

func (b *BytesChannel) WriteChannel(channelID interface{}, bytes []byte, bytesErr error) error {

	var (
		channel chan bytesChannelData
		err     error
	)

	if channel, err = b.getChannel(channelID); err != nil {
		return err
	}

	channel <- bytesChannelData{
		bytes: bytes,
		err:   bytesErr,
	}

	return nil
}

func (b *BytesChannel) ReadChannel(channelID interface{}) ([]byte, error) {

	var (
		channel chan bytesChannelData
		err     error
	)

	if channel, err = b.getChannel(channelID); err != nil {
		return nil, err
	}

	data := <-channel

	if err = data.err; err != nil {
		return nil, err
	}

	return data.bytes, nil
}

func (b *BytesChannel) ReadChannelJSON(channelID interface{}, data interface{}) error {

	var (
		bytes []byte
		err   error
	)

	if bytes, err = b.ReadChannel(channelID); err != nil {
		return err
	}

	return json.Unmarshal(bytes, data)

}
