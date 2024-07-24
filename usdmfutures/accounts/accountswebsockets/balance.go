package accountswebsockets

import "github.com/shopspring/decimal"

type Balance struct {
	Asset              string          `json:"a"`
	WalletBalance      decimal.Decimal `json:"wb"`
	CrossWalletBalance decimal.Decimal `json:"cw"`
	BalanceChange      decimal.Decimal `json:"bc"`
}
