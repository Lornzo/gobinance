package threadsafetypes

import "sync"

type MapStringInterface struct {
	value map[string]interface{}
	lock  sync.RWMutex
}

func (m *MapStringInterface) Set(key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.value == nil {
		m.value = make(map[string]interface{})
	}
	m.value[key] = value
}

func (m *MapStringInterface) Get(key string) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.value == nil {
		return nil
	}

	if v, ok := m.value[key]; ok {
		return v
	}
	return nil
}

func (m *MapStringInterface) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.value == nil {
		return
	}

	delete(m.value, key)
}
