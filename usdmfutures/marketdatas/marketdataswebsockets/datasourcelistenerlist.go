package marketdataswebsockets

type dataSourceListenerList interface {
	addListener(listener dataSourceListener) int
	rmListener(listener dataSourceListener) int
	update(streamName string, data []byte) error
}
