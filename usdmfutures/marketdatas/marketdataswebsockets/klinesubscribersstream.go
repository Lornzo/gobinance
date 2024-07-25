package marketdataswebsockets

import (
	"fmt"
	"strings"
	"sync"
)

type kLineSubscribersForStream interface {
	isEmpty() bool
	subscribe(subscriber KLineSubscriber) error
	unsubscribe(subscriber KLineSubscriber) error
	updateKLine(kLine KLine, err error)
}

type kLineSubscribersStream struct {
	subscriberMap map[string]kLineSubscribersForStream
	lock          sync.RWMutex
}

func (k *kLineSubscribersStream) subscribe(subscriber KLineSubscriber) error {

	k.lock.Lock()
	defer k.lock.Unlock()

	k.init()

	var streamName string = k.getStreamName(subscriber.GetSymbol(), subscriber.GetInterval())

	if _, subscribersExist := k.subscriberMap[streamName]; !subscribersExist {
		k.subscriberMap[streamName] = &kLineSubscribers{}
	}

	return k.subscriberMap[streamName].subscribe(subscriber)

}

func (k *kLineSubscribersStream) unsubscribe(subscriber KLineSubscriber) error {

	k.lock.Lock()
	defer k.lock.Unlock()

	k.init()

	var (
		streamName       string = k.getStreamName(subscriber.GetSymbol(), subscriber.GetInterval())
		subscribers      kLineSubscribersForStream
		subscribersExist bool
		err              error
	)

	if subscribers, subscribersExist = k.subscriberMap[streamName]; !subscribersExist {
		return fmt.Errorf("stream name not found : %s", streamName)
	}

	if err = subscribers.unsubscribe(subscriber); err != nil {
		return err
	}

	if subscribers.isEmpty() {
		delete(k.subscriberMap, streamName)
	}

	return nil

}

func (k *kLineSubscribersStream) updateKLine(kLine KLine, err error) {

	k.lock.RLock()

	defer k.lock.RUnlock()

	k.init()

	var (
		streamName       string = k.getStreamName(kLine.GetSymbol(), kLine.GetInterval())
		subscribers      kLineSubscribersForStream
		subscribersExist bool
	)

	if subscribers, subscribersExist = k.subscriberMap[streamName]; !subscribersExist {
		return
	}

	subscribers.updateKLine(kLine, err)

}

func (k *kLineSubscribersStream) existStreamName(streamName string) bool {

	k.lock.RLock()

	defer k.lock.RUnlock()

	if k.subscriberMap == nil {
		return false
	}

	if _, exist := k.subscriberMap[streamName]; exist {
		return true
	}

	return false

}

func (k *kLineSubscribersStream) getStreamName(symbol string, interval string) string {
	return fmt.Sprint(strings.ToLower(symbol), "@kline_", interval)
}

func (k *kLineSubscribersStream) init() {
	if k.subscriberMap == nil {
		k.subscriberMap = make(map[string]kLineSubscribersForStream)
	}
}
