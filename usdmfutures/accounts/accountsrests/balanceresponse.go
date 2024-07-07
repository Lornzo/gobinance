package accountsrests

import "github.com/shopspring/decimal"

type BalanceResponseItem struct {
	AccountAlias       string          `json:"accountAlias"`       // 账户唯一识别码
	Asset              string          `json:"asset"`              // 资产
	Balance            decimal.Decimal `json:"balance"`            // 总余额
	CrossWalletBalance decimal.Decimal `json:"crossWalletBalance"` // 全仓余额
	CrossUnPnl         decimal.Decimal `json:"crossUnPnl"`         // 全仓持仓未实现盈亏
	AvailableBalance   decimal.Decimal `json:"availableBalance"`   // 下单可用余额
	MaxWithdrawAmount  decimal.Decimal `json:"maxWithdrawAmount"`  // 最大可转出余额
	MarginAvailable    bool            `json:"marginAvailable"`    // 是否可用作联合保证金
	UpdateTime         int64           `json:"updateTime"`         // 更新时间
}

type BalanceResponse []BalanceResponseItem
