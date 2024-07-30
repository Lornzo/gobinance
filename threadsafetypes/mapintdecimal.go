package threadsafetypes

import (
	"sync"

	"github.com/shopspring/decimal"
)

type MapIntDecimal struct {
	item map[int]decimal.Decimal
	lock sync.RWMutex
}

func (m *MapIntDecimal) Set(index int, value decimal.Decimal) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.item == nil {
		m.item = make(map[int]decimal.Decimal)
	}

	m.item[index] = value
}

func (m *MapIntDecimal) Get(index int) decimal.Decimal {

	m.lock.RLock()

	defer m.lock.RUnlock()

	if m.item == nil {
		return decimal.Zero
	}

	var (
		value      decimal.Decimal
		valueExist bool
	)

	if value, valueExist = m.item[index]; !valueExist {
		return decimal.Zero
	}

	return value

}

func (m *MapIntDecimal) GetMap() map[int]decimal.Decimal {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.item
}
func (m *MapIntDecimal) List(indexes ...int) []decimal.Decimal {

	var (
		dstMap map[int]decimal.Decimal = m.GetMap()
		arr    []decimal.Decimal
	)

	for _, index := range indexes {

		if v, ok := dstMap[index]; ok {
			arr = append(arr, v)
		} else {
			arr = append(arr, decimal.Zero)
		}

	}

	return arr

}
