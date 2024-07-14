package commons

import (
	"fmt"
	"strings"
	"sync"
)

func NewBinanceWebsocketBase() *BinanceWebsocketBase {
	return &BinanceWebsocketBase{}
}

type BinanceWebsocketBase struct {
	APIKey      string
	APISecret   string
	BaseURL     string
	Pathes      []string
	running     bool
	runningLock sync.RWMutex
}

func (b *BinanceWebsocketBase) SetAPIKeyAndSecret(key string, secret string) {
	b.APIKey = key
	b.APISecret = secret
}

func (b *BinanceWebsocketBase) SetBaseURL(url string) {
	b.BaseURL = url
}

func (b *BinanceWebsocketBase) SetPathes(pathes ...string) {
	b.Pathes = pathes
}

func (b *BinanceWebsocketBase) GetWebsocketURL() string {
	var url string = b.BaseURL
	if len(b.Pathes) > 0 {
		url += fmt.Sprint("/", strings.Join(b.Pathes, "/"))
	}
	return url
}

func (b *BinanceWebsocketBase) SetRunning(isRunning bool) {
	b.runningLock.Lock()
	defer b.runningLock.Unlock()
	b.running = isRunning
}

func (b *BinanceWebsocketBase) IsRunning() bool {
	b.runningLock.RLock()
	defer b.runningLock.RUnlock()
	return b.running
}
