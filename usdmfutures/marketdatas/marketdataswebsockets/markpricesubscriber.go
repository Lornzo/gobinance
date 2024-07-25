package marketdataswebsockets

import (
	"fmt"
	"sync"
)

type MarkPriceSubscriber interface {
	GetID() string
	GetSymbol() string
	GetInterval() string
	UpdateMarkPrice(markPrice MarkPrice, err error)
}

type markPriceSubscribers struct {
	subscribers map[string]MarkPriceSubscriber
	lock        sync.RWMutex
}

func (m *markPriceSubscribers) isEmpty() bool {

	m.lock.RLock()

	defer m.lock.RUnlock()

	if m.subscribers == nil {
		return true
	}

	if len(m.subscribers) == 0 {
		return true
	}

	return false

}

func (m *markPriceSubscribers) subscribe(subscriber MarkPriceSubscriber) error {

	m.lock.Lock()

	defer m.lock.Unlock()

	m.init()

	if _, exist := m.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("markprice subscriber has already exist : %s", subscriber.GetID())
	}

	m.subscribers[subscriber.GetID()] = subscriber

	return nil

}

func (m *markPriceSubscribers) unsubscribe(subscriber MarkPriceSubscriber) error {

	m.lock.Lock()

	defer m.lock.Unlock()

	m.init()

	if _, exist := m.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("markprice subscriber not found : %s", subscriber.GetID())
	}

	delete(m.subscribers, subscriber.GetID())

	return nil
}

func (m *markPriceSubscribers) updateMarkPrice(markPrice MarkPrice, err error) {

	m.lock.RLock()

	defer m.lock.RUnlock()

	if m.subscribers == nil {
		return
	}

	for _, subscriber := range m.subscribers {
		subscriber.UpdateMarkPrice(markPrice, err)
	}
}

func (m *markPriceSubscribers) init() {
	if m.subscribers == nil {
		m.subscribers = make(map[string]MarkPriceSubscriber)
	}
}
