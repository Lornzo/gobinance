package marketdataswebsockets

type binanceDataSourceListenerListItem struct {
	id               int
	concreteListener interface {
		update(data []byte) error
	}
}

func (b *binanceDataSourceListenerListItem) getID() int {
	return b.id
}

func (b *binanceDataSourceListenerListItem) update(data []byte) error {
	return b.concreteListener.update(data)
}
