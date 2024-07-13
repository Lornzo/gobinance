package marketdataswebsockets

import (
	"context"
	"fmt"
)

type dataSource interface {
	start(ctx context.Context) error
	close() error
	addListener(listener dataSourceListener) error
	rmListener(listener dataSourceListener) error
}

func NewDataSource(baseURL string) *DataSource {
	return &DataSource{
		concreteDataSource: newBinanceDataSource(baseURL),
		kLineListeners:     newBinanceKLineListeners(),
	}

}

type DataSource struct {
	concreteDataSource dataSource
	kLineListeners     kLineListeners
}

func (d *DataSource) Start(ctx context.Context) error {
	return d.concreteDataSource.start(ctx)
}

func (d *DataSource) Close() error {
	return d.concreteDataSource.close()
}

func (d *DataSource) SubscribeKLine(subscriber KLineSubscriber) error {

	var (
		kLineStream string        = d.getKLineStream(subscriber.GetSymbol(), subscriber.GetInterval())
		listener    kLineListener = d.kLineListeners.getListener(kLineStream)
	)

	if listener == nil {
		listener = newBinanceKLineListener(subscriber.GetSymbol(), subscriber.GetInterval())
		d.kLineListeners.addListener(listener)
		d.concreteDataSource.addListener(listener)
	}

	return listener.subscribe(subscriber)

}

func (d *DataSource) UnSubscribeKLine(subscriber KLineSubscriber) error {

	var (
		streamName string        = d.getKLineStream(subscriber.GetSymbol(), subscriber.GetInterval())
		listener   kLineListener = d.kLineListeners.getListener(streamName)
		err        error
	)

	if listener == nil {
		return nil
	}

	if err = listener.unsubscribe(subscriber); err != nil {
		return err
	}

	if listener.hasSubscribers() {
		return nil
	}

	if err = d.concreteDataSource.rmListener(listener); err != nil {
		return err
	}

	return d.kLineListeners.rmListener(listener.getStreamName())

}

func (d *DataSource) getKLineStream(symbol string, interval string) string {
	return fmt.Sprint(symbol, "@kline_", interval)
}
