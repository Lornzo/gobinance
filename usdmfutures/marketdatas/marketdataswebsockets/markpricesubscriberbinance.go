package marketdataswebsockets

import "fmt"

type MarkPriceSubscriberBinance struct {
	ID       string
	Symbol   string
	Interval int
}

func (m MarkPriceSubscriberBinance) GetID() string {
	return m.ID
}

func (m MarkPriceSubscriberBinance) GetSymbol() string {
	return m.Symbol
}

func (m MarkPriceSubscriberBinance) GetInterval() string {

	if m.Interval == 0 {
		return ""
	}

	return fmt.Sprint(m.Interval, "s")
}

func (m MarkPriceSubscriberBinance) UpdateMarkPrice(markPrice MarkPrice, err error) {

	fmt.Println(m.Symbol, "@", m.Interval, ":", markPrice, err)

}
