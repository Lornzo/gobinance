package marketdataswebsockets

import (
	"fmt"
	"sync"
)

type AggTradeSubscriber interface {
	GetID() string
	GetSymbol() string
	UpdateAggTrade(aggTrade AggTrade, err error)
}

type aggTradeSubscribers struct {
	subscribers map[string]AggTradeSubscriber
	lock        sync.RWMutex
}

func (a *aggTradeSubscribers) isEmpty() bool {

	a.lock.RLock()

	defer a.lock.RUnlock()

	if a.subscribers == nil {
		return true
	}

	if len(a.subscribers) == 0 {
		return true
	}

	return false

}

func (a *aggTradeSubscribers) subscribe(subscriber AggTradeSubscriber) error {

	a.lock.Lock()

	defer a.lock.Unlock()

	a.init()

	if _, exist := a.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("aggtrade subscriber has already exist : %s", subscriber.GetID())
	}

	a.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (a *aggTradeSubscribers) unsubscribe(subscriber AggTradeSubscriber) error {

	a.lock.Lock()

	defer a.lock.Unlock()

	a.init()

	if _, exist := a.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("aggtrade subscriber not found : %s", subscriber.GetID())
	}

	delete(a.subscribers, subscriber.GetID())

	return nil
}

func (a *aggTradeSubscribers) updateAggTrade(aggTrade AggTrade, err error) {

	a.lock.RLock()

	defer a.lock.RUnlock()

	if a.subscribers == nil {
		return
	}

	for _, subscriber := range a.subscribers {
		subscriber.UpdateAggTrade(aggTrade, err)
	}

}

func (a *aggTradeSubscribers) init() {
	if a.subscribers == nil {
		a.subscribers = make(map[string]AggTradeSubscriber)
	}
}
