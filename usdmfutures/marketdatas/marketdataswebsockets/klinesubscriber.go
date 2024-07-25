package marketdataswebsockets

import (
	"fmt"
	"sync"
)

type KLineSubscriber interface {
	GetID() string
	GetSymbol() string
	GetInterval() string
	UpdateKLine(kLine KLine, err error)
}

type kLineSubscribers struct {
	subscribers map[string]KLineSubscriber
	lock        sync.RWMutex
}

func (k *kLineSubscribers) isEmpty() bool {
	k.lock.RLock()
	defer k.lock.RUnlock()

	if k.subscribers == nil {
		return true
	}

	if len(k.subscribers) == 0 {
		return true
	}

	return false

}

func (k *kLineSubscribers) subscribe(subscriber KLineSubscriber) error {

	k.lock.Lock()
	defer k.lock.Unlock()

	k.init()

	if _, exist := k.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("kline subscriber has already exist : %s", subscriber.GetID())
	}

	k.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (k *kLineSubscribers) unsubscribe(subscriber KLineSubscriber) error {

	k.lock.Lock()
	defer k.lock.Unlock()

	k.init()

	if _, exist := k.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("kline subscriber not found : %s", subscriber.GetID())
	}

	delete(k.subscribers, subscriber.GetID())

	return nil
}

func (k *kLineSubscribers) updateKLine(kLine KLine, err error) {

	k.lock.RLock()
	defer k.lock.RUnlock()

	if k.subscribers == nil {
		return
	}

	for _, subscriber := range k.subscribers {
		subscriber.UpdateKLine(kLine, err)
	}

}

func (k *kLineSubscribers) init() {
	if k.subscribers == nil {
		k.subscribers = make(map[string]KLineSubscriber)
	}
}
