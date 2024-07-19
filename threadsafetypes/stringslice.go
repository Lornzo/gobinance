package threadsafetypes

import "sync"

type StringSlice struct {
	value []string
	lock  sync.RWMutex
}

func (s *StringSlice) Set(str ...string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.value = str
}

func (s *StringSlice) Add(str ...string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.value = append(s.value, str...)
}

func (s *StringSlice) Get() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.value
}
