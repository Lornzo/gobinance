package trades

import (
	"errors"
	"fmt"
)

type placeOrderFormBuilder struct {
	account Account
	order   PlaceOrderForm
}

func (p placeOrderFormBuilder) check() error {

	if p.account == nil {
		return errors.New("need object account in builder")
	}

	if p.order == nil {
		return errors.New("need object order in builder")
	}

	return nil

}

func (p placeOrderFormBuilder) buildRequestBody() (map[string]interface{}, error) {

	var (
		bodyMap   map[string]interface{}
		err       error
		queries   []string
		signature string
	)

	if queries, err = p.buildQueriesArray(); err != nil {
		return nil, err
	}

	if signature, err = p.account.GetSignatureWithQueries(queries...); err != nil {
		return nil, err
	}

	if bodyMap, err = p.buildMap(); err != nil {
		return nil, err
	}

	bodyMap["signature"] = signature

	return bodyMap, nil

}

func (p placeOrderFormBuilder) buildQueriesArray() ([]string, error) {
	var (
		err     error
		queries []string
		dstMap  map[string]interface{}
	)

	if dstMap, err = p.buildMap(); err != nil {
		return nil, err
	}

	for key, value := range dstMap {
		queries = append(queries, fmt.Sprint(key, "=", value))
	}

	return queries, nil

}

func (p placeOrderFormBuilder) buildMap() (map[string]interface{}, error) {

	var (
		err    error
		dstMap map[string]interface{} = make(map[string]interface{})
	)

	if err = p.check(); err != nil {
		return nil, err
	}

	dstMap["apiKey"] = p.account.GetAPIKey()

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

// type placeOrderFormBody struct {
// 	Symbol           string          `json:"symbol,omitempty"`
// 	Side             string          `json:"side,omitempty"`
// 	PositionSide     string          `json:"positionSide,omitempty"`
// 	Type             string          `json:"type,omitempty"`
// 	ReduceOnly       string          `json:"reduceOnly,omitempty"`
// 	Quantity         decimal.Decimal `json:"quantity,omitempty"`
// 	Price            decimal.Decimal `json:"price,omitempty"`
// 	NewClientOrderID string          `json:"newClientOrderId,omitempty"`
// 	StopPrice        decimal.Decimal `json:"stopPrice,omitempty"`
// 	ClosePosition    string          `json:"closePosition,omitempty"`
// 	ActivationPrice  decimal.Decimal `json:"activationPrice,omitempty"`
// 	CallbackRate     decimal.Decimal `json:"callbackRate,omitempty"`
// 	TimeInForce      string          `json:"timeInForce,omitempty"`
// 	WorkingType      string          `json:"workingType,omitempty"`
// 	PriceProtect     string          `json:"priceProtect,omitempty"`
// 	NewOrderRespType string          `json:"newOrderRespType,omitempty"`
// 	RecvWindow       int64           `json:"recvWindow,omitempty"`
// 	Timestamp        int64           `json:"timestamp,omitempty"`
// }

// func (p *placeOrderFormBody) AssignWithForm(form PlaceOrderForm) {
// 	p.Symbol = form.GetSymbol()
// 	p.Side = form.GetSide()
// 	p.PositionSide = form.GetPositionSide()
// 	p.Type = form.GetType()
// 	p.ReduceOnly = form.GetReduceOnly()
// 	p.Quantity = form.GetQuantity()
// 	p.Price = form.GetPrice()
// 	p.NewClientOrderID = form.GetNewClientOrderID()
// 	p.StopPrice = form.GetStopPrice()
// 	p.ClosePosition = form.GetClosePosition()
// 	p.ActivationPrice = form.GetActivationPrice()
// 	p.CallbackRate = form.GetCallbackRate()
// 	p.TimeInForce = form.GetTimeInForce()
// 	p.WorkingType = form.GetWorkingType()
// 	p.PriceProtect = form.GetPriceProtect()
// 	p.NewOrderRespType = form.GetNewOrderRespType()
// 	p.RecvWindow = form.GetRecvWindow()
// 	p.Timestamp = form.GetTimestamp()
// }

// func (p *placeOrderFormBody) ToQueriesArray() []string {

// 	var (
// 		queries      []string
// 		queriesBytes []byte
// 		queriesMap   map[string]interface{}
// 		err          error
// 	)

// 	if queriesBytes, err = json.Marshal(*p); err != nil {
// 		return queries
// 	}

// 	if err = json.Unmarshal(queriesBytes, &queriesMap); err != nil {
// 		return queries
// 	}

// 	for key, value := range queriesMap {
// 		queries = append(queries, fmt.Sprint(key, "=", value))
// 	}

// 	sort.Strings(queries)

// 	return queries
// }

// func (p *placeOrderFormBody) ToQueriesString() string {
// 	return strings.Join(p.ToQueriesArray(), "&")
// }
