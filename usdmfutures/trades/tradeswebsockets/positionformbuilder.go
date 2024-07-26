package tradeswebsockets

type positionFormBuilder struct {
	form PositionForm
}

func (p positionFormBuilder) buildMap() (map[string]interface{}, error) {

	var dstMap map[string]interface{} = make(map[string]interface{})

	if p.form.GetSymbol() != "" {
		dstMap["symbol"] = p.form.GetSymbol()
	}

	if p.form.GetRecvWindow() != 0 {
		dstMap["recvWindow"] = p.form.GetRecvWindow()
	}

	if p.form.GetTimestamp() != 0 {
		dstMap["timestamp"] = p.form.GetTimestamp()
	}

	return dstMap, nil

}
