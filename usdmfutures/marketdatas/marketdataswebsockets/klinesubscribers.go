package marketdataswebsockets

type kLineSubscribers interface {
	add(subscriber KLineSubscriber) error
	rm(subscriber KLineSubscriber) error
	isEmpty() bool
	update(kLine KLine)
}
