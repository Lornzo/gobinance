package accountswebsockets

import (
	"fmt"
	"sync"
)

type ErrorSubscriber interface {
	GetID() string
	UpdateError(err error)
}

type errorSubscribers struct {
	lock        sync.RWMutex
	subscribers map[string]ErrorSubscriber
}

func (e *errorSubscribers) subscribe(subscriber ErrorSubscriber) error {

	e.lock.Lock()
	defer e.lock.Unlock()
	e.init()

	if _, exist := e.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	e.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (e *errorSubscribers) unsubscribe(subscriber ErrorSubscriber) error {

	e.lock.Lock()
	defer e.lock.Unlock()

	e.init()

	if _, exist := e.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}

	delete(e.subscribers, subscriber.GetID())

	return nil
}

func (e *errorSubscribers) update(err error) {

	e.lock.RLock()
	defer e.lock.RUnlock()

	if e.subscribers == nil {
		return
	}

	for _, subscriber := range e.subscribers {
		subscriber.UpdateError(err)
	}

}

func (e *errorSubscribers) init() {
	if e.subscribers == nil {
		e.subscribers = make(map[string]ErrorSubscriber)
	}
}
