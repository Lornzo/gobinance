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
		kLineStream string        = fmt.Sprint(subscriber.GetSymbol(), "@kline_", subscriber.GetInterval())
		listener    kLineListener = d.kLineListeners.getListener(kLineStream)
	)

	if listener == nil {
		listener = newBinanceKLineListener(subscriber.GetSymbol(), subscriber.GetInterval())
		d.kLineListeners.addListener(listener)
		d.concreteDataSource.addListener(listener)
	}

	return listener.subscribe(subscriber)

}
