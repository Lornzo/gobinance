package accountswebsockets

import (
	"fmt"
	"sync"
)

type RateLimitsSubscriber interface {
	GetID() string
	UpdateRateLimits(rateLimits RateLimits, err error)
}

type rateLimitsSubscribers struct {
	lock        sync.RWMutex
	subscirbers map[string]RateLimitsSubscriber
}

func (r *rateLimitsSubscribers) subscribe(subscriber RateLimitsSubscriber) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.init()

	if _, exist := r.subscirbers[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	r.subscirbers[subscriber.GetID()] = subscriber

	return nil
}

func (r *rateLimitsSubscribers) unsubscribe(subscriber RateLimitsSubscriber) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.init()

	if _, exist := r.subscirbers[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}

	delete(r.subscirbers, subscriber.GetID())

	return nil

}

func (r *rateLimitsSubscribers) update(rateLimits RateLimits, err error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	if r.subscirbers == nil {
		return
	}

	for _, subscriber := range r.subscirbers {
		subscriber.UpdateRateLimits(rateLimits, err)
	}

}

func (r *rateLimitsSubscribers) init() {
	if r.subscirbers == nil {
		r.subscirbers = make(map[string]RateLimitsSubscriber)
	}
}
