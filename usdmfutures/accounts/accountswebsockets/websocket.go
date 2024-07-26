package accountswebsockets

import (
	"context"
	"fmt"

	"github.com/Lornzo/gobinance/binancewebsockets"
	"github.com/Lornzo/gobinance/usdmfutures/accounts/accountsrests"
	"github.com/Lornzo/gobinance/usdmfutures/usdmfutureswebsockettypes"
)

type Websocket interface {
	ExtendListenKey(ctx context.Context) (string, error)
	CloseListenKey(ctx context.Context) (string, error)
	RequestUserPosition(ctx context.Context) (AccountPositions, error)
	RequestBalance(ctx context.Context) (AccountBalances, error)
	RequestAccountInformation(ctx context.Context) (AccountInformation, error)
	SubscribeRateLimits(subscriber RateLimitsSubscriber) error
	UnSubscribeRateLimits(subscriber RateLimitsSubscriber) error
	SubscribeError(subscriber ErrorSubscriber) error
	UnSubscribeError(subscriber ErrorSubscriber) error
	SubscribeListenKeyExpired(subscriber ListenKeyExpiredSubscriber) error
	UnSubscribeListenKeyExpired(subscriber ListenKeyExpiredSubscriber) error
	SubscribeAccountUpdate(subscriber AccountUpdateSubscriber) error
	UnSubscribeAccountUpdate(subscriber AccountUpdateSubscriber) error
	SubscribeMarginCall(subscriber MarginCallSubscriber) error
	UnSubscribeMarginCall(subscriber MarginCallSubscriber) error
	SubscribeOrderUpdate(subscriber OrderTradeUpdateSubscriber) error
	UnSubscribeOrderUpdate(subscriber OrderTradeUpdateSubscriber) error
	SubscribeAccountConfigUpdate(subscriber AccountConfigUpdateSubscriber) error
	UnSubscribeAccountConfigUpdate(subscriber AccountConfigUpdateSubscriber) error
	SubscribeAccountInfoUpdate(subscriber AccountInfoUpdateSubscriber) error
	UnSubscribeAccountInfoUpdate(subscriber AccountInfoUpdateSubscriber) error
	SubscribeStrategyUpdate(subscriber StrategyUpdateSubscriber) error
	UnSubscribeStrategyUpdate(subscriber StrategyUpdateSubscriber) error
	SubscribeGridUpdate(subscriber GridUpdateSubscriber) error
	UnSubscribeGridUpdate(subscriber GridUpdateSubscriber) error
	SubscribeConditionalOrderTriggerReject(subscriber ConditionalOrderTriggerRejectSubscriber) error
	UnSubscribeConditionalOrderTriggerReject(subscriber ConditionalOrderTriggerRejectSubscriber) error
	Close() error
	Run(ctx context.Context)
	RunNewThread(ctx context.Context)
}

func NewWebsocket(ctx context.Context, binanceWebsocketURL string, binanceRestURL string, account Account) (Websocket, error) {

	var (
		rest restful = accountsrests.Restful{
			Account: account,
			BaseURL: binanceRestURL,
		}
		err           error
		ws            *accountsWebsocket = &accountsWebsocket{}
		listenKeyResp accountsrests.ListenKeyResponse
	)

	if listenKeyResp, err = rest.ListenKeyCreate().DoRequest(ctx); err != nil {
		return nil, err
	}

	ws.concreateRest = rest
	ws.listenKey.Set(listenKeyResp.ListenKey)
	ws.account.Set(account.GetAPIKey(), account.GetAPISecret())
	ws.errorSubscribers = &usdmfutureswebsockettypes.ErrorSubscribers{}

	if ws.ws, err = binancewebsockets.NewWebsocket(fmt.Sprint(binanceWebsocketURL, "/ws/", ws.listenKey.Get())); err != nil {
		return nil, err
	}

	return ws, nil

}
