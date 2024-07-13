package marketdataswebsockets

import (
	"encoding/json"
	"fmt"
)

func newBinanceKLineListener(symbol string, interval string) *binanceKLineListener {
	return &binanceKLineListener{
		symbol:      symbol,
		interval:    interval,
		subscribers: newKLineSubscriberMap(),
	}
}

type binanceKLineListener struct {
	symbol      string
	interval    string
	subscribers kLineSubscribers
}

func (b *binanceKLineListener) getStreamName() string {
	var streamName string = fmt.Sprint(b.symbol, "@kline_", b.interval)
	return streamName
}

func (b *binanceKLineListener) update(data []byte) error {

	type binanceKLineDataEvent struct {
		Stream string       `json:"stream"`
		Data   binanceKLine `json:"data"`
	}

	var (
		kLineDataEvent binanceKLineDataEvent
		err            error
	)

	if err = json.Unmarshal(data, &kLineDataEvent); err != nil {
		return err
	}

	if kLineDataEvent.Stream != b.getStreamName() {
		return nil
	}

	b.subscribers.update(kLineDataEvent.Data)

	return nil
}

func (b *binanceKLineListener) subscribe(subscriber KLineSubscriber) error {
	return b.subscribers.add(subscriber)
}

func (b *binanceKLineListener) unsubscribe(subscriber KLineSubscriber) error {
	return b.subscribers.rm(subscriber)
}

func (b *binanceKLineListener) hasSubscribers() bool {
	return !b.subscribers.isEmpty()
}
