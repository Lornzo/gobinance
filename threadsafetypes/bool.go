package threadsafetypes

import "sync"

type Bool struct {
	value bool
	lock  sync.RWMutex
}

func (b *Bool) Set(value bool) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.value = value
}

func (b *Bool) Get() bool {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.value
}
