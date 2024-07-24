package accountswebsockets

import "github.com/shopspring/decimal"

type MarginCall interface {
	GetEventName() string
	GetEventTimestampMilli() int64
	GetCrossWalletBalance() decimal.Decimal
	GetPositions() []MarginCallPosition
}

type marginCall struct {
	EventName           string               `json:"e"`
	EventTimestampMilli int64                `json:"E"`
	CrossWalletBalance  decimal.Decimal      `json:"cw"` // 除去逐仓仓位保证金的钱包余额, 仅在全仓 margin call 情况下推送此字段
	Positions           []MarginCallPosition `json:"p"`
}

func (m marginCall) GetEventName() string {
	return m.EventName
}

func (m marginCall) GetEventTimestampMilli() int64 {
	return m.EventTimestampMilli
}

func (m marginCall) GetCrossWalletBalance() decimal.Decimal {
	return m.CrossWalletBalance
}

func (m marginCall) GetPositions() []MarginCallPosition {
	return m.Positions
}
