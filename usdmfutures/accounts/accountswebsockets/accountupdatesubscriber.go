package accountswebsockets

import (
	"fmt"
	"sync"
)

type AccountUpdateSubscriber interface {
	GetID() string
	UpdateAccountUpdate(accountUpdate AccountUpdate, err error)
}

type accountUpdateSubscribers struct {
	lock        sync.RWMutex
	subscribers map[string]AccountUpdateSubscriber
}

func (a *accountUpdateSubscribers) subscribe(subscriber AccountUpdateSubscriber) error {

	a.lock.Lock()
	defer a.lock.Unlock()
	a.init()

	if _, exist := a.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	a.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (a *accountUpdateSubscribers) unsubscribe(subscriber AccountUpdateSubscriber) error {

	a.lock.Lock()
	defer a.lock.Unlock()

	a.init()

	if _, exist := a.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}

	delete(a.subscribers, subscriber.GetID())

	return nil

}

func (a *accountUpdateSubscribers) update(accountUpdate AccountUpdate, err error) {

	a.lock.RLock()
	defer a.lock.RUnlock()

	if a.subscribers == nil {
		return
	}

	for _, subscriber := range a.subscribers {
		subscriber.UpdateAccountUpdate(accountUpdate, err)
	}

}

func (a *accountUpdateSubscribers) init() {
	if a.subscribers == nil {
		a.subscribers = make(map[string]AccountUpdateSubscriber)
	}
}
