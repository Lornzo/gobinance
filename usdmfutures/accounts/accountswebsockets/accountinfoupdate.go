package accountswebsockets

type AccountInfoUpdate interface {
	GetEventName() string
	GetEventTimestampMilli() int64
	GetTimestampMilli() int64
	GetJointMarginStatus() bool
}

type accountInfoUpdate struct {
	EventName           string `json:"e"`
	EventTimestampMilli int64  `json:"E"`
	TimestampMilli      int64  `json:"T"`
	AccountInfo         struct {
		JointMarginStatus bool `json:"j"`
	} `json:"ai"`
}

func (a accountInfoUpdate) GetEventName() string {
	return a.EventName
}

func (a accountInfoUpdate) GetEventTimestampMilli() int64 {
	return a.EventTimestampMilli
}

func (a accountInfoUpdate) GetTimestampMilli() int64 {
	return a.TimestampMilli
}
func (a accountInfoUpdate) GetJointMarginStatus() bool {
	return a.AccountInfo.JointMarginStatus
}
