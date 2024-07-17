package trades

import (
	"errors"
)

type cancelOrderFormBuilder struct {
	account         Account
	order           CancelOrderForm
	concreteBuilder signatureBuilder
}

func (c cancelOrderFormBuilder) check() error {

	if c.account == nil {
		return errors.New("need object account in builder")
	}

	if c.order == nil {
		return errors.New("need object order in builder")
	}

	if c.concreteBuilder == nil {
		return errors.New("need signatureBuilder in builder")
	}

	return nil

}

func (c cancelOrderFormBuilder) buildRequestBody() (map[string]interface{}, error) {

	var (
		bodyMap map[string]interface{}
		err     error
	)

	if bodyMap, err = c.buildMap(); err != nil {
		return nil, err
	}

	c.concreteBuilder.SetAPIKeyAndSecret(c.account.GetAPIKey(), c.account.GetAPISecret())

	return c.concreteBuilder.BuildRequestBodyWithSignature(bodyMap)

}

func (c cancelOrderFormBuilder) buildMap() (map[string]interface{}, error) {

	var (
		err    error
		dstMap map[string]interface{} = make(map[string]interface{})
	)

	if err = c.check(); err != nil {
		return nil, err
	}

	if symbol := c.order.GetSymbol(); symbol != "" {
		dstMap["symbol"] = symbol
	}

	if orderID := c.order.GetOrderID(); orderID != 0 {
		dstMap["orderId"] = orderID
	}

	if originClientOrderID := c.order.GetOriginClientOrderID(); originClientOrderID != "" {
		dstMap["origClientOrderId"] = originClientOrderID
	}

	if recvWindow := c.order.GetRecvWindow(); recvWindow != 0 {
		dstMap["recvWindow"] = recvWindow
	}

	if timestamp := c.order.GetTimestamp(); timestamp != 0 {
		dstMap["timestamp"] = timestamp
	}

	return dstMap, nil
}
