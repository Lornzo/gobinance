package accountswebsockets

import (
	"fmt"
	"sync"
)

type MarginCallSubscriber interface {
	GetID() string
	UpdateMarginCall(marginCall MarginCall, err error)
}

type marginCallSubscribers struct {
	lock        sync.RWMutex
	subscribers map[string]MarginCallSubscriber
}

func (m *marginCallSubscribers) subscribe(subscriber MarginCallSubscriber) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.init()

	if _, exist := m.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	m.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (m *marginCallSubscribers) unsubscribe(subscriber MarginCallSubscriber) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.init()

	if _, exist := m.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}

	delete(m.subscribers, subscriber.GetID())

	return nil

}

func (m *marginCallSubscribers) update(marginCall MarginCall, err error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.subscribers == nil {
		return
	}

	for _, subscriber := range m.subscribers {
		subscriber.UpdateMarginCall(marginCall, err)
	}

}

func (m *marginCallSubscribers) init() {
	if m.subscribers == nil {
		m.subscribers = make(map[string]MarginCallSubscriber)
	}
}
