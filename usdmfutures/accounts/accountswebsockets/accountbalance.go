package accountswebsockets

import "github.com/shopspring/decimal"

type AccountBalance struct {
	Asset              string          `json:"asset"`
	Balance            decimal.Decimal `json:"balance"`
	CrossWalletBalance decimal.Decimal `json:"crossWalletBalance"`
	CrossUnPNL         decimal.Decimal `json:"crossunPnl"`
	AvailableBalance   decimal.Decimal `json:"availableBalance"`
	MaxWithdrawAmount  decimal.Decimal `json:"maxWithdrawAmount"`
}

type AccountBalances []AccountBalance
