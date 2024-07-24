package accountswebsockets

import (
	"fmt"
	"sync"
)

type AccountInfoUpdateSubscriber interface {
	GetID() string
	UpdateAccountInfoUpdate(infoUpdate AccountInfoUpdate, err error)
}

type accountInfoUpdateSubscribers struct {
	lock        sync.RWMutex
	subscribers map[string]AccountInfoUpdateSubscriber
}

func (a *accountInfoUpdateSubscribers) subscribe(subscriber AccountInfoUpdateSubscriber) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.init()

	if _, exist := a.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	a.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (a *accountInfoUpdateSubscribers) unsubscribe(subscriber AccountInfoUpdateSubscriber) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.init()

	if _, exist := a.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}

	delete(a.subscribers, subscriber.GetID())

	return nil

}

func (a *accountInfoUpdateSubscribers) update(infoUpdate AccountInfoUpdate, err error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	if a.subscribers == nil {
		return
	}

	for _, subscriber := range a.subscribers {
		subscriber.UpdateAccountInfoUpdate(infoUpdate, err)
	}

}

func (a *accountInfoUpdateSubscribers) init() {
	if a.subscribers == nil {
		a.subscribers = make(map[string]AccountInfoUpdateSubscriber)
	}
}
