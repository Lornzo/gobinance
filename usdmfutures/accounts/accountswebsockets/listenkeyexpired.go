package accountswebsockets

type ListenKeyExpired interface {
	GetEventName() string
	GetEventTimestampMilli() int64
}

type listenKeyExpired struct {
	EventName           string `json:"e"`
	EventTimestampMilli int64  `json:"E"`
}

func (l listenKeyExpired) GetEventName() string {
	return l.EventName
}

func (l listenKeyExpired) GetEventTimestampMilli() int64 {
	return l.EventTimestampMilli
}
