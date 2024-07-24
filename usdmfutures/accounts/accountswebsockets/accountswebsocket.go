package accountswebsockets

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Lornzo/gobinance/binancewebsockets"
	"github.com/Lornzo/gobinance/channels"
	"github.com/Lornzo/gobinance/threadsafetypes"
	"github.com/Lornzo/gobinance/usdmfutures/accounts/accountsrests"
)

type accountsWebsocket struct {
	concreateRest                            restful
	ws                                       binancewebsockets.Websocket
	account                                  threadsafetypes.Account
	listenKey                                threadsafetypes.String
	bytesChannel                             channels.BytesChannel
	errorSubscribers                         errorSubscribers
	listenKeyExpiredSubscribers              listenKeyExpiredSubscribers
	accountUpdateSubscribers                 accountUpdateSubscribers
	marginCallSubscribers                    marginCallSubscribers
	orderTradeUpdateSubscribers              orderTradeUpdateSubscribers
	accountConfigUpdateSubscribers           accountConfigUpdateSubscribers
	accountInfoUpdateSubscribers             accountInfoUpdateSubscribers
	strategyUpdateSubscribers                strategyUpdateSubscribers
	gridUpdateSubscribers                    gridUpdateSubscribers
	conditionalOrderTriggerRejectSubscribers conditionalOrderTriggerRejectSubscribers
	rateLimitsSubscribers                    rateLimitsSubscribers
}

func (a *accountsWebsocket) CreateListenKey(ctx context.Context) (string, error) {
	var (
		resp accountsrests.ListenKeyResponse
		err  error
	)

	if resp, err = a.concreateRest.ListenKeyCreate().DoRequest(ctx); err != nil {
		return "", err
	}

	return resp.ListenKey, nil
}

func (a *accountsWebsocket) ExtendListenKey(ctx context.Context) (string, error) {

	var (
		resp accountsrests.ListenKeyResponse
		err  error
	)

	if resp, err = a.concreateRest.ListenKeyUpdate().DoRequest(ctx); err != nil {
		return "", err
	}

	return resp.ListenKey, nil

}

func (a *accountsWebsocket) CloseListenKey(ctx context.Context) (string, error) {

	var err error = a.concreateRest.ListenKeyDelete().DoRequest(ctx)

	if err != nil {
		return "", err
	}

	return a.listenKey.Get(), nil

}

func (a *accountsWebsocket) RequestUserPosition(ctx context.Context) (AccountPositions, error) {

	var (
		requestName string = fmt.Sprint(a.listenKey.Get(), "@", "position")
		bytes       []byte
		err         error
		resp        struct {
			RequestName string `json:"req"`
			Result      struct {
				Positions AccountPositions `json:"positions"`
			} `json:"res"`
		}
	)

	if bytes, err = a.request(ctx, requestName); err != nil {
		return AccountPositions{}, err
	}

	if err = json.Unmarshal(bytes, &resp); err != nil {
		return AccountPositions{}, err
	}

	return resp.Result.Positions, nil

}

func (a *accountsWebsocket) RequestBalance(ctx context.Context) (AccountBalances, error) {

	var (
		requestName string = fmt.Sprint(a.listenKey.Get(), "@", "balance")
		bytes       []byte
		err         error
		resp        struct {
			RequestName string `json:"req"`
			Result      struct {
				AccountAlias string          `json:"accountAlias"`
				Balances     AccountBalances `json:"balances"`
			} `json:"res"`
		}
	)

	if bytes, err = a.request(ctx, requestName); err != nil {
		return nil, nil
	}

	if err = json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}

	return resp.Result.Balances, nil
}

func (a *accountsWebsocket) RequestAccountInformation(ctx context.Context) (AccountInformation, error) {

	var (
		requestName string = fmt.Sprint(a.listenKey.Get(), "@", "account")
		bytes       []byte
		err         error
		resp        struct {
			RequestName string             `json:"req"`
			Result      AccountInformation `json:"res"`
		}
	)

	if bytes, err = a.request(ctx, requestName); err != nil {
		return AccountInformation{}, err
	}

	if err = json.Unmarshal(bytes, &resp); err != nil {
		return AccountInformation{}, err
	}

	return resp.Result, nil

}

func (a *accountsWebsocket) requestBytes(ctx context.Context, id interface{}, method string, params interface{}) ([]byte, error) {

	type response struct {
		bytes []byte
		err   error
	}

	var (
		err     error
		request struct {
			ID     interface{} `json:"id"`
			Method string      `json:"method"`
			Params interface{} `json:"params"`
		}
		resp chan response = make(chan response)
	)

	defer close(resp)

	request.ID = id
	request.Method = method
	request.Params = params

	if err = a.bytesChannel.CreateChannel(request.ID); err != nil {
		return nil, err
	}

	defer a.bytesChannel.CloseChannel(request.ID)

	go func() {

		var (
			bytes      []byte
			channelErr error
		)

		bytes, channelErr = a.bytesChannel.ReadChannel(request.ID)

		resp <- response{
			bytes: bytes,
			err:   channelErr,
		}

	}()

	if err = a.ws.WriteJSON(request); err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("context timeout")
	case data := <-resp:
		return data.bytes, data.err
	}
}

func (a *accountsWebsocket) request(ctx context.Context, names ...string) ([]byte, error) {

	if len(names) == 0 {
		return nil, errors.New("request names is empty")
	}

	var (
		bytes []byte
		err   error
		data  struct {
			ID     interface{} `json:"id"`
			Result []struct {
				RequestName string      `json:"req"`
				Result      interface{} `json:"res"`
			} `json:"result"`
		}
	)

	if bytes, err = a.requestBytes(ctx, a.bytesChannel.GetIntID(), "REQUEST", names); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	if len(data.Result) == 0 {
		return nil, errors.New("result not found")
	}

	if len(data.Result) == 1 {
		return json.Marshal(data.Result[0])
	}

	return json.Marshal(data.Result)

}

func (a *accountsWebsocket) SubscribeRateLimits(subscriber RateLimitsSubscriber) error {
	return a.rateLimitsSubscribers.subscribe(subscriber)
}

func (a *accountsWebsocket) UnSubscribeRateLimits(subscriber RateLimitsSubscriber) error {
	return a.rateLimitsSubscribers.unsubscribe(subscriber)
}

func (a *accountsWebsocket) SubscribeError(subscriber ErrorSubscriber) error {
	return a.errorSubscribers.subscribe(subscriber)
}

func (a *accountsWebsocket) UnSubscribeError(subscriber ErrorSubscriber) error {
	return a.errorSubscribers.unsubscribe(subscriber)
}

// listen key 過期
func (a *accountsWebsocket) SubscribeListenKeyExpired(subscriber ListenKeyExpiredSubscriber) error {
	return a.listenKeyExpiredSubscribers.subscribe(subscriber)
}

func (a *accountsWebsocket) UnSubscribeListenKeyExpired(subscriber ListenKeyExpiredSubscriber) error {
	return a.listenKeyExpiredSubscribers.unsubscribe(subscriber)
}

// Account Balance and Position Update
func (a *accountsWebsocket) SubscribeAccountUpdate(subscriber AccountUpdateSubscriber) error {
	return a.accountUpdateSubscribers.subscribe(subscriber)
}

func (a *accountsWebsocket) UnSubscribeAccountUpdate(subscriber AccountUpdateSubscriber) error {
	return a.accountUpdateSubscribers.unsubscribe(subscriber)
}

// 追加保證金
func (a *accountsWebsocket) SubscribeMarginCall(subscriber MarginCallSubscriber) error {
	return a.marginCallSubscribers.subscribe(subscriber)
}

func (a *accountsWebsocket) UnSubscribeMarginCall(subscriber MarginCallSubscriber) error {
	return a.marginCallSubscribers.unsubscribe(subscriber)
}

// 訂單交易更新
func (a *accountsWebsocket) SubscribeOrderUpdate(subscriber OrderTradeUpdateSubscriber) error {
	return a.orderTradeUpdateSubscribers.subscribe(subscriber)
}

func (a *accountsWebsocket) UnSubscribeOrderUpdate(subscriber OrderTradeUpdateSubscriber) error {
	return a.orderTradeUpdateSubscribers.unsubscribe(subscriber)
}

// 杠杆倍数 更新推送
func (a *accountsWebsocket) SubscribeAccountConfigUpdate(subscriber AccountConfigUpdateSubscriber) error {
	return a.accountConfigUpdateSubscribers.subscribe(subscriber)
}

func (a *accountsWebsocket) UnSubscribeAccountConfigUpdate(subscriber AccountConfigUpdateSubscriber) error {
	return a.accountConfigUpdateSubscribers.unsubscribe(subscriber)
}

// 账户配置
func (a *accountsWebsocket) SubscribeAccountInfoUpdate(subscriber AccountInfoUpdateSubscriber) error {
	return a.accountInfoUpdateSubscribers.subscribe(subscriber)
}

func (a *accountsWebsocket) UnSubscribeAccountInfoUpdate(subscriber AccountInfoUpdateSubscriber) error {
	return a.accountInfoUpdateSubscribers.unsubscribe(subscriber)
}

// 策略交易更新推送
func (a *accountsWebsocket) SubscribeStrategyUpdate(subscriber StrategyUpdateSubscriber) error {
	return a.strategyUpdateSubscribers.subscribe(subscriber)
}

func (a *accountsWebsocket) UnSubscribeStrategyUpdate(subscriber StrategyUpdateSubscriber) error {
	return a.strategyUpdateSubscribers.unsubscribe(subscriber)
}

// 网格交易更新推送
func (a *accountsWebsocket) SubscribeGridUpdate(subscriber GridUpdateSubscriber) error {
	return a.gridUpdateSubscribers.subscribe(subscriber)
}

func (a *accountsWebsocket) UnSubscribeGridUpdate(subscriber GridUpdateSubscriber) error {
	return a.gridUpdateSubscribers.unsubscribe(subscriber)
}

// 条件订单(TP/SL)触发后拒绝更新推送
func (a *accountsWebsocket) SubscribeConditionalOrderTriggerReject(subscriber ConditionalOrderTriggerRejectSubscriber) error {
	return a.conditionalOrderTriggerRejectSubscribers.subscribe(subscriber)
}

func (a *accountsWebsocket) UnSubscribeConditionalOrderTriggerReject(subscriber ConditionalOrderTriggerRejectSubscriber) error {
	return a.conditionalOrderTriggerRejectSubscribers.unsubscribe(subscriber)
}

func (a *accountsWebsocket) Close() error {
	return a.ws.Close()
}

func (a *accountsWebsocket) Run(ctx context.Context) {
	a.ws.Run(ctx, a.runHandler)
}

func (a *accountsWebsocket) RunNewThread(ctx context.Context) {
	a.ws.RunNewThread(ctx, a.runHandler)
}

func (a *accountsWebsocket) runHandler(msgType int, msg []byte, msgError error) {
	var (
		msgMap map[string]interface{}
		err    error
	)

	if msgType == -1 {
		a.disconnectHandler(msgType, msgError)
		return
	}

	if msgError != nil {
		a.errorSubscribers.update(msgError)
		return
	}

	if err = json.Unmarshal(msg, &msgMap); err != nil {
		a.errorSubscribers.update(err)
		return
	}

	if event, eventExist := msgMap["e"]; eventExist {
		a.eventNameHandler(fmt.Sprint(event), msg, msgMap)
		return
	}

	if requestID, idExist := msgMap["id"]; idExist {
		a.requestHandler(requestID, msg)
		return
	}
}

func (a *accountsWebsocket) disconnectHandler(msgType int, err error) {

	var disconnectError error = fmt.Errorf("websocket disconnect, msg type: %d, msg: %w", msgType, err)

	a.errorSubscribers.update(disconnectError)

	a.listenKeyExpiredSubscribers.update(nil, disconnectError)

	a.accountUpdateSubscribers.update(nil, disconnectError)

	a.marginCallSubscribers.update(nil, disconnectError)

	a.orderTradeUpdateSubscribers.update(nil, disconnectError)

	a.accountConfigUpdateSubscribers.update(nil, disconnectError)

	a.accountInfoUpdateSubscribers.update(nil, disconnectError)

	a.strategyUpdateSubscribers.update(nil, disconnectError)

	a.gridUpdateSubscribers.update(nil, disconnectError)

	a.conditionalOrderTriggerRejectSubscribers.update(nil, disconnectError)

	a.rateLimitsSubscribers.update(nil, disconnectError)

}

func (a *accountsWebsocket) eventNameHandler(eventName string, msg []byte, msgMap map[string]interface{}) {

	switch eventName {
	case EVENT_TYPE_LISTEN_KEY_EXPIRED:
		a.listenKeyExpiredhandler(msg)
		return
	case EVENT_TYPE_ACCOUNT_UPDATE:
		a.accountUpdateHandler(msg)
		return
	case EVENT_TYPE_MARGIN_CALL:
		a.marginCallHandler(msg)
		return
	case EVENT_TYPE_ORDER_TRADE_UPDATE:
		a.orderTradeUpdateHandler(msg)
		return
	case EVENT_TYPE_ACCOUNT_CONFIG_UPDATE:

		if _, exist := msgMap["ac"]; exist {
			a.accountConfigUpdateHandler(msg)
			return
		}

		if _, exist := msgMap["ai"]; exist {
			a.accountInfoUpdateHandler(msg)
			return
		}

		var accountConfigUpdateError error = fmt.Errorf("unknown account config update message : %s", string(msg))
		a.errorSubscribers.update(accountConfigUpdateError)
		a.accountConfigUpdateSubscribers.update(nil, accountConfigUpdateError)
		a.accountInfoUpdateSubscribers.update(nil, accountConfigUpdateError)
		return

	case EVENT_TYPE_STRATEGY_UPDATE:
		a.strategyUpdateHandler(msg)
		return
	case EVENT_TYPE_GRID_UPDATE:
		a.gridUpdateHandler(msg)
		return
	case EVENT_TYPE_CONDITIONAL_ORDER_TRIGGER_REJECT:
		a.conditionalOrderTriggerRejectHandler(msg)
		return
	default:
		a.errorSubscribers.update(fmt.Errorf("unknown event name : %s", eventName))
	}

}

func (a *accountsWebsocket) requestHandler(requestID interface{}, msg []byte) {

	var err error = a.bytesChannel.WriteChannel(requestID, msg, nil)

	if err != nil {
		a.errorSubscribers.update(err)
	}
}

func (a *accountsWebsocket) listenKeyExpiredhandler(msg []byte) {

	var (
		resp listenKeyExpired
		err  error
	)

	if err = json.Unmarshal(msg, &resp); err != nil {
		a.errorSubscribers.update(err)
		a.listenKeyExpiredSubscribers.update(nil, err)
		return
	}

	a.listenKeyExpiredSubscribers.update(resp, nil)

}

func (a *accountsWebsocket) accountUpdateHandler(msg []byte) {

	var (
		resp accountUpdate
		err  error
	)

	if err = json.Unmarshal(msg, &resp); err != nil {
		a.errorSubscribers.update(err)
		a.accountUpdateSubscribers.update(nil, err)
		return
	}

	a.accountUpdateSubscribers.update(resp, nil)

}

func (a *accountsWebsocket) marginCallHandler(msg []byte) {

	var (
		resp marginCall
		err  error
	)

	if err = json.Unmarshal(msg, &resp); err != nil {
		a.errorSubscribers.update(err)
		a.marginCallSubscribers.update(nil, err)
		return
	}

	a.marginCallSubscribers.update(resp, nil)

}

func (a *accountsWebsocket) orderTradeUpdateHandler(msg []byte) {

	var (
		resp orderTradeUpdate
		err  error
	)

	if err = json.Unmarshal(msg, &resp); err != nil {
		a.orderTradeUpdateSubscribers.update(nil, err)
		return
	}

	a.orderTradeUpdateSubscribers.update(resp, nil)

}

func (a *accountsWebsocket) accountConfigUpdateHandler(msg []byte) {

	var (
		resp accountConfigUpdate
		err  error
	)

	if err = json.Unmarshal(msg, &resp); err != nil {
		a.accountConfigUpdateSubscribers.update(nil, err)
		return
	}

	a.accountConfigUpdateSubscribers.update(resp, nil)

}

func (a *accountsWebsocket) accountInfoUpdateHandler(msg []byte) {

	var (
		resp accountInfoUpdate
		err  error
	)

	if err = json.Unmarshal(msg, &resp); err != nil {
		a.accountInfoUpdateSubscribers.update(nil, err)
		return
	}

	a.accountInfoUpdateSubscribers.update(resp, nil)
}

func (a *accountsWebsocket) strategyUpdateHandler(msg []byte) {
	var (
		resp strategyUpdate
		err  error
	)

	if err = json.Unmarshal(msg, &resp); err != nil {
		a.strategyUpdateSubscribers.update(nil, err)
		return
	}

	a.strategyUpdateSubscribers.update(resp, nil)
}

func (a *accountsWebsocket) gridUpdateHandler(msg []byte) {
	var (
		resp gridUpdate
		err  error
	)

	if err = json.Unmarshal(msg, &resp); err != nil {
		a.gridUpdateSubscribers.update(nil, err)
		return
	}

	a.gridUpdateSubscribers.update(resp, nil)
}

func (a *accountsWebsocket) conditionalOrderTriggerRejectHandler(msg []byte) {
	var (
		resp conditionalOrderTriggerReject
		err  error
	)

	if err = json.Unmarshal(msg, &resp); err != nil {
		a.conditionalOrderTriggerRejectSubscribers.update(nil, err)
		return
	}

	a.conditionalOrderTriggerRejectSubscribers.update(resp, nil)
}

// func (a *accountsWebsocket) rateLimithandler(rateLimits []rateLimit, err error) {

// 	var limits RateLimits = make(RateLimits, len(rateLimits))

// 	for _, rateLimit := range rateLimits {
// 		limits = append(limits, rateLimit)
// 	}

// 	a.rateLimitsSubscribers.update(limits, err)
// }
