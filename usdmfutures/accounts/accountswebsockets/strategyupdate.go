package accountswebsockets

type StrategyUpdate interface {
	GetEventName() string
	GetEventTimestampMilli() int64
	GetTimestampMilli() int64
	GetStrategyID() int64
	GetStrategyType() string
	GetStrategyStatus() string
	GetSymbol() string
	GetUpdateTimestampMilli() int64
	GetOprationCode() int64
}

type strategyUpdate struct {
	EventName           string `json:"e"`
	EventTimestampMilli int64  `json:"E"`
	TimestampMilli      int64  `json:"T"`
	StrategyUpdate      struct {
		StrategyID           int64  `json:"si"`
		StrategyType         string `json:"st"`
		StrategyStatus       string `json:"ss"`
		Symbol               string `json:"s"`
		UpdateTimestampMilli int64  `json:"ut"`
		OprationCode         int64  `json:"c"`
	} `json:"su"`
}

func (s strategyUpdate) GetEventName() string {
	return s.EventName
}

func (s strategyUpdate) GetEventTimestampMilli() int64 {
	return s.EventTimestampMilli
}

func (s strategyUpdate) GetTimestampMilli() int64 {
	return s.TimestampMilli
}

func (s strategyUpdate) GetStrategyID() int64 {
	return s.StrategyUpdate.StrategyID
}

func (s strategyUpdate) GetStrategyType() string {
	return s.StrategyUpdate.StrategyType
}

func (s strategyUpdate) GetStrategyStatus() string {
	return s.StrategyUpdate.StrategyStatus
}

func (s strategyUpdate) GetSymbol() string {
	return s.StrategyUpdate.Symbol
}

func (s strategyUpdate) GetUpdateTimestampMilli() int64 {
	return s.StrategyUpdate.UpdateTimestampMilli
}

func (s strategyUpdate) GetOprationCode() int64 {
	return s.StrategyUpdate.OprationCode
}
