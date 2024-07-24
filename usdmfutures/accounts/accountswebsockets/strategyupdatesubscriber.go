package accountswebsockets

import (
	"fmt"
	"sync"
)

type StrategyUpdateSubscriber interface {
	GetID() string
	UpdateStrategyUpdate(strategyUpdate StrategyUpdate, err error)
}

type strategyUpdateSubscribers struct {
	lock        sync.RWMutex
	subscribers map[string]StrategyUpdateSubscriber
}

func (s *strategyUpdateSubscribers) subscribe(subscriber StrategyUpdateSubscriber) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.init()

	if _, exist := s.subscribers[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	s.subscribers[subscriber.GetID()] = subscriber

	return nil
}

func (s *strategyUpdateSubscribers) unsubscribe(subscriber StrategyUpdateSubscriber) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.init()

	if _, exist := s.subscribers[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}

	delete(s.subscribers, subscriber.GetID())

	return nil

}

func (s *strategyUpdateSubscribers) update(strategyUpdate StrategyUpdate, err error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.subscribers == nil {
		return
	}

	for _, subscriber := range s.subscribers {
		subscriber.UpdateStrategyUpdate(strategyUpdate, err)
	}

}

func (s *strategyUpdateSubscribers) init() {
	if s.subscribers == nil {
		s.subscribers = make(map[string]StrategyUpdateSubscriber)
	}
}
