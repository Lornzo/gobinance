package accountsrests

import "github.com/shopspring/decimal"

type AccountAsset struct {
	Asset                  string          `json:"asset"`                  //资产
	WalletBalance          decimal.Decimal `json:"walletBalance"`          //余额
	UnrealizedProfit       decimal.Decimal `json:"unrealizedProfit"`       // 未实现盈亏
	MarginBalance          decimal.Decimal `json:"marginBalance"`          // 保证金余额
	MaintMargin            decimal.Decimal `json:"maintMargin"`            // 维持保证金
	InitialMargin          decimal.Decimal `json:"initialMargin"`          // 当前所需起始保证金
	PositionInitialMargin  decimal.Decimal `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
	OpenOrderInitialMargin decimal.Decimal `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
	CrossWalletBalance     decimal.Decimal `json:"crossWalletBalance"`     //全仓账户余额
	CrossUnPnl             decimal.Decimal `json:"crossUnPnl"`             // 全仓持仓未实现盈亏
	AvailableBalance       decimal.Decimal `json:"availableBalance"`       // 可用余额
	MaxWithdrawAmount      decimal.Decimal `json:"maxWithdrawAmount"`      // 最大可转出余额
	MarginAvailable        bool            `json:"marginAvailable"`        // 是否可用作联合保证金
	UpdateTime             int64           `json:"updateTime"`             //更新时间
}

func (a AccountAsset) GetAssetName() string {
	return a.Asset
}

type AccountAssets []AccountAsset

func (a AccountAssets) GetAssetByAssetName(asset string) (AccountAsset, bool) {
	for _, v := range a {
		if v.GetAssetName() == asset {
			return v, true
		}
	}
	return AccountAsset{}, false
}

type AccountPosition struct {
	Symbol                 string          `json:"symbol"`                 // 交易对
	InitialMargin          decimal.Decimal `json:"initialMargin"`          // 当前所需起始保证金(基于最新标记价格)
	MaintMargin            decimal.Decimal `json:"maintMargin"`            //维持保证金
	UnrealizedProfit       decimal.Decimal `json:"unrealizedProfit"`       // 持仓未实现盈亏
	PositionInitialMargin  decimal.Decimal `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
	OpenOrderInitialMargin decimal.Decimal `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
	Leverage               decimal.Decimal `json:"leverage"`               // 杠杆倍率
	Isolated               bool            `json:"isolated"`               // 是否是逐仓模式
	EntryPrice             decimal.Decimal `json:"entryPrice"`             // 持仓成本价
	MaxNotional            decimal.Decimal `json:"maxNotional"`            // 当前杠杆下用户可用的最大名义价值
	BidNotional            decimal.Decimal `json:"bidNotional"`            // 买单净值，忽略
	AskNotional            decimal.Decimal `json:"askNotional"`            // 卖单净值，忽略
	PositionSide           string          `json:"positionSide"`           // 持仓方向
	PositionAmt            decimal.Decimal `json:"positionAmt"`            // 持仓数量
	UpdateTime             int64           `json:"updateTime"`             // 更新时间
}

func (a AccountPosition) GetSymbol() string {
	return a.Symbol
}

type AccountPositions []AccountPosition

func (a AccountPositions) GetPositionBySymbol(symbol string) (AccountPosition, bool) {
	for _, v := range a {
		if v.GetSymbol() == symbol {
			return v, true
		}
	}
	return AccountPosition{}, false
}

type AccountResponse struct {
	MultiAssetsMargin           bool             `json:"multiAssetsMargin"`
	TradeGroupId                int64            `json:"tradeGroupId"`
	FeeTier                     int64            `json:"feeTier"`                     // 手续费等级
	FeeBurn                     bool             `json:"feeBurn"`                     // "true": 手续费抵扣开; "false": 手续费抵扣关
	CanTrade                    bool             `json:"canTrade"`                    // 是否可以交易
	CanDeposit                  bool             `json:"canDeposit"`                  // 是否可以入金
	CanWithdraw                 bool             `json:"canWithdraw"`                 // 是否可以出金
	UpdateTime                  int64            `json:"updateTime"`                  // 保留字段，请忽略
	TotalInitialMargin          decimal.Decimal  `json:"totalInitialMargin"`          // 当前所需起始保证金总额(存在逐仓请忽略), 仅计算usdt资产
	TotalMaintMargin            decimal.Decimal  `json:"totalMaintMargin"`            // 维持保证金总额, 仅计算usdt资产
	TotalWalletBalance          decimal.Decimal  `json:"totalWalletBalance"`          // 账户总余额, 仅计算usdt资产
	TotalUnrealizedProfit       decimal.Decimal  `json:"totalUnrealizedProfit"`       // 持仓未实现盈亏总额, 仅计算usdt资产
	TotalMarginBalance          decimal.Decimal  `json:"totalMarginBalance"`          // 保证金总余额, 仅计算usdt资产
	TotalPositionInitialMargin  decimal.Decimal  `json:"totalPositionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格), 仅计算usdt资产
	TotalOpenOrderInitialMargin decimal.Decimal  `json:"totalOpenOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格), 仅计算usdt资产
	TotalCrossWalletBalance     decimal.Decimal  `json:"totalCrossWalletBalance"`     // 全仓账户余额, 仅计算usdt资产
	TotalCrossUnPnl             decimal.Decimal  `json:"totalCrossUnPnl"`             // 全仓持仓未实现盈亏总额, 仅计算usdt资产
	AvailableBalance            decimal.Decimal  `json:"availableBalance"`            // 可用余额, 仅计算usdt资产
	MaxWithdrawAmount           decimal.Decimal  `json:"maxWithdrawAmount"`           // 最大可转出余额, 仅计算usdt资产
	Assets                      AccountAssets    `json:"assets"`                      // 資產
	Positions                   AccountPositions `json:"positions"`                   // 头寸，将返回所有市场symbol。根据用户持仓模式展示持仓方向，即单向模式下只返回BOTH持仓情况，双向模式下只返回 LONG 和 SHORT 持仓情况。
}
