package accountswebsockets

import "github.com/shopspring/decimal"

type GridUpdate interface {
	GetEventName() string
	GetEventTimestampMilli() int64
	GetTimestampMilli() int64
	GetStrategyID() int64                      // 策略 ID
	GetStrategyType() string                   // 策略类型
	GetStrategyStatus() string                 // 策略状态
	GetSymbol() string                         // 交易对
	GetRealizedPNL() decimal.Decimal           // 已实现 PNL
	GetUnMatchedAveragePrice() decimal.Decimal // 未配对均价
	GetUnMatchedQuantity() decimal.Decimal     // 未配对数量
	GetUnMatchedFee() decimal.Decimal          // 未配对手续费
	GetMatchedPNL() decimal.Decimal            // 已配对 PNL
	GetUpdateTimestampMilli() int64            // 更新时间
}

type gridUpdate struct {
	EventName           string `json:"e"`
	EventTimestampMilli int64  `json:"E"`
	TimestampMilli      int64  `json:"T"`
	GridUpdate          struct {
		StrategyID            int64           `json:"si"`
		StrategyType          string          `json:"st"`
		StrategyStatus        string          `json:"ss"`
		Symbol                string          `json:"s"`
		RealizedPNL           decimal.Decimal `json:"r"`
		UnMatchedAveragePrice decimal.Decimal `json:"up"`
		UnMatchedQuantity     decimal.Decimal `json:"uq"`
		UnMatchedFee          decimal.Decimal `json:"uf"`
		MatchedPNL            decimal.Decimal `json:"mp"`
		UpdateTimestampMilli  int64           `json:"ut"`
	} `json:"gu"`
}

func (g gridUpdate) GetEventName() string {
	return g.EventName
}

func (g gridUpdate) GetEventTimestampMilli() int64 {
	return g.EventTimestampMilli
}

func (g gridUpdate) GetTimestampMilli() int64 {
	return g.TimestampMilli
}

func (g gridUpdate) GetStrategyID() int64 {
	return g.GridUpdate.StrategyID
}

func (g gridUpdate) GetStrategyType() string {
	return g.GridUpdate.StrategyType
}

func (g gridUpdate) GetStrategyStatus() string {
	return g.GridUpdate.StrategyStatus
}

func (g gridUpdate) GetSymbol() string {
	return g.GridUpdate.Symbol
}

func (g gridUpdate) GetRealizedPNL() decimal.Decimal {
	return g.GridUpdate.RealizedPNL
}

func (g gridUpdate) GetUnMatchedAveragePrice() decimal.Decimal {
	return g.GridUpdate.UnMatchedAveragePrice
}

func (g gridUpdate) GetUnMatchedQuantity() decimal.Decimal {
	return g.GridUpdate.UnMatchedQuantity
}

func (g gridUpdate) GetUnMatchedFee() decimal.Decimal {
	return g.GridUpdate.UnMatchedFee
}

func (g gridUpdate) GetMatchedPNL() decimal.Decimal {
	return g.GridUpdate.MatchedPNL
}

func (g gridUpdate) GetUpdateTimestampMilli() int64 {
	return g.GridUpdate.UpdateTimestampMilli
}
