package accountspushs

type DataAccountConfigUpdate struct {
	DataEvent
	Timestamp     int64
	AccountConfig struct {
		Symbol   string `json:"s"` // 交易對
		Leverage int64  `json:"l"` // 杠桿倍數
	} `json:"ac"`
}

func (d DataAccountConfigUpdate) GetTimestamp() int64 {
	return d.Timestamp
}

func (d DataAccountConfigUpdate) GetSymbol() string {
	return d.AccountConfig.Symbol
}

func (d DataAccountConfigUpdate) GetLeverage() int64 {
	return d.AccountConfig.Leverage
}

type AccountConfigUpdate interface {
	GetEventName() string
	GetEventTs() int64
	GetTimestamp() int64
	GetSymbol() string
	GetLeverage() int64
}
