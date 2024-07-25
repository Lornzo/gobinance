package binancewebsockets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func NewWebsocket(wsURL string) (Websocket, error) {

	var (
		err          error
		resp         *http.Response
		dstWebsocket *websocketWithSubscribers = &websocketWithSubscribers{
			subscribers: subscribers{
				list: make(map[string]Subscriber),
			},
		}
	)

	if dstWebsocket.Conn, resp, err = websocket.DefaultDialer.Dial(wsURL, nil); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusSwitchingProtocols {
		err = fmt.Errorf("websocket connection faild, code: %d", resp.StatusCode)
		return nil, err
	}

	return dstWebsocket, nil

}

type Websocket interface {
	Subscribe(subscriber Subscriber) error
	UnSubscribe(subscriber Subscriber) error
	MakeRequestByIntIndex(ctx context.Context, request Request) (Response, error)
	MakeRequestByUUIDIndex(ctx context.Context, request Request) (Response, error)
	MakeRequest(ctx context.Context, requestID interface{}, request Request) (Response, error)
	Run(ctx context.Context, handlers ...MessageHander)
	RunNewThread(ctx context.Context, handlers ...MessageHander)
	WriteJSON(interface{}) error
	Close() error
}
