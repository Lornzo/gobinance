package accountspushs

type DataListenKeyExpired struct {
	DataEvent
}

type ListenKeyExpired interface {
	GetEventName() string
	GetEventTs() int64
}
