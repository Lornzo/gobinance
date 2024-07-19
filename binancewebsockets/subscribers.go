package binancewebsockets

import (
	"fmt"
	"sync"
)

type subscribers struct {
	list map[string]Subscriber
	lock sync.RWMutex
}

func (s *subscribers) add(subscriber Subscriber) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, exist := s.list[subscriber.GetID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetID())
	}

	s.list[subscriber.GetID()] = subscriber

	return nil
}

func (s *subscribers) remove(subscriber Subscriber) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, exist := s.list[subscriber.GetID()]; !exist {
		return fmt.Errorf("subscriber not found : %s", subscriber.GetID())
	}
	delete(s.list, subscriber.GetID())
	return nil
}

func (s *subscribers) getMap() map[string]Subscriber {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.list
}

func (s *subscribers) update(msgType int, msg []byte, err error) {
	var dstSubscribers map[string]Subscriber = s.getMap()
	for _, subscriber := range dstSubscribers {
		subscriber.Update(msgType, msg, err)
	}
}
