package accountspushs

import "github.com/shopspring/decimal"

type DataGridUpdate struct {
	DataEvent
	Timestamp  int64 `json:"T"` // 事件時間
	GridUpdate struct {
		StrategyID          int64           `json:"si"` // 策略ID
		StrategyType        string          `json:"st"` // 策略類型
		StrategyStatus      string          `json:"ss"` // 策略狀態
		Symbol              string          `json:"s"`  // 交易對
		RealizedProfit      decimal.Decimal `json:"r"`  // 已實現PNL
		UnmatchAveragePrice decimal.Decimal `json:"up"` // 未配對均價
		UnmatchQuatity      decimal.Decimal `json:"uq"` // 未配對數量
		UnmatchFee          decimal.Decimal `json:"uf"` // 未配對手續費
		MatchPrice          decimal.Decimal `json:"mp"` // 已配對PNL
		UpdateTs            int64           `json:"ut"` // 更新時間
	} `json:"gu"`
}

func (d DataGridUpdate) GetTimestamp() int64 {
	return d.Timestamp
}

func (d DataGridUpdate) GetStrategyID() int64 {
	return d.GridUpdate.StrategyID
}

func (d DataGridUpdate) GetStrategyType() string {
	return d.GridUpdate.StrategyType
}

func (d DataGridUpdate) GetStrategyStatus() string {
	return d.GridUpdate.StrategyStatus
}

func (d DataGridUpdate) GetSymbol() string {
	return d.GridUpdate.Symbol
}

func (d DataGridUpdate) GetRealizedProfit() decimal.Decimal {
	return d.GridUpdate.RealizedProfit
}

func (d DataGridUpdate) GetUnmatchAveragePrice() decimal.Decimal {
	return d.GridUpdate.UnmatchAveragePrice
}

func (d DataGridUpdate) GetUnmatchQuatity() decimal.Decimal {
	return d.GridUpdate.UnmatchQuatity
}

func (d DataGridUpdate) GetUnmatchFee() decimal.Decimal {
	return d.GridUpdate.UnmatchFee
}

func (d DataGridUpdate) GetMatchPrice() decimal.Decimal {
	return d.GridUpdate.MatchPrice
}

func (d DataGridUpdate) GetUpdateTimestamp() int64 {
	return d.GridUpdate.UpdateTs
}

type GridUpdate interface {
	GetEventName() string
	GetEventTs() int64
	GetTimestamp() int64
	GetStrategyID() int64
	GetStrategyType() string
	GetStrategyStatus() string
	GetSymbol() string
	GetRealizedProfit() decimal.Decimal
	GetUnmatchAveragePrice() decimal.Decimal
	GetUnmatchQuatity() decimal.Decimal
	GetUnmatchFee() decimal.Decimal
	GetMatchPrice() decimal.Decimal
	GetUpdateTimestamp() int64
}
