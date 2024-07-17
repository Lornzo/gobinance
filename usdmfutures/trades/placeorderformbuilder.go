package trades

import (
	"errors"
)

type placeOrderFormBuilder struct {
	account Account
	order   PlaceOrderForm
	builder signatureBuilder
}

func (p placeOrderFormBuilder) check() error {

	if p.account == nil {
		return errors.New("need object account in builder")
	}

	if p.order == nil {
		return errors.New("need object order in builder")
	}

	if p.builder == nil {
		return errors.New("need signatureBuilder in builder")
	}

	return nil

}

func (p placeOrderFormBuilder) buildRequestBody() (map[string]interface{}, error) {

	var (
		bodyMap map[string]interface{}
		err     error
	)

	if bodyMap, err = p.buildMap(); err != nil {
		return nil, err
	}

	p.builder.SetAPIKeyAndSecret(p.account.GetAPIKey(), p.account.GetAPISecret())

	return p.builder.BuildRequestBodyWithSignature(bodyMap)

}

func (p placeOrderFormBuilder) buildMap() (map[string]interface{}, error) {

	var (
		err    error
		dstMap map[string]interface{} = make(map[string]interface{})
	)

	if err = p.check(); err != nil {
		return nil, err
	}

	if symbol := p.order.GetSymbol(); symbol != "" {
		dstMap["symbol"] = symbol
	}

	if side := p.order.GetSide(); side != "" {
		dstMap["side"] = side
	}

	if positionSide := p.order.GetPositionSide(); positionSide != "" {
		dstMap["positionSide"] = positionSide
	}

	if orderType := p.order.GetType(); orderType != "" {
		dstMap["type"] = orderType
	}

	if reduceOnly := p.order.GetReduceOnly(); reduceOnly != "" {
		dstMap["reduceOnly"] = reduceOnly
	}

	if quantity := p.order.GetQuantity(); !quantity.IsZero() {
		dstMap["quantity"] = quantity
	}

	if price := p.order.GetPrice(); !price.IsZero() {
		dstMap["price"] = price
	}

	if newClientOrderID := p.order.GetNewClientOrderID(); newClientOrderID != "" {
		dstMap["newClientOrderID"] = newClientOrderID
	}

	if stopPrice := p.order.GetStopPrice(); !stopPrice.IsZero() {
		dstMap["stopPrice"] = stopPrice
	}

	if closePosition := p.order.GetClosePosition(); closePosition != "" {
		dstMap["closePosition"] = closePosition
	}

	if activationPrice := p.order.GetActivationPrice(); !activationPrice.IsZero() {
		dstMap["activationPrice"] = activationPrice
	}

	if callbackRate := p.order.GetCallbackRate(); !callbackRate.IsZero() {
		dstMap["callbackRate"] = callbackRate
	}

	if timeInForce := p.order.GetTimeInForce(); timeInForce != "" {
		dstMap["timeInForce"] = timeInForce
	}

	if workingType := p.order.GetWorkingType(); workingType != "" {
		dstMap["workingType"] = workingType
	}

	if priceProtect := p.order.GetPriceProtect(); priceProtect != "" {
		dstMap["priceProtect"] = priceProtect
	}

	if newOrderRespType := p.order.GetNewOrderRespType(); newOrderRespType != "" {
		dstMap["newOrderRespType"] = newOrderRespType
	}

	if recvWindow := p.order.GetRecvWindow(); recvWindow != 0 {
		dstMap["recvWindow"] = recvWindow
	}

	if timestamp := p.order.GetTimestamp(); timestamp != 0 {
		dstMap["timestamp"] = timestamp
	}

	return dstMap, nil
}
