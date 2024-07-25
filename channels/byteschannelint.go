package channels

import (
	"encoding/json"
	"fmt"
	"sync"
)

type BytesChannelInt struct {
	lock        sync.RWMutex
	channels    map[int]chan bytesChannelData
	counterLock sync.Mutex
	counter     int
}

func (b *BytesChannelInt) GetID() int {
	b.counterLock.Lock()
	defer b.counterLock.Unlock()
	b.counter++
	return b.counter
}

func (b *BytesChannelInt) Close() error {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.channels == nil {
		return nil
	}

	for channelIndex, channel := range b.channels {
		close(channel)
		delete(b.channels, channelIndex)
	}

	return nil

}

func (b *BytesChannelInt) CreateChannel(channelID int) error {

	b.lock.Lock()
	defer b.lock.Unlock()

	if b.channels == nil {
		b.channels = make(map[int]chan bytesChannelData)
	}

	if _, exist := b.channels[channelID]; exist {
		return fmt.Errorf("channel already exist: %d", channelID)
	}

	b.channels[channelID] = make(chan bytesChannelData)

	return nil
}

func (b *BytesChannelInt) CloseChannel(channelID int) error {

	b.lock.Lock()
	defer b.lock.Unlock()

	if b.channels == nil {
		return fmt.Errorf("channel not exist: %d", channelID)
	}

	if channel, exist := b.channels[channelID]; exist {
		close(channel)
		delete(b.channels, channelID)
		return nil
	}

	return fmt.Errorf("channel not exist: %d", channelID)

}

func (b *BytesChannelInt) getChannel(channelID int) (chan bytesChannelData, error) {
	b.lock.RLock()
	defer b.lock.RUnlock()

	if b.channels == nil {
		return nil, fmt.Errorf("channel not exist: %d", channelID)
	}

	if channel, exist := b.channels[channelID]; exist {
		return channel, nil
	}

	return nil, fmt.Errorf("channel not exist: %d", channelID)
}

func (b *BytesChannelInt) WriteChannel(channelID int, bytes []byte, bytesErr error) error {

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

func (b *BytesChannelInt) ReadChannel(channelID int) ([]byte, error) {

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

func (b *BytesChannelInt) ReadChannelJSON(channelID int, data interface{}) error {

	var (
		bytes []byte
		err   error
	)

	if bytes, err = b.ReadChannel(channelID); err != nil {
		return err
	}

	if err = json.Unmarshal(bytes, data); err != nil {
		return err
	}

	return nil

}
