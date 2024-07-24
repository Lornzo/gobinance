package usdmfutureswebsockettypes

import (
	"fmt"
	"sync"
)

type ErrorSubscriber interface {
	GetID() string
	UpdateError(err error)
}

type ErrorSubscribers struct {
	lock        sync.RWMutex
	subscribers map[string]ErrorSubscriber
}

func (e *ErrorSubscribers) UpdateError(err error) {

	e.lock.RLock()

	defer e.lock.RUnlock()

	if e.subscribers == nil {
		return
	}

	for _, subscriber := range e.subscribers {
		subscriber.UpdateError(err)
	}

}

func (e *ErrorSubscribers) Subscribe(subscriber ErrorSubscriber) error {

	e.lock.Lock()

	defer e.lock.Unlock()

	e.init()

	if _, exist := e.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("error ! subscriber has already exist : %s", subscriber.GetID())
	}

	e.subscribers[subscriber.GetID()] = subscriber

	return nil

}

func (e *ErrorSubscribers) UnSubscribe(subscriber ErrorSubscriber) error {

	e.lock.Lock()

	defer e.lock.Unlock()

	e.init()

	if _, exist := e.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("error ! subscriber not found : %s", subscriber.GetID())
	}

	delete(e.subscribers, subscriber.GetID())

	return nil

}

func (e *ErrorSubscribers) init() {
	if e.subscribers == nil {
		e.subscribers = make(map[string]ErrorSubscriber)
	}
}
