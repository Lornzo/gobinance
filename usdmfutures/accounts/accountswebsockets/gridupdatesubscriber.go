package accountswebsockets

import (
	"fmt"
	"sync"
)

type GridUpdateSubscriber interface {
	GetID() string
	UpdateGridUpdate(gridUpdate GridUpdate, err error)
}

type gridUpdateSubscribers struct {
	lock        sync.RWMutex
	subscribers map[string]GridUpdateSubscriber
}

func (g *gridUpdateSubscribers) subscribe(subscriber GridUpdateSubscriber) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.init()

	if _, exist := g.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	g.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (g *gridUpdateSubscribers) unsubscribe(subscriber GridUpdateSubscriber) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.init()

	if _, exist := g.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}

	delete(g.subscribers, subscriber.GetID())

	return nil

}

func (g *gridUpdateSubscribers) update(gridUpdate GridUpdate, err error) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	if g.subscribers == nil {
		return
	}

	for _, subscriber := range g.subscribers {
		subscriber.UpdateGridUpdate(gridUpdate, err)
	}

}

func (g *gridUpdateSubscribers) init() {
	if g.subscribers == nil {
		g.subscribers = make(map[string]GridUpdateSubscriber)
	}
}
