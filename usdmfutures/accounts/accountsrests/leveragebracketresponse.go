package accountsrests

import "github.com/shopspring/decimal"

type LeverageBracketData struct {
	Bracket          int64           `json:"bracket"`          // 层级
	InitialLeverage  int64           `json:"initialLeverage"`  // 该层允许的最高初始杠杆倍数
	NotionalCap      decimal.Decimal `json:"notionalCap"`      // 该层对应的名义价值上限
	NotionalFloor    decimal.Decimal `json:"notionalFloor"`    // 该层对应的名义价值下限
	MaintMarginRatio decimal.Decimal `json:"maintMarginRatio"` // 该层对应的维持保证金率
	Cum              int64           `json:"cum"`              // 速算数
}

type LeverageBracketDatas []LeverageBracketData

type LeverageBracketResponseItem struct {
	Symbol       string               `json:"symbol"`
	NotionalCoef decimal.Decimal      `json:"notionalCoef"` // 用户bracket相对默认bracket的倍数，仅在和交易对默认不一样时显示
	Brackets     LeverageBracketDatas `json:"brackets"`
}

type LeverageBracketResponse []LeverageBracketResponseItem
