package marketdataswebsockets

import (
	"fmt"
	"strings"
	"sync"
)

type markPriceSubscribersForStream interface {
	isEmpty() bool
	subscribe(subscriber MarkPriceSubscriber) error
	unsubscribe(subscriber MarkPriceSubscriber) error
	updateMarkPrice(markPrice MarkPrice, err error)
}

type markPriceSubscribersStream struct {
	subscribersMap map[string]markPriceSubscribersForStream
	lock           sync.RWMutex
}

func (m *markPriceSubscribersStream) subscribe(subscriber MarkPriceSubscriber) error {

	m.lock.Lock()

	defer m.lock.Unlock()

	m.init()

	var streamName string = m.getStreamName(subscriber.GetSymbol(), subscriber.GetInterval())

	if _, subscribersExist := m.subscribersMap[streamName]; !subscribersExist {
		m.subscribersMap[streamName] = &markPriceSubscribers{}
	}

	return m.subscribersMap[streamName].subscribe(subscriber)

}

func (m *markPriceSubscribersStream) unsubscribe(subscriber MarkPriceSubscriber) error {

	m.lock.Lock()

	defer m.lock.Unlock()

	m.init()

	var (
		streamName       string = m.getStreamName(subscriber.GetSymbol(), subscriber.GetInterval())
		subscribers      markPriceSubscribersForStream
		subscribersExist bool
		err              error
	)

	if subscribers, subscribersExist = m.subscribersMap[streamName]; !subscribersExist {
		return fmt.Errorf("stream name not found : %s", streamName)
	}

	if err = subscribers.unsubscribe(subscriber); err != nil {
		return err
	}

	if subscribers.isEmpty() {
		delete(m.subscribersMap, streamName)
	}

	return nil

}

func (m *markPriceSubscribersStream) updateMarkPrice(streamName string, markPrice MarkPrice, err error) {

	m.lock.RLock()

	defer m.lock.RUnlock()

	if m.subscribersMap == nil {
		return
	}

	var (
		subscribers      markPriceSubscribersForStream
		subscribersExist bool
	)

	if subscribers, subscribersExist = m.subscribersMap[streamName]; !subscribersExist {
		return
	}

	subscribers.updateMarkPrice(markPrice, err)

}

func (m *markPriceSubscribersStream) existStreamName(streamName string) bool {

	m.lock.RLock()

	defer m.lock.RUnlock()

	if m.subscribersMap == nil {
		return false
	}

	if _, exist := m.subscribersMap[streamName]; exist {
		return true
	}

	return false

}

func (m *markPriceSubscribersStream) getStreamName(symbol string, interval string) string {

	var streamName string = fmt.Sprint(strings.ToLower(symbol), "@markPrice")

	if interval != "" {
		streamName = fmt.Sprint(streamName, "@", interval)
	}

	return streamName

}

// func (m *markPriceSubscribersStream) getStreamNameByStreamName(streamName string) string {

// 	var arr []string = strings.Split(streamName, "@")

// 	if len(arr) == 0 {
// 		return ""
// 	}

// 	arr[0] = strings.ToLower(arr[0])

// 	return strings.Join(arr, "@")
// }

func (m *markPriceSubscribersStream) init() {
	if m.subscribersMap == nil {
		m.subscribersMap = make(map[string]markPriceSubscribersForStream)
	}
}
