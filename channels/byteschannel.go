package channels

import (
	"errors"
	"sync"
)

type bytesChannelData struct {
	bytes []byte
	err   error
}

type BytesChannel struct {
	channels map[string]chan bytesChannelData
	lock     sync.RWMutex
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

func (b *BytesChannel) CreateChannel(channelID string) error {

	b.lock.Lock()

	defer b.lock.Unlock()

	if b.channels == nil {
		b.channels = make(map[string]chan bytesChannelData)
	}

	if _, exist := b.channels[channelID]; exist {
		return errors.New("channel already exist")
	}
	b.channels[channelID] = make(chan bytesChannelData)

	return nil
}

func (b *BytesChannel) CloseChannel(channelID string) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.channels == nil {
		return errors.New("channel not exist")
	}

	if channel, exist := b.channels[channelID]; exist {
		close(channel)
		delete(b.channels, channelID)
		return nil
	}
	return errors.New("channel not exist")

}

func (b *BytesChannel) getChannel(channelID string) (chan bytesChannelData, error) {
	b.lock.RLock()
	defer b.lock.RUnlock()

	if b.channels == nil {
		return nil, errors.New("channel not exist")

	}

	if channel, exist := b.channels[channelID]; exist {
		return channel, nil
	}
	return nil, errors.New("channel not exist")
}

func (b *BytesChannel) WriteChannel(channelID string, bytes []byte, bytesErr error) error {

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

func (b *BytesChannel) ReadChannel(channelID string) ([]byte, error) {

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
