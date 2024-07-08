package marketdatasrests

import "github.com/shopspring/decimal"

type KLine struct {
	OpenTimestamp   int64           `json:"open_ts"`           // 開盤時間(milliseconds)
	OpenPrice       decimal.Decimal `json:"open_price"`        // 開盤價
	HighPrice       decimal.Decimal `json:"high_price"`        // 最高價
	LowPrice        decimal.Decimal `json:"low_price"`         // 最低價
	ClosePrice      decimal.Decimal `json:"close_price"`       // 收盤價
	Volume          decimal.Decimal `json:"volume"`            // 成交量
	CloseTimestamp  int64           `json:"close_ts"`          // 收盤時間(milliseconds)
	Turnover        decimal.Decimal `json:"turnover"`          // 成交额
	TurnoverNum     int64           `json:"turnover_num"`      // 成交筆數
	ActiveBuyVolume decimal.Decimal `json:"active_buy_volume"` // 主動買入成交量
	ActiveBuyTurn   decimal.Decimal `json:"active_buy_turn"`   // 主動買入成交额
	UnknowNum       decimal.Decimal `json:"unknow_num"`        // 未知
}

type KLinesResponse []KLine
