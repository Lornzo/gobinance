package marketdataswebsockets

type kLineListeners interface {
	hasListener(streamName string) bool
	getListener(streamName string) kLineListener
	addListener(listener kLineListener) error
	rmListener(streamName string) error
}
