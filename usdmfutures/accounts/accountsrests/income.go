package accountsrests

import (
	"context"
	"fmt"

	"github.com/Lornzo/gobinance/commons"
)

func NewIncome(baseUrl string, apiKey string, apiSecret string) *Income {
	var incomeAPI *Income = &Income{
		Base: commons.NewBinanceRestfulBase(),
	}
	incomeAPI.Base.SetBaseURL(baseUrl)
	incomeAPI.Base.SetAPIKeyAndSecret(apiKey, apiSecret)
	return incomeAPI
}

// 获取账户损益资金流水
//   - 如果startTime 和 endTime 均未发送, 只会返回最近7天的数据。
//   - 如果incomeType没有发送，返回所有类型账户损益资金流水。
//   - "trandId" 在相同用户的同一种收益流水类型中是唯一的。
//   - 仅保留最近3个月的数据。
type Income struct {
	Base       commons.RestfulBase
	Symbol     string // 交易对
	IncomeType string // 收益类型
	StartTime  int64  // 起始时间
	EndTime    int64  // 结束时间
	Page       int64  // 分頁數
	Limit      int64  // 返回的结果集数量 默认值:100 最大值:1000
	RecvWindow int64
	Timestamp  int64
}

func (i *Income) GetAPIPathes() []string {
	return []string{"fapi", "v1", "income"}
}

func (i *Income) GetWeights() int {
	return 30
}

func (i *Income) SetSymbol(symbol string) {
	i.Symbol = symbol
}

// 收益类型：
//   - TRANSFER 转账
//   - WELCOME_BONUS 欢迎奖金
//   - REALIZED_PNL 已实现盈亏
//   - FUNDING_FEE 资金费用
//   - COMMISSION 佣金
//   - INSURANCE_CLEAR 强平
//   - REFERRAL_KICKBACK 推荐人返佣
//   - COMMISSION_REBATE 被推荐人返佣
//   - API_REBATE API佣金回扣
//   - CONTEST_REWARD 交易大赛奖金
//   - CROSS_COLLATERAL_TRANSFER cc转账
//   - OPTIONS_PREMIUM_FEE 期权购置手续费
//   - OPTIONS_SETTLE_PROFIT 期权行权收益
//   - INTERNAL_TRANSFER 内部账户，给普通用户划转
//   - AUTO_EXCHANGE 自动兑换
//   - DELIVERED_SETTELMENT 下架结算
//   - COIN_SWAP_DEPOSIT 闪兑转入
//   - COIN_SWAP_WITHDRAW 闪兑转出
//   - POSITION_LIMIT_INCREASE_FEE 仓位限制上调费用
func (i *Income) SetIncomeType(incomeType string) {
	i.IncomeType = incomeType
}

func (i *Income) SetStartTime(startTs int64) {
	i.StartTime = startTs
}

func (i *Income) SetEndTime(endTs int64) {
	i.EndTime = endTs
}

func (i *Income) SetPage(page int64) {
	i.Page = page
}

func (i *Income) SetLimit(limit int64) {
	i.Limit = limit
}

func (i *Income) SetRecvWindow(recvWindow int64) {
	i.RecvWindow = recvWindow
}

func (i *Income) SetTimestamp(ts int64) {
	i.Timestamp = ts
}

func (i *Income) DoRequest(ctx context.Context) (IncomeResponse, error) {
	var (
		resp IncomeResponse
		err  error
	)

	i.initApiUrl()

	if _, err = i.Base.GET(ctx, &resp); err != nil {
		return IncomeResponse{}, err
	}

	return resp, nil
}

func (i *Income) initApiUrl() {
	i.Base.SetPathes(i.GetAPIPathes()...)

	if i.Symbol != "" {
		i.Base.SetQuery("symbol", i.Symbol)
	}

	if i.IncomeType != "" {
		i.Base.SetQuery("incomeType", i.IncomeType)
	}

	if i.StartTime > 0 {
		i.Base.SetQuery("startTime", fmt.Sprint(i.StartTime))
	}

	if i.EndTime > 0 {
		i.Base.SetQuery("endTime", fmt.Sprint(i.EndTime))
	}

	if i.Page > 0 {
		i.Base.SetQuery("page", fmt.Sprint(i.Page))
	}

	if i.Limit > 0 {
		i.Base.SetQuery("limit", fmt.Sprint(i.Limit))
	}

	if i.RecvWindow > 0 {
		i.Base.SetQuery("recvWindow", fmt.Sprint(i.RecvWindow))
	}

	if i.Timestamp > 0 {
		i.Base.SetQuery("timestamp", fmt.Sprint(i.Timestamp))
	}

	i.Base.UseSignature(true)
}
