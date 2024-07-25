package marketdataswebsockets

import "github.com/shopspring/decimal"

type MarkPrice interface {
	GetEventName() string                     // 事件类型
	GetEventTimestampMilli() int64            // 事件时间
	GetSymbol() string                        // 交易对
	GetMarkPrice() decimal.Decimal            // 标记价格
	GetIndexPrice() decimal.Decimal           // 现货指数价格
	GetEstimatedSettlePrice() decimal.Decimal // 预估结算价,仅在结算前最后一小时有参考价值
	GetFundingRate() decimal.Decimal          // 资金费率
	GetNextFundingTimeMilli() int64           // 下次资金时间
}

type markPrice struct {
	EventName            string          `json:"e"`
	EventTimestampMilli  int64           `json:"E"`
	Symbol               string          `json:"s"`
	MarkPrice            decimal.Decimal `json:"p"`
	IndexPrice           decimal.Decimal `json:"i"`
	EstimatedSettlePrice decimal.Decimal `json:"P"`
	FundingRate          decimal.Decimal `json:"r"`
	NextFundingTimeMilli int64           `json:"T"`
}

func (m markPrice) GetEventName() string {
	return m.EventName
}

func (m markPrice) GetEventTimestampMilli() int64 {
	return m.EventTimestampMilli
}

func (m markPrice) GetSymbol() string {
	return m.Symbol
}

func (m markPrice) GetMarkPrice() decimal.Decimal {
	return m.MarkPrice
}

func (m markPrice) GetIndexPrice() decimal.Decimal {
	return m.IndexPrice
}

func (m markPrice) GetEstimatedSettlePrice() decimal.Decimal {
	return m.EstimatedSettlePrice
}

func (m markPrice) GetFundingRate() decimal.Decimal {
	return m.FundingRate
}

func (m markPrice) GetNextFundingTimeMilli() int64 {
	return m.NextFundingTimeMilli
}
