package gobinanceenums

const (

	// 交易對類型
	SYMBOL_TYPE_FUTURE string = "FUTURE" // 期貨

	// 合約類型
	CONTRACT_TYPE_PERPETUAL            string = "PERPETUAL"            // 永續合約
	CONTRACT_TYPE_CURRENT_MONTH        string = "CURRENT_MONTH"        // 當月交割合約
	CONTRACT_TYPE_NEXT_MONTH           string = "NEXT_MONTH"           // 次月交割合約
	CONTRACT_TYPE_CURRENT_QUARTER      string = "CURRENT_QUARTER"      // 當季交割合約
	CONTRACT_TYPE_NEXT_QUARTER         string = "NEXT_QUARTER"         // 次季交割合約
	CONTRACT_TYPE_PERPETUAL_DELIVERING string = "PERPETUAL_DELIVERING" // 交割結算中合約

	// 合約狀態
	CONTRACT_STATUS_PENDING_TRADING string = "PENDING_TRADING" // 待上市
	CONTRACT_STATUS_TRADING         string = "TRADING"         // 交易中
	CONTRACT_STATUS_PRE_DELIVERING  string = "PRE_DELIVERING"  // 预交割
	CONTRACT_STATUS_DELIVERING      string = "DELIVERING"      // 交割中
	CONTRACT_STATUS_DELIVERED       string = "DELIVERED"       // 已交割
	CONTRACT_STATUS_PRE_SETTLE      string = "PRE_SETTLE"      // 预结算
	CONTRACT_STATUS_SETTLING        string = "SETTLING"        // 结算中
	CONTRACT_STATUS_CLOSE           string = "CLOSE"           // 已下架

	// 訂單狀態
	STATUS_NEW              string = "NEW"              // 新建订单
	STATUS_PARTIALLY_FILLED string = "PARTIALLY_FILLED" // 部分成交
	STATUS_FILLED           string = "FILLED"           // 全部成交
	STATUS_CANCELED         string = "CANCELED"         // 已撤销
	STATUS_REJECTED         string = "REJECTED"         // 订单被拒绝
	STATUS_EXPIRED          string = "EXPIRED"          // 订单过期(根据timeInForce参数规则)

	// 訂單種類
	ORDER_TYPE_LIMIT                string = "LIMIT"                // 限价单
	ORDER_TYPE_MARKET               string = "MARKET"               // 市价单
	ORDER_TYPE_STOP                 string = "STOP"                 // 止损限价单
	ORDER_TYPE_STOP_MARKET          string = "STOP_MARKET"          // 止损市价单
	ORDER_TYPE_TAKE_PROFIT          string = "TAKE_PROFIT"          // 止盈限价单
	ORDER_TYPE_TAKE_PROFIT_MARKET   string = "TAKE_PROFIT_MARKET"   // 止盈市价单
	ORDER_TYPE_TRAILING_STOP_MARKET string = "TRAILING_STOP_MARKET" // 跟踪止损单

	// 訂單方向
	SIDE_BUY  = "BUY"  // 買入
	SIDE_SELL = "SELL" // 賣出

	// 持倉方向
	POSITION_SIDE_BOTH  string = "BOTH"  // 单一持仓方向
	POSITION_SIDE_LONG  string = "LONG"  // 多头(双向持仓下)
	POSITION_SIDE_SHORT string = "SHORT" // 空头(双向持仓下)

	// 有效方式
	TIME_IN_FORCE_GTC string = "GTC" // Good Till Cancel 成交为止（下单后仅有1年有效期，1年后自动取消）
	TIME_IN_FORCE_IOC string = "IOC" // Immediate or Cancel 无法立即成交(吃单)的部分就撤销
	TIME_IN_FORCE_FOK string = "FOK" // Fill or Kill 无法全部立即成交就撤销
	TIME_IN_FORCE_GTX string = "GTX" // Good Till Crossing 无法成为挂单方就撤销
	TIME_IN_FORCE_GTD string = "GTD" // Good Till Date 在特定时间之前有效，到期自动撤销

	// 條件價格觸發類型
	WORKING_TYPE_MARK_PRICE     string = "MARK_PRICE"     // 標記價格
	WORKING_TYPE_CONTRACT_PRICE string = "CONTRACT_PRICE" // 合約價格

	// 响应类型 (newOrderRespType)
	NEW_ORDER_RESP_TYPE_ACK    string = "ACK"
	NEW_ORDER_RESP_TYPE_RESULT string = "RESULT"

	// K线间隔 (m -> 分钟; h -> 小时; d -> 天; w -> 周; M -> 月)
	INTERVAL_ONE_MINUTE      string = "1m"
	INTERVAL_THREE_MINUTES   string = "3m"
	INTERVAL_FIVE_MINUTES    string = "5m"
	INTERVAL_FIFTEEN_MINUTES string = "15m"
	INTERVAL_THIRTY_MINUTES  string = "30m"
	INTERVAL_ONE_HOUR        string = "1h"
	INTERVAL_TWO_HOURS       string = "2h"
	INTERVAL_FOUR_HOURS      string = "4h"
	INTERVAL_SIX_HOURS       string = "6h"
	INTERVAL_EIGHT_HOURS     string = "8h"
	INTERVAL_TWELVE_HOURS    string = "12h"
	INTERVAL_ONE_DAY         string = "1d"
	INTERVAL_THREE_DAYS      string = "3d"
	INTERVAL_ONE_WEEK        string = "1w"
	INTERVAL_ONE_MONTH       string = "1M"

	// 防止自成交模式:
	STP_MODE_NONE         string = "NONE"
	STP_MODE_EXPIRE_TAKER string = "EXPIRE_TAKER"
	STP_MODE_EXPIRE_BOTH  string = "EXPIRE_BOTH"
	STP_MODE_EXPIRE_MAKER string = "EXPIRE_MAKER"

	// 盘口价下单模式:
	PRICE_MATCH_OPPONENT        string = "OPPONENT"    // 盘口对手价
	PRICE_MATCH_OPPONENT_FIVE   string = "OPPONENT_5"  // 盘口对手5档价
	PRICE_MATCH_OPPONENT_TEN    string = "OPPONENT_10" // 盘口对手10档价
	PRICE_MATCH_OPPONENT_TWENTY string = "OPPONENT_20" // 盘口对手20档价
	PRICE_MATCH_QUEUE           string = "QUEUE"       // 盘口同向价
	PRICE_MATCH_QUEUE_FIVE      string = "QUEUE_5"     // 盘口同向排队5档价
	PRICE_MATCH_QUEUE_TEN       string = "QUEUE_10"    // 盘口同向排队10档价
	PRICE_MATCH_QUEUE_TWENTY    string = "QUEUE_20"    // 盘口同向排队20档价
)
