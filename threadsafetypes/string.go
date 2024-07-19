package threadsafetypes

import "sync"

type String struct {
	value string
	lock  sync.RWMutex
}

func (s *String) Set(str string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.value = str

}

func (s *String) Get() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.value
}
