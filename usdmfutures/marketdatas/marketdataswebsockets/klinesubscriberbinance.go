package marketdataswebsockets

import (
	"fmt"
)

type KLineSubscriberBinance struct {
	ID       string
	Symbol   string
	Interval string
}

func (k KLineSubscriberBinance) GetID() string {
	return k.ID
}

func (k KLineSubscriberBinance) GetSymbol() string {
	return k.Symbol
}

func (k KLineSubscriberBinance) GetInterval() string {
	return k.Interval
}

func (k KLineSubscriberBinance) UpdateKLine(kLine KLine, err error) {
	fmt.Println(kLine, err)
	// do something
}
