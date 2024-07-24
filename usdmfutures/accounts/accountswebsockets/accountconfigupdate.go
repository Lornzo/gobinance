package accountswebsockets

type AccountConfigUpdate interface {
	GetEventName() string
	GetEventTimestampMilli() int64
	GetTimestampMilli() int64
	GetSymbol() string
	GetLeverage() int64
}

type accountConfigUpdate struct {
	EventName           string `json:"e"`
	EventTimestampMilli int64  `json:"E"`
	TimestampMilli      int64  `json:"T"`
	AccountConfig       struct {
		Symbol   string `json:"s"`
		Leverage int64  `json:"l"`
	} `json:"ac,omitempty"`
}

func (a accountConfigUpdate) GetEventName() string {
	return a.EventName
}

func (a accountConfigUpdate) GetEventTimestampMilli() int64 {
	return a.EventTimestampMilli
}

func (a accountConfigUpdate) GetTimestampMilli() int64 {
	return a.TimestampMilli
}

func (a accountConfigUpdate) GetSymbol() string {
	return a.AccountConfig.Symbol
}

func (a accountConfigUpdate) GetLeverage() int64 {
	return a.AccountConfig.Leverage
}
