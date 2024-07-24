package accountswebsockets

import (
	"fmt"
	"sync"
)

type ConditionalOrderTriggerRejectSubscriber interface {
	GetID() string
	UpdateConditionalOrderTriggerReject(reject ConditionalOrderTriggerReject, err error)
}

type conditionalOrderTriggerRejectSubscribers struct {
	lock        sync.RWMutex
	subscribers map[string]ConditionalOrderTriggerRejectSubscriber
}

func (c *conditionalOrderTriggerRejectSubscribers) subscribe(subscriber ConditionalOrderTriggerRejectSubscriber) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.init()

	if _, exist := c.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	c.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (c *conditionalOrderTriggerRejectSubscribers) unsubscribe(subscriber ConditionalOrderTriggerRejectSubscriber) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.init()

	if _, exist := c.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}

	delete(c.subscribers, subscriber.GetID())

	return nil

}

func (c *conditionalOrderTriggerRejectSubscribers) update(reject ConditionalOrderTriggerReject, err error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.subscribers == nil {
		return
	}

	for _, subscriber := range c.subscribers {
		subscriber.UpdateConditionalOrderTriggerReject(reject, err)
	}

}

func (c *conditionalOrderTriggerRejectSubscribers) init() {
	if c.subscribers == nil {
		c.subscribers = make(map[string]ConditionalOrderTriggerRejectSubscriber)
	}
}
