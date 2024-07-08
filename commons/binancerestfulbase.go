package commons

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"hash"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

func NewBinanceRestfulBase() *BinanceRestfulBase {
	return &BinanceRestfulBase{
		Queries: NewMapQueries(),
	}
}

type BinanceRestfulBase struct {
	APIKey      string
	APISecret   string
	BaseURL     string
	Pathes      []string
	Queries     Queries
	IsSignature bool
	IsAPIKey    bool
}

func (b BinanceRestfulBase) Copy() RestfulBase {
	return &b
}

func (b *BinanceRestfulBase) SetAPIKeyAndSecret(key string, secret string) {
	b.APIKey = key
	b.APISecret = secret
}

func (b *BinanceRestfulBase) SetBaseURL(url string) {
	b.BaseURL = url
}

func (b *BinanceRestfulBase) SetPathes(pathes ...string) {
	b.Pathes = pathes
}

func (b *BinanceRestfulBase) AddPathes(pathes ...string) {
	b.Pathes = append(b.Pathes, pathes...)
}

func (b *BinanceRestfulBase) SetQuery(name string, value string) {
	b.Queries.SetQuery(name, value)
}

func (b *BinanceRestfulBase) DelQuery(name string) {
	b.Queries.DelQuery(name)
}

func (b *BinanceRestfulBase) GetQuery(name string) string {
	return b.Queries.GetQueryValue(name)
}

func (b *BinanceRestfulBase) UseApiKeyHeader(isUse bool) {
	b.IsAPIKey = isUse
}

func (b *BinanceRestfulBase) UseSignature(isUse bool) {
	b.IsSignature = isUse
	b.UseApiKeyHeader(isUse)
}

func (b *BinanceRestfulBase) GET(ctx context.Context, result interface{}) (int, error) {
	var request *resty.Request = b.getRestyRequest(ctx)
	request.Method = http.MethodGet
	if result != nil {
		request.SetResult(result)
	}
	return b.doRequest(request)
}

func (b *BinanceRestfulBase) POST(ctx context.Context, body interface{}, result interface{}) (int, error) {
	var request *resty.Request = b.getRestyRequest(ctx)
	request.Method = http.MethodPost
	if body != nil {
		request.SetBody(body)

	}
	if result != nil {
		request.SetResult(result)
	}
	return b.doRequest(request)
}

func (b *BinanceRestfulBase) DELETE(ctx context.Context, result interface{}) (int, error) {
	var request *resty.Request = b.getRestyRequest(ctx)
	request.Method = http.MethodDelete
	if result != nil {
		request.SetResult(result)
	}
	return b.doRequest(request)
}

func (b *BinanceRestfulBase) PUT(ctx context.Context, body interface{}, result interface{}) (int, error) {
	var request *resty.Request = b.getRestyRequest(ctx)
	request.Method = http.MethodPut
	if body != nil {
		request.SetBody(body)
	}
	if result != nil {
		request.SetResult(result)
	}
	return b.doRequest(request)
}

func (b *BinanceRestfulBase) doRequest(request *resty.Request) (int, error) {

	type ErrorResult struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	var (
		response    *resty.Response
		errorResult ErrorResult
		err         error
	)

	request.URL = b.getAPIURL()
	if response, err = request.SetError(&errorResult).Send(); err != nil {
		return http.StatusBadRequest, err
	}

	if otherErr := response.Error(); otherErr != nil {
		return http.StatusBadRequest, fmt.Errorf("%v", otherErr)
	}

	if response.IsError() {
		return response.StatusCode(), fmt.Errorf("code : %d | msg : %s", errorResult.Code, errorResult.Msg)
	}

	return response.StatusCode(), nil
}

func (b *BinanceRestfulBase) getRestyRequest(ctx context.Context) *resty.Request {

	var request *resty.Request = resty.New().R().SetContext(ctx)

	if b.IsAPIKey {
		request.SetHeaderVerbatim("X-MBX-APIKEY", b.APIKey)
	}

	return request
}

func (b *BinanceRestfulBase) getAPIURL() string {

	var (
		url     string   = b.BaseURL
		queries []string = b.Queries.ToArray()
	)

	if len(b.Pathes) > 0 {
		url = fmt.Sprint(url, "/", strings.Join(b.Pathes, "/"))
	}

	if b.IsSignature {
		queries = append(queries, "signature="+b.getSignature())
	}

	if len(queries) > 0 {
		url += "?" + strings.Join(queries, "&")
	}

	return url
}

func (b *BinanceRestfulBase) getSignature() string {
	var h hash.Hash = hmac.New(sha256.New, []byte(b.APISecret))
	h.Write([]byte(b.Queries.ToString()))
	var dstSignature string = fmt.Sprintf("%x", h.Sum(nil))
	return dstSignature

}
