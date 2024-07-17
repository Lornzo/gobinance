package trades

import (
	"errors"
	"fmt"
)

type cancelOrderFormBuilder struct {
	account Account
	order   CancelOrderForm
}

func (c cancelOrderFormBuilder) check() error {

	if c.account == nil {
		return errors.New("need object account in builder")
	}

	if c.order == nil {
		return errors.New("need object order in builder")
	}

	return nil

}

func (c cancelOrderFormBuilder) buildRequestBody() (map[string]interface{}, error) {

	var (
		bodyMap   map[string]interface{}
		err       error
		queries   []string
		signature string
	)

	if queries, err = c.buildQueriesArray(); err != nil {
		return nil, err
	}

	if signature, err = c.account.GetSignatureWithQueries(queries...); err != nil {
		return nil, err
	}

	if bodyMap, err = c.buildMap(); err != nil {
		return nil, err
	}

	bodyMap["signature"] = signature

	return bodyMap, nil

}

func (c cancelOrderFormBuilder) buildQueriesArray() ([]string, error) {

	var (
		err     error
		queries []string
		dstMap  map[string]interface{}
	)

	if dstMap, err = c.buildMap(); err != nil {
		return nil, err
	}

	for key, value := range dstMap {
		queries = append(queries, fmt.Sprint(key, "=", value))
	}

	return queries, nil
}

func (c cancelOrderFormBuilder) buildMap() (map[string]interface{}, error) {

	var (
		err    error
		dstMap map[string]interface{} = make(map[string]interface{})
	)

	if err = c.check(); err != nil {
		return nil, err
	}

	dstMap["apiKey"] = c.account.GetAPIKey()

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
