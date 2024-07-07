package accountspushs

type DataStrategyUpdate struct {
	DataEvent
	Timestamp      int64 `json:"T"` // 事件時間
	StrategyUpdate struct {
		StrategyID     int64  `json:"si"` // 策略ID
		StrategyType   string `json:"st"` // 策略類型
		StrategyStatus string `json:"ss"` // 策略狀態
		Symbol         string `json:"s"`  // 交易對
		UpdateTs       int64  `json:"ut"` // 更新時間
		Code           int64  `json:"c"`  // op code
	} `json:"su"`
}

func (d DataStrategyUpdate) GetTimestamp() int64 {
	return d.Timestamp
}

func (d DataStrategyUpdate) GetStrategyID() int64 {
	return d.StrategyUpdate.StrategyID
}

func (d DataStrategyUpdate) GetStrategyType() string {
	return d.StrategyUpdate.StrategyType
}

func (d DataStrategyUpdate) GetStrategyStatus() string {
	return d.StrategyUpdate.StrategyStatus
}

func (d DataStrategyUpdate) GetStrategySymbol() string {
	return d.StrategyUpdate.Symbol
}

func (d DataStrategyUpdate) GetUpdateTimestamp() int64 {
	return d.StrategyUpdate.UpdateTs
}

func (d DataStrategyUpdate) GetStrategyOpCode() int64 {
	return d.StrategyUpdate.Code
}

type StrategyUpdate interface {
	GetTimestamp() int64
	GetStrategyID() int64
	GetStrategyType() string
	GetStrategyStatus() string
	GetStrategySymbol() string
	GetUpdateTimestamp() int64
	GetStrategyOpCode() int64
}
