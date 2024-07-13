package marketdataswebsockets

type KLineSubscriber interface {
	GetSymbol() string
	GetInterval() string
	GetUUID() string
	Update(kLine KLine)
}
