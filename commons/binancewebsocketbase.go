package commons

import (
	"fmt"
	"strings"
)

func NewBinanceWebsocketBase() *BinanceWebsocketBase {
	return &BinanceWebsocketBase{}
}

type BinanceWebsocketBase struct {
	APIKey    string
	APISecret string
	BaseURL   string
	Pathes    []string
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
