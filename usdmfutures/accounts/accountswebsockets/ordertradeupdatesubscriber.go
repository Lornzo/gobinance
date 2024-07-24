package accountswebsockets

import (
	"fmt"
	"sync"
)

type OrderTradeUpdateSubscriber interface {
	GetID() string
	UpdateOrderTrade(orderTrade OrderTradeUpdate, err error)
}

type orderTradeUpdateSubscribers struct {
	lock        sync.RWMutex
	subscribers map[string]OrderTradeUpdateSubscriber
}

func (o *orderTradeUpdateSubscribers) subscribe(subscriber OrderTradeUpdateSubscriber) error {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.init()

	if _, exist := o.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	o.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (o *orderTradeUpdateSubscribers) unsubscribe(subscriber OrderTradeUpdateSubscriber) error {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.init()

	if _, exist := o.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}

	delete(o.subscribers, subscriber.GetID())

	return nil

}

func (o *orderTradeUpdateSubscribers) update(orderTrade OrderTradeUpdate, err error) {
	o.lock.RLock()
	defer o.lock.RUnlock()

	if o.subscribers == nil {
		return
	}

	for _, subscriber := range o.subscribers {
		subscriber.UpdateOrderTrade(orderTrade, err)
	}

}

func (o *orderTradeUpdateSubscribers) init() {
	if o.subscribers == nil {
		o.subscribers = make(map[string]OrderTradeUpdateSubscriber)
	}
}
