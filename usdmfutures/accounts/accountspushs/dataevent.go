package accountspushs

type DataEvent struct {
	EventName string `json:"e"`
	EventTs   int64  `json:"E"`
}

func (d DataEvent) GetEventName() string {
	return d.EventName
}

func (d DataEvent) GetEventTs() int64 {
	return d.EventTs
}
