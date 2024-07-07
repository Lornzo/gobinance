package accountspushs

import "github.com/shopspring/decimal"

type DataMarginCallPosition struct {
	Symbol           string          `json:"s"`  // symbol
	Position         string          `json:"ps"` // 持倉方向
	PositionAmount   decimal.Decimal `json:"pa"` // 倉位
	MarginMode       string          `json:"mt"` // 保證金模式
	PositionMargin   decimal.Decimal `json:"iw"` // 仓位保证金
	MarketPrice      decimal.Decimal `json:"mp"` // 標記價格
	UnrealizedProfit decimal.Decimal `json:"up"` // 持倉未實現盈虧
	MustMargin       decimal.Decimal `json:"mm"` // 持仓需要的维持保证金
}

type DataMarginCallPositions []DataMarginCallPosition

type DataMarginCall struct {
	DataEvent
	CrossWallet decimal.Decimal         `json:"cw"` // 除去逐仓仓位保证金的钱包余额, 仅在全仓 margin call 情况下推送此字段
	Positions   DataMarginCallPositions `json:"p"`  // 仓位信息
}

func (d DataMarginCall) GetCrossWallet() decimal.Decimal {
	return d.CrossWallet
}

func (d DataMarginCall) GetPositions() DataMarginCallPositions {
	return d.Positions
}

type MarginCall interface {
	GetEventName() string
	GetEventTs() int64
	GetCrossWallet() decimal.Decimal
	GetPositions() DataMarginCallPositions
}
