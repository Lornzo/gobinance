package accountswebsockets

import (
	"fmt"
	"sync"
)

type ListenKeyExpiredSubscriber interface {
	GetID() string
	UpdateListenKeyExpired(listenKeyExpired ListenKeyExpired, err error)
}

type listenKeyExpiredSubscribers struct {
	lock        sync.RWMutex
	subscribers map[string]ListenKeyExpiredSubscriber
}

func (l *listenKeyExpiredSubscribers) subscribe(subscriber ListenKeyExpiredSubscriber) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.init()

	if _, exist := l.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	l.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (l *listenKeyExpiredSubscribers) unsubscribe(subscriber ListenKeyExpiredSubscriber) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.init()

	if _, exist := l.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}

	delete(l.subscribers, subscriber.GetID())

	return nil

}

func (l *listenKeyExpiredSubscribers) update(listenKeyExpired ListenKeyExpired, err error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if l.subscribers == nil {
		return
	}

	for _, subscriber := range l.subscribers {
		subscriber.UpdateListenKeyExpired(listenKeyExpired, err)
	}

}

func (l *listenKeyExpiredSubscribers) init() {
	if l.subscribers == nil {
		l.subscribers = make(map[string]ListenKeyExpiredSubscriber)
	}
}
