package accountswebsockets

import (
	"context"
	"fmt"

	"github.com/Lornzo/gobinance/binancewebsockets"
	"github.com/Lornzo/gobinance/usdmfutures/accounts/accountsrests"
	"github.com/Lornzo/gobinance/usdmfutures/usdmfutureswebsockettypes"
)

type Websocket interface {
	RequestUserPosition(ctx context.Context) (AccountPositions, error)
	RequestBalance(ctx context.Context) (AccountBalances, error)
	RequestAccountInformation(ctx context.Context) (AccountInformation, error)
	CloseListenKey(ctx context.Context) (string, error)
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
