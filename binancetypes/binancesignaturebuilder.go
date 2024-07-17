package binancetypes

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"sort"
	"strings"
	"time"
)

type BinanceSignatureBuilderAccount interface {
	GetAPIKey() string
	GetAPISecret() string
}

func NewBinanceSignatureBuilder() *BinanceSignatureBuilder {
	return &BinanceSignatureBuilder{}
}

type BinanceSignatureBuilder struct {
	APIKey    string
	APISecret string
}

func (b *BinanceSignatureBuilder) SetAPIKeyAndSecret(key string, secret string) {
	b.APIKey = key
	b.APISecret = secret
}

func (b *BinanceSignatureBuilder) SetAPIKeyAndSecretWithAccount(account BinanceSignatureBuilderAccount) {
	b.APIKey = account.GetAPIKey()
	b.APISecret = account.GetAPISecret()
}

func (b *BinanceSignatureBuilder) BuildRequestBodyWithSignature(body map[string]interface{}) (map[string]interface{}, error) {

	var (
		dstMap    map[string]interface{} = body
		signature string
		err       error
	)

	if b.APIKey == "" {
		return nil, errors.New("api key is empty")
	}

	dstMap["apiKey"] = b.APIKey

	if _, exist := dstMap["timestamp"]; !exist {
		dstMap["timestamp"] = time.Now().UnixMilli()
	}

	if signature, err = b.getSignatureWithQueries(b.getQueriesArrayWithMapForSignature(dstMap)...); err != nil {
		return nil, err
	}

	dstMap["signature"] = signature

	return dstMap, nil

}

func (b *BinanceSignatureBuilder) getQueriesArrayWithMapForSignature(datas map[string]interface{}) []string {

	var queries []string

	for key, value := range datas {
		queries = append(queries, fmt.Sprint(key, "=", value))
	}

	return queries

}

func (b *BinanceSignatureBuilder) getSignatureWithQueries(queries ...string) (string, error) {

	var dstQueries []string = queries

	sort.Strings(dstQueries)

	return b.getSignatureWithQueriesString(strings.Join(dstQueries, "&"))
}

func (b *BinanceSignatureBuilder) getSignatureWithQueriesString(queriesString string) (string, error) {

	if b.APISecret == "" {
		return "", errors.New("api secret is empty")
	}

	var (
		err          error
		dstSignature string
		h            hash.Hash = hmac.New(sha256.New, []byte(b.APISecret))
	)

	if _, err = h.Write([]byte(queriesString)); err != nil {
		return "", err
	}

	dstSignature = fmt.Sprintf("%x", h.Sum(nil))
	return dstSignature, nil
}
