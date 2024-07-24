package accountswebsockets

import (
	"fmt"
	"sync"
)

type AccountConfigUpdateSubscriber interface {
	GetID() string
	UpdateAccountConfigUpdate(configUpdate AccountConfigUpdate, err error)
}

type accountConfigUpdateSubscribers struct {
	lock        sync.RWMutex
	subscribers map[string]AccountConfigUpdateSubscriber
}

func (a *accountConfigUpdateSubscribers) subscribe(subscriber AccountConfigUpdateSubscriber) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.init()

	if _, exist := a.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	a.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (a *accountConfigUpdateSubscribers) unsubscribe(subscriber AccountConfigUpdateSubscriber) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.init()

	if _, exist := a.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}

	delete(a.subscribers, subscriber.GetID())

	return nil

}

func (a *accountConfigUpdateSubscribers) update(configUpdate AccountConfigUpdate, err error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	if a.subscribers == nil {
		return
	}

	for _, subscriber := range a.subscribers {
		subscriber.UpdateAccountConfigUpdate(configUpdate, err)
	}

}

func (a *accountConfigUpdateSubscribers) init() {
	if a.subscribers == nil {
		a.subscribers = make(map[string]AccountConfigUpdateSubscriber)
	}
}
