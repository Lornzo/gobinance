package marketdataswebsockets

type dataSourceListener interface {
	getStreamName() string
	update(data []byte) error
}
