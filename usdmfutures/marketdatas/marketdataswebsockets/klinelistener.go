package marketdataswebsockets

type kLineListener interface {
	dataSourceListener
	subscribe(subscriber KLineSubscriber) error
	unsubscribe(subscriber KLineSubscriber) error
	hasSubscribers() bool
}
