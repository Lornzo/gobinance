package accountspushs

type DataConditionalOrderTriggerReject struct {
	DataEvent
	Timestamp int64 `json:"T"`
	Order     struct {
		Symbol  string `json:"s"` // 交易對
		OrderID int64  `json:"i"` // 訂單號
		Reason  string `json:"r"` // 拒絕原因
	} `json:"or"`
}

func (d DataConditionalOrderTriggerReject) GetTimestamp() int64 {
	return d.Timestamp
}

func (d DataConditionalOrderTriggerReject) GetSymbol() string {
	return d.Order.Symbol
}

func (d DataConditionalOrderTriggerReject) GetOrderID() int64 {
	return d.Order.OrderID
}

func (d DataConditionalOrderTriggerReject) GetReason() string {
	return d.Order.Reason
}

type ConditionalOrderTriggerReject interface {
	GetEventName() string
	GetEventTs() int64
	GetTimestamp() int64
	GetSymbol() string
	GetOrderID() int64
	GetReason() string
}
