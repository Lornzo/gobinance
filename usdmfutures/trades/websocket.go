package trades

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Lornzo/gobinance/channels"
	"github.com/Lornzo/gobinance/commons"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func NewWebsocket(wsURL string) *Websocket {
	var dstWebsocket *Websocket = &Websocket{
		Base: commons.NewBinanceWebsocketBase(),
	}
	dstWebsocket.Base.SetBaseURL(wsURL)
	return dstWebsocket
}

type Websocket struct {
	Base         commons.WebsocketBase
	wsConn       *websocket.Conn
	bytesChannel channels.BytesChannel
}

func (w *Websocket) Start(ctx context.Context) error {

	if w.Base.IsRunning() {
		return errors.New("websocket is running")
	}

	w.Base.SetRunning(true)

	var (
		err   error
		resp  *http.Response
		wsURL string = w.Base.GetWebsocketURL()
	)

	defer func() {
		if err != nil {
			w.Base.SetRunning(false)
		}
	}()

	if w.wsConn, resp, err = websocket.DefaultDialer.Dial(wsURL, nil); err != nil {
		return err
	}

	if resp.StatusCode != http.StatusSwitchingProtocols {
		err = fmt.Errorf("websocket connection faild, code: %d", resp.StatusCode)
		return err
	}

	go w.listenToSource()

	return nil

}

func (w *Websocket) listenToSource() {

	for {

		var (
			msgType  int
			msgBytes []byte
			err      error
			resp     struct {
				ID     string                 `json:"id"`
				Status int                    `json:"status"`
				Result map[string]interface{} `json:"result,omitempty"`
				Error  struct {
					Code int    `json:"code"`
					Msg  string `json:"msg"`
				} `json:"error,omitempty"`
			}
		)

		msgType, msgBytes, err = w.wsConn.ReadMessage()

		if msgType == -1 {
			break
		}

		if err != nil {
			w.bytesChannel.WriteChannel(resp.ID, nil, err)
			continue
		}

		if err = json.Unmarshal(msgBytes, &resp); err != nil {
			log.Println("error unmarshal => ", err)
			continue
		}

		if resp.Status != http.StatusOK {
			w.bytesChannel.WriteChannel(resp.ID, nil, fmt.Errorf("status %d => code %d : %s ", resp.Status, resp.Error.Code, resp.Error.Msg))
			continue
		}

		if resp.Result == nil {
			w.bytesChannel.WriteChannel(resp.ID, nil, fmt.Errorf("result is nil"))
			continue
		}

		if msgBytes, err = json.Marshal(resp.Result); err != nil {
			w.bytesChannel.WriteChannel(resp.ID, nil, err)
			continue
		}

		if err = w.bytesChannel.WriteChannel(resp.ID, msgBytes, nil); err != nil {
			log.Println("error write channel => ", err)
		}

	}
}

func (w *Websocket) makeRequest(request interface{}) error {

	if !w.Base.IsRunning() {
		return errors.New("websocket is not running")
	}

	if err := w.wsConn.WriteJSON(request); err != nil {
		return err
	}

	return nil

}

func (w *Websocket) Close() error {

	if !w.Base.IsRunning() {
		return nil
	}

	var err = w.wsConn.Close()

	if err != nil {
		return err
	}

	w.Base.SetRunning(false)

	return nil

}

func (w *Websocket) CancelOrder(ctx context.Context, account Account, order CancelOrderForm) (CanceledOrder, error) {

	var (
		canceledOrder CanceledOrder
		response      []byte
		err           error
		channelID     string = uuid.NewString()
		request       struct {
			ID     string                 `json:"id"`
			Method string                 `json:"method"`
			Params map[string]interface{} `json:"params"`
		}
		builder cancelOrderFormBuilder = cancelOrderFormBuilder{
			account: account,
			order:   order,
		}
	)

	request.ID = channelID
	request.Method = "order.cancel"

	if err = w.bytesChannel.CreateChannel(channelID); err != nil {
		return CanceledOrder{}, err
	}

	defer w.bytesChannel.CloseChannel(channelID)

	if request.Params, err = builder.buildRequestBody(); err != nil {
		return CanceledOrder{}, err
	}

	if err = w.makeRequest(request); err != nil {
		return CanceledOrder{}, err
	}

	if response, err = w.bytesChannel.ReadChannel(channelID); err != nil {
		return CanceledOrder{}, err
	}

	if err = json.Unmarshal(response, &canceledOrder); err != nil {
		return CanceledOrder{}, err
	}

	return canceledOrder, nil

}

func (w *Websocket) CreateOrder(ctx context.Context, account Account, order PlaceOrderForm) (PlacedOrder, error) {

	type Request struct {
		ID     string                 `json:"id"`
		Method string                 `json:"method"`
		Params map[string]interface{} `json:"params"`
	}

	var (
		placedOrder PlacedOrder
		resp        []byte
		req         Request = Request{
			ID:     order.GetNewClientOrderID(),
			Method: "order.place",
		}
		builder placeOrderFormBuilder = placeOrderFormBuilder{
			account: account,
			order:   order,
		}
		err error
	)

	if req.ID == "" {
		req.ID = uuid.NewString()
	}

	if err = w.bytesChannel.CreateChannel(req.ID); err != nil {
		return PlacedOrder{}, err
	}

	defer w.bytesChannel.CloseChannel(req.ID)

	if req.Params, err = builder.buildRequestBody(); err != nil {
		return PlacedOrder{}, err
	}

	if err = w.makeRequest(req); err != nil {
		return PlacedOrder{}, err
	}

	if resp, err = w.bytesChannel.ReadChannel(req.ID); err != nil {
		return PlacedOrder{}, err
	}

	if err = json.Unmarshal(resp, &placedOrder); err != nil {
		return PlacedOrder{}, err
	}

	return placedOrder, nil

}
