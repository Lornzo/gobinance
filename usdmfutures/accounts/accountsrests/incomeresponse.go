package accountsrests

import "github.com/shopspring/decimal"

type IncomeResponseItem struct {
	Symbol     string          `json:"symbol"`     // 交易对，仅针对涉及交易对的资金流
	IncomeType string          `json:"incomeType"` // 资金流类型
	Income     decimal.Decimal `json:"income"`     // 资金流数量，正数代表流入，负数代表流出
	Asset      string          `json:"asset"`      // 资产内容
	Info       string          `json:"info"`       // 备注信息，取决于流水类型
	Timestamp  int64           `json:"time"`       // 时间
	TranID     string          `json:"tranId"`     // 划转ID
	TradeID    string          `json:"tradeId"`    // 引起流水产生的原始交易ID
}

type IncomeResponse []IncomeResponseItem
