package binancewebsockets

type BinanceRequest struct {
	Method string
	Params interface{}
}

func (b BinanceRequest) GetMethod() string {
	return b.Method
}

func (b BinanceRequest) GetParams() interface{} {
	return b.Params
}
