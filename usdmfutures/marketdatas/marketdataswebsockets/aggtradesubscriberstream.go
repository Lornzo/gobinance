package marketdataswebsockets

import (
	"fmt"
	"strings"
	"sync"
)

type aggTradeSubscribersForStream interface {
	isEmpty() bool
	subscribe(subscriber AggTradeSubscriber) error
	unsubscribe(subscriber AggTradeSubscriber) error
	updateAggTrade(aggTrade AggTrade, err error)
}

type aggTradeSubscribersStream struct {
	subscriberMap map[string]aggTradeSubscribersForStream
	lock          sync.RWMutex
}

func (a *aggTradeSubscribersStream) subscribe(subscriber AggTradeSubscriber) error {

	a.lock.Lock()

	defer a.lock.Unlock()

	a.init()

	var streamName string = a.getStreamName(subscriber.GetSymbol())

	if _, subscribersExist := a.subscriberMap[streamName]; !subscribersExist {
		a.subscriberMap[streamName] = &aggTradeSubscribers{}
	}

	return a.subscriberMap[streamName].subscribe(subscriber)

}

func (a *aggTradeSubscribersStream) unsubscribe(subscriber AggTradeSubscriber) error {

	a.lock.Lock()

	defer a.lock.Unlock()

	a.init()

	var (
		streamName       string = a.getStreamName(subscriber.GetSymbol())
		subscribers      aggTradeSubscribersForStream
		subscribersExist bool
		err              error
	)

	if subscribers, subscribersExist = a.subscriberMap[streamName]; !subscribersExist {
		return fmt.Errorf("stream name not found : %s", streamName)
	}

	if err = subscribers.unsubscribe(subscriber); err != nil {
		return err
	}

	if subscribers.isEmpty() {
		delete(a.subscriberMap, streamName)
	}

	return nil

}

func (a *aggTradeSubscribersStream) updateAggTrade(aggTrade AggTrade, err error) {

	a.lock.RLock()

	defer a.lock.RUnlock()

	a.init()

	var streamName string = a.getStreamName(aggTrade.GetSymbol())

	if subscribers, exist := a.subscriberMap[streamName]; exist {
		subscribers.updateAggTrade(aggTrade, err)
	}

}

func (a *aggTradeSubscribersStream) existStreamName(streamName string) bool {

	a.lock.RLock()

	defer a.lock.RUnlock()

	if _, exist := a.subscriberMap[streamName]; exist {
		return true
	}

	return false
}

func (a *aggTradeSubscribersStream) getStreamName(symbol string) string {
	return fmt.Sprint(strings.ToLower(symbol), "@aggTrade")
}

func (a *aggTradeSubscribersStream) init() {
	if a.subscriberMap == nil {
		a.subscriberMap = make(map[string]aggTradeSubscribersForStream)
	}
}
