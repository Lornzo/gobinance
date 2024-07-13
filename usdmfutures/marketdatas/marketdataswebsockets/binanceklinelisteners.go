package marketdataswebsockets

import (
	"fmt"
	"sync"
)

func newBinanceKLineListeners() *binanceKLineListeners {
	return &binanceKLineListeners{
		listeners: make(map[binanceStreamName]kLineListener),
	}
}

type binanceKLineListeners struct {
	listeners     map[binanceStreamName]kLineListener
	listenersLock sync.RWMutex
}

func (b *binanceKLineListeners) hasListener(streamName string) bool {
	if listener := b.getListener(streamName); listener != nil {
		return true
	}
	return false
}

func (b *binanceKLineListeners) getListener(streamName string) kLineListener {

	b.listenersLock.RLock()
	defer b.listenersLock.RUnlock()

	if listener, exist := b.listeners[binanceStreamName(streamName)]; exist {
		return listener
	}

	return nil
}

func (b *binanceKLineListeners) addListener(listener kLineListener) error {

	if exist := b.hasListener(listener.getStreamName()); exist {
		return fmt.Errorf("listener already exist for stream %s", listener.getStreamName())
	}

	b.listenersLock.Lock()
	defer b.listenersLock.Unlock()

	b.listeners[binanceStreamName(listener.getStreamName())] = listener

	return nil

}

func (b *binanceKLineListeners) rmListener(streamName string) error {

	if !b.hasListener(streamName) {
		return fmt.Errorf("listener does not exist for stream %s", streamName)
	}

	b.listenersLock.Lock()
	defer b.listenersLock.Unlock()

	delete(b.listeners, binanceStreamName(streamName))

	return nil
}
