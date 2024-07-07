package accountspushs

import "github.com/shopspring/decimal"

type AccountEventPosition interface {
	GetSymbol() string
	GetPositionAmount() decimal.Decimal
	GetEntryPrice() decimal.Decimal
	GetBalanceEntryPrice() decimal.Decimal
	GetCumulativeRealized() decimal.Decimal
	GetUnrealizedProfit() decimal.Decimal
	GetMarginMode() string
	GetPositionMargin() decimal.Decimal
	GetPosition() string
}

type DataAccountEventPosition struct {
	Symbol             string          `json:"s"`   // 交易对
	PositionAmount     decimal.Decimal `json:"pa"`  // 仓位
	EntryPrice         decimal.Decimal `json:"ep"`  // 入倉價格
	BalanceEntryPrice  decimal.Decimal `json:"bep"` // 盈亏平衡价
	CumulativeRealized decimal.Decimal `json:"cr"`  // (費錢)累计实现损益
	UnrealizedProfit   decimal.Decimal `json:"up"`  // 持倉未实现盈亏
	MarginMode         string          `json:"mt"`  // 保证金模式
	PositionMargin     decimal.Decimal `json:"iw"`  // 仓位保证金
	Position           string          `json:"ps"`  // 仓位方向
}

func (d DataAccountEventPosition) GetSymbol() string {
	return d.Symbol
}

func (d DataAccountEventPosition) GetPositionAmount() decimal.Decimal {
	return d.PositionAmount
}

func (d DataAccountEventPosition) GetEntryPrice() decimal.Decimal {
	return d.EntryPrice
}

func (d DataAccountEventPosition) GetBalanceEntryPrice() decimal.Decimal {
	return d.BalanceEntryPrice
}

func (d DataAccountEventPosition) GetCumulativeRealized() decimal.Decimal {
	return d.CumulativeRealized
}

func (d DataAccountEventPosition) GetUnrealizedProfit() decimal.Decimal {
	return d.UnrealizedProfit
}

func (d DataAccountEventPosition) GetMarginMode() string {
	return d.MarginMode
}

func (d DataAccountEventPosition) GetPositionMargin() decimal.Decimal {
	return d.PositionMargin
}

func (d DataAccountEventPosition) GetPosition() string {
	return d.Position
}

func (d DataAccountEventPosition) GetAccountEventPosition() AccountEventPosition {
	return d
}

type DataAccountEventPositions []DataAccountEventPosition

type AccountEventBalance interface {
	GetAsset() string
	GetWalletBalance() decimal.Decimal
	GetCrossWalletBalance() decimal.Decimal
	GetBalanceChange() decimal.Decimal
}

type DataAccountEventBalance struct {
	Asset              string          `json:"a"`  // 资产
	WalletBalance      decimal.Decimal `json:"wb"` // 钱包余额
	CrossWalletBalance decimal.Decimal `json:"cw"` // 除去逐仓仓位保证金的钱包余额
	BalanceChange      decimal.Decimal `json:"bc"` // 除去盈亏与交易手续费以外的钱包余额改变量
}

func (d DataAccountEventBalance) GetAsset() string {
	return d.Asset
}

func (d DataAccountEventBalance) GetWalletBalance() decimal.Decimal {
	return d.WalletBalance
}

func (d DataAccountEventBalance) GetCrossWalletBalance() decimal.Decimal {
	return d.CrossWalletBalance
}

func (d DataAccountEventBalance) GetBalanceChange() decimal.Decimal {
	return d.BalanceChange
}

type DataAccountEventBalances []DataAccountEventBalance

type AccountEvent interface {
	GetAccountEventName() string
	GetBalances() []DataAccountEventBalance
	GetPositions() []DataAccountEventPosition
}

type DataAccountEvent struct {
	AccountEventName string                    `json:"m"` // 事件推出原因
	Balances         DataAccountEventBalances  `json:"B"` // 账户余额
	Positions        DataAccountEventPositions `json:"P"` // 仓位
}

func (d DataAccountEvent) GetAccountEventName() string {
	return d.AccountEventName
}

func (d DataAccountEvent) GetBalances() []DataAccountEventBalance {
	return d.Balances
}

func (d DataAccountEvent) GetPositions() []DataAccountEventPosition {
	return d.Positions
}

type DataAccountUpdate struct {
	DataEvent
	Timestamp    int64            `json:"T"`
	AccountEvent DataAccountEvent `json:"a"` // 账户更新事件
}

func (d DataAccountUpdate) GetTimestamp() int64 {
	return d.Timestamp
}

func (d DataAccountUpdate) GetAccountEventName() string {
	return d.AccountEvent.GetAccountEventName()
}

func (d DataAccountUpdate) GetBalances() DataAccountEventBalances {
	return d.AccountEvent.GetBalances()
}

func (d DataAccountUpdate) GetPositions() DataAccountEventPositions {
	return d.AccountEvent.GetPositions()
}

type AccountPosition interface {
	GetEventName() string
	GetEventTs() int64
	GetTimestamp() int64
	GetAccountEventName() string
	GetPositions() DataAccountEventPositions
}

type AccountBalance interface {
	GetEventName() string
	GetEventTs() int64
	GetTimestamp() int64
	GetAccountEventName() string
	GetBalances() DataAccountEventBalances
}
