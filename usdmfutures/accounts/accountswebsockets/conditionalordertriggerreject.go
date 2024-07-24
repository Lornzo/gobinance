package accountswebsockets

type ConditionalOrderTriggerReject interface {
	GetEventName() string
	GetEventTimestampMilli() int64
	GetTimestampMilli() int64
	GetSymbol() string
	GetOrderID() int64
	GetReason() string
}

type conditionalOrderTriggerReject struct {
	EventName           string `json:"e"`
	EventTimestampMilli int64  `json:"E"`
	TimestampMilli      int64  `json:"T"`
	OrderReject         struct {
		Symbol  string `json:"s"`
		OrderID int64  `json:"i"`
		Reason  string `json:"r"`
	} `json:"or"`
}

func (c conditionalOrderTriggerReject) GetEventName() string {
	return c.EventName
}

func (c conditionalOrderTriggerReject) GetEventTimestampMilli() int64 {
	return c.EventTimestampMilli
}

func (c conditionalOrderTriggerReject) GetTimestampMilli() int64 {
	return c.TimestampMilli
}

func (c conditionalOrderTriggerReject) GetSymbol() string {
	return c.OrderReject.Symbol
}

func (c conditionalOrderTriggerReject) GetOrderID() int64 {
	return c.OrderReject.OrderID
}

func (c conditionalOrderTriggerReject) GetReason() string {
	return c.OrderReject.Reason
}
