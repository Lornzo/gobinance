package commons

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Lornzo/gobinance/threadsafetypes"
)

func NewBinanceWebsocketBase() *BinanceWebsocketBase {
	return &BinanceWebsocketBase{}
}

type BinanceWebsocketBase struct {
	baseURL threadsafetypes.String
	pathes  threadsafetypes.StringSlice
	running threadsafetypes.Bool
	debug   threadsafetypes.Bool
}

func (b *BinanceWebsocketBase) SetBaseURL(url string) error {

	if b.IsRunning() {
		return fmt.Errorf("cannot set base url while websocket is running")
	}

	b.baseURL.Set(url)

	return nil
}

func (b *BinanceWebsocketBase) SetPathes(pathes ...string) error {

	if b.IsRunning() {
		return fmt.Errorf("cannot set pathes while websocket is running")
	}

	b.pathes.Set(pathes...)

	return nil
}

func (b *BinanceWebsocketBase) GetWebsocketURL() string {

	var (
		wg     sync.WaitGroup
		wsURL  string
		pathes []string
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		wsURL = b.baseURL.Get()
	}()

	go func() {
		defer wg.Done()
		pathes = b.pathes.Get()
	}()

	wg.Wait()

	if len(pathes) > 0 {
		wsURL += fmt.Sprint("/", strings.Join(pathes, "/"))
	}
	return wsURL
}

func (b *BinanceWebsocketBase) SetRunning(isRunning bool) {
	b.running.Set(isRunning)
}

func (b *BinanceWebsocketBase) IsRunning() bool {
	return b.running.Get()
}

func (b *BinanceWebsocketBase) SetDebug(isDebug bool) {
	b.debug.Set(isDebug)
}

func (b *BinanceWebsocketBase) IsDebug() bool {
	return b.debug.Get()
}
