package marketdataswebsockets

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/Lornzo/gobinance/commons"
	"github.com/gorilla/websocket"
)

func newBinanceDataSource(baseURL string) *binanceDataSource {

	var dataSource *binanceDataSource = &binanceDataSource{
		listenerList: newBinanceDataSourceListenerList(),
		base:         commons.NewBinanceWebsocketBase(),
	}

	dataSource.base.SetBaseURL(baseURL)

	return dataSource
}

type binanceDataSource struct {
	wsConn       *websocket.Conn
	listenerList dataSourceListenerList
	base         commons.WebsocketBase
	running      bool
	runningLock  sync.RWMutex
}

func (b *binanceDataSource) setRunning(isRunning bool) {
	b.runningLock.Lock()
	defer b.runningLock.Unlock()
	b.running = isRunning
}

func (b *binanceDataSource) isRunning() bool {
	b.runningLock.RLock()
	defer b.runningLock.RUnlock()
	return b.running
}

func (b *binanceDataSource) start(ctx context.Context) error {

	if b.isRunning() {
		return errors.New("binance data source is running")
	}

	b.setRunning(true)

	b.initWebsocketURL()

	var (
		resp  *http.Response
		err   error
		wsURL string = b.base.GetWebsocketURL()
	)

	defer func() {
		if err != nil {
			b.setRunning(false)
		}
	}()

	if b.wsConn, resp, err = websocket.DefaultDialer.Dial(wsURL, nil); err != nil {
		return err
	}

	if resp.StatusCode != http.StatusSwitchingProtocols {
		err = fmt.Errorf("websocket connection faild, code: %d", resp.StatusCode)
		return err
	}

	go func() {
		b.listenToSource()
	}()

	return nil
}

func (b *binanceDataSource) listenToSource() {

	if b.wsConn == nil {
		return
	}

	defer b.setRunning(false)

	for {

		var (
			msgType int
			msg     []byte
			err     error
		)

		if msgType, msg, err = b.wsConn.ReadMessage(); err != nil {
			b.errorHandler(err)
		}

		if msgType == -1 {
			break
		}

		if err = b.msgHandler(msg); err != nil {
			b.errorHandler(err)
		}

	}
}

func (b *binanceDataSource) errorHandler(err error) {

}

func (b *binanceDataSource) close() error {

	if !b.isRunning() {
		return nil
	}

	return b.wsConn.Close()

}

func (b *binanceDataSource) msgHandler(msg []byte) error {

	var (
		msgMap map[string]interface{}
		err    error
	)

	if err = json.Unmarshal(msg, &msgMap); err != nil {
		return err
	}

	if stream, exist := msgMap["stream"]; exist {
		go func() {
			var runtineErr error = b.listenerList.update(fmt.Sprint(stream), msg)
			if runtineErr != nil {
				b.errorHandler(runtineErr)
			}
		}()
	}

	return nil
}

func (b *binanceDataSource) addListener(listener dataSourceListener) error {

	type SubscribeRequest struct {
		Method string   `json:"method"`
		Params []string `json:"params"`
		ID     int      `json:"id"`
	}

	var (
		err     error
		request SubscribeRequest = SubscribeRequest{
			Method: "SUBSCRIBE",
			Params: []string{listener.getStreamName()},
			ID:     b.listenerList.addListener(listener),
		}
	)

	if request.ID == 0 {
		return fmt.Errorf("stream name %s has already exist", listener.getStreamName())
	}

	if !b.isRunning() {
		return nil
	}

	defer func() {
		if err != nil {
			b.listenerList.rmListener(listener)
		}
	}()

	if err = b.wsConn.WriteJSON(request); err != nil {
		return err
	}

	return nil

}

func (b *binanceDataSource) rmListener(listener dataSourceListener) error {

	type UnSubscribeRequest struct {
		Method string   `json:"method"`
		Params []string `json:"params"`
		ID     int      `json:"id"`
	}

	var request UnSubscribeRequest = UnSubscribeRequest{
		Method: "UNSUBSCRIBE",
		Params: []string{listener.getStreamName()},
		ID:     b.listenerList.rmListener(listener),
	}

	if !b.isRunning() || request.ID == 0 {
		return nil
	}

	return b.wsConn.WriteJSON(request)

}

func (b *binanceDataSource) initWebsocketURL() {
	b.base.SetPathes("stream")
}
