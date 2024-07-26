package tradeswebsockets

import (
	"errors"

	"github.com/Lornzo/gobinance/binancetypes"
)

type mapBuilder interface {
	buildMap() (map[string]interface{}, error)
}

func newFormBuilder(account Account, builder mapBuilder) formBuilder {
	return formBuilder{
		account:          account,
		mapBuilder:       builder,
		signatureBuilder: binancetypes.NewBinanceSignatureBuilder(),
	}
}

type formBuilder struct {
	account          Account
	signatureBuilder signatureBuilder
	mapBuilder       interface {
		buildMap() (map[string]interface{}, error)
	}
}

func (p formBuilder) check() error {

	if p.account == nil {
		return errors.New("need object account in builder")
	}

	if p.signatureBuilder == nil {
		return errors.New("need signatureBuilder in builder")
	}

	return nil

}

func (p formBuilder) buildRequestBody() (map[string]interface{}, error) {

	var (
		bodyMap map[string]interface{}
		err     error
	)

	if err = p.check(); err != nil {
		return nil, err
	}

	if bodyMap, err = p.mapBuilder.buildMap(); err != nil {
		return nil, err
	}

	p.signatureBuilder.SetAPIKeyAndSecret(p.account.GetAPIKey(), p.account.GetAPISecret())

	return p.signatureBuilder.BuildRequestBodyWithSignature(bodyMap)

}
