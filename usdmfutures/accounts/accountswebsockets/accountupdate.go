package accountswebsockets

type AccountUpdate interface {
	GetEventName() string
	GetEventTimestampMilli() int64
	GetTimestampMilli() int64
	GetUpdateType() string
	GetBalances() []Balance
	GetPositions() []Position
}

type accountUpdate struct {
	EventName           string `json:"e"`
	EventTimestampMilli int64  `json:"E"`
	TimestampMilli      int64  `json:"T"`
	AccountEvent        struct {
		MessageType string     `json:"m"`
		Balances    []Balance  `json:"B"`
		Positions   []Position `json:"P"`
	} `json:"a"`
}

func (a accountUpdate) GetEventName() string {
	return a.EventName
}

func (a accountUpdate) GetEventTimestampMilli() int64 {
	return a.EventTimestampMilli
}

func (a accountUpdate) GetTimestampMilli() int64 {
	return a.TimestampMilli
}

func (a accountUpdate) GetUpdateType() string {
	return a.AccountEvent.MessageType
}

func (a accountUpdate) GetBalances() []Balance {
	return a.AccountEvent.Balances
}

func (a accountUpdate) GetPositions() []Position {
	return a.AccountEvent.Positions
}
