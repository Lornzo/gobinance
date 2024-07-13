package marketdataswebsockets

import (
	"fmt"
	"sync"
)

func newKLineSubscriberMap() *kLineSubscriberMap {
	return &kLineSubscriberMap{
		subscribers: make(map[string]KLineSubscriber),
	}
}

type kLineSubscriberMap struct {
	subscribers     map[string]KLineSubscriber
	subscribersLock sync.RWMutex
}

func (k *kLineSubscriberMap) get(id string) KLineSubscriber {
	k.subscribersLock.RLock()
	defer k.subscribersLock.RUnlock()

	var (
		subscriber KLineSubscriber
		exist      bool
	)

	if subscriber, exist = k.subscribers[id]; exist {
		return subscriber
	}

	return nil
}

func (k *kLineSubscriberMap) exist(id string) bool {
	if subscriber := k.get(id); subscriber != nil {
		return true
	}
	return false
}

func (k *kLineSubscriberMap) add(subscriber KLineSubscriber) error {

	if exist := k.exist(subscriber.GetUUID()); exist {
		return fmt.Errorf("subscriber %s already exist", subscriber.GetUUID())
	}

	k.subscribersLock.Lock()
	defer k.subscribersLock.Unlock()

	k.subscribers[subscriber.GetUUID()] = subscriber

	return nil

}

func (k *kLineSubscriberMap) rm(subscriber KLineSubscriber) error {

	if exist := k.exist(subscriber.GetUUID()); !exist {
		return fmt.Errorf("subscriber %s not exist", subscriber.GetUUID())
	}

	k.subscribersLock.Lock()
	defer k.subscribersLock.Unlock()

	delete(k.subscribers, subscriber.GetUUID())

	return nil
}

func (k *kLineSubscriberMap) length() int {
	k.subscribersLock.RLock()
	defer k.subscribersLock.RUnlock()
	return len(k.subscribers)
}

func (k *kLineSubscriberMap) isEmpty() bool {
	return k.length() == 0
}

func (k *kLineSubscriberMap) update(kLine KLine) {
	k.subscribersLock.Lock()
	defer k.subscribersLock.Unlock()
	for _, subscriber := range k.subscribers {
		subscriber.Update(kLine)
	}
}
