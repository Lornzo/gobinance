package marketdataswebsockets

import (
	"sync"
)

type dataSourceListenerListItem interface {
	getID() int
	update(data []byte) error
}

func newBinanceDataSourceListenerList() *binanceDataSourceListenerList {
	return &binanceDataSourceListenerList{
		list: make(map[string]dataSourceListenerListItem),
	}
}

// 執行緒安全
type binanceDataSourceListenerList struct {
	list        map[string]dataSourceListenerListItem
	lock        sync.RWMutex
	idCounter   int
	counterLock sync.Mutex
}

func (b *binanceDataSourceListenerList) getID() int {
	b.counterLock.Lock()
	defer b.counterLock.Unlock()
	b.idCounter++
	return b.idCounter
}

func (b *binanceDataSourceListenerList) existStreamName(stream string) bool {
	b.lock.RLock()
	defer b.lock.RUnlock()
	_, exist := b.list[stream]
	return exist
}

func (b *binanceDataSourceListenerList) addListener(listener dataSourceListener) int {

	var (
		streamName string = listener.getStreamName()
		id         int
	)

	if exist := b.existStreamName(streamName); exist {
		return 0
	}

	id = b.getID()

	b.lock.Lock()
	defer b.lock.Unlock()

	b.list[streamName] = &binanceDataSourceListenerListItem{
		id:               id,
		concreteListener: listener,
	}

	return id

}

func (b *binanceDataSourceListenerList) rmListener(listener dataSourceListener) int {

	var (
		streamName string                     = listener.getStreamName()
		listner    dataSourceListenerListItem = b.getItemByStreamName(streamName)
		deletedID  int
	)

	if listner == nil {
		return 0
	}

	b.lock.Lock()
	defer b.lock.Unlock()

	deletedID = listner.getID()
	delete(b.list, streamName)

	return deletedID
}

func (b *binanceDataSourceListenerList) getItemByStreamName(streamName string) dataSourceListenerListItem {

	b.initList()

	b.lock.RLock()
	defer b.lock.RUnlock()

	if listener, exist := b.list[streamName]; exist {
		return listener
	}

	return nil
}

func (b *binanceDataSourceListenerList) update(streamName string, data []byte) error {

	var listener dataSourceListenerListItem = b.getItemByStreamName(streamName)

	if listener == nil {
		return nil
	}

	return listener.update(data)

}

func (b *binanceDataSourceListenerList) initList() {

	b.lock.Lock()

	defer b.lock.Unlock()

	if b.list == nil {
		b.list = make(map[string]dataSourceListenerListItem)
	}
}
