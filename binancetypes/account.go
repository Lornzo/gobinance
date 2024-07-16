package binancetypes

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"sort"
	"strings"
)

type Account struct {
	APIKey    string
	APISecret string
}

func (a Account) GetAPIKey() string {
	return a.APIKey
}

func (a Account) GetAPISecret() string {
	return a.APISecret
}

func (a Account) GetSignatureWithQueries(queries ...string) (string, error) {

	var dstQueries []string = queries

	sort.Strings(dstQueries)

	return a.GetSignatureWithQueriesString(strings.Join(dstQueries, "&"))
}

func (a Account) GetSignatureWithQueriesString(queriesString string) (string, error) {

	if a.APISecret == "" {
		return "", errors.New("APIKey is empty")
	}

	var (
		err          error
		dstSignature string
		h            hash.Hash = hmac.New(sha256.New, []byte(a.APISecret))
	)

	if _, err = h.Write([]byte(queriesString)); err != nil {
		return "", err
	}

	dstSignature = fmt.Sprintf("%x", h.Sum(nil))
	return dstSignature, nil
}
