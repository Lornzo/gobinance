package gobinanceenums

type ClosePosition bool

func (c *ClosePosition) GetLabelName() string {
	return "closePosition"
}

func (c *ClosePosition) GetValue() bool {
	return bool(*c)
}

func (c *ClosePosition) GetClosePosition() string {

	if c.GetValue() {
		return "true"
	}

	return "false"

}

func (c *ClosePosition) SetClosePosition(closePosition bool) {
	*c = ClosePosition(closePosition)
}

func (c *ClosePosition) SetClosePositionTRUE() {
	*c = ClosePosition(true)
}

func (c *ClosePosition) SetClosePositionFALSE() {
	*c = ClosePosition(false)
}

func (c *ClosePosition) GetQueryString() string {
	return c.GetLabelName() + "=" + c.GetClosePosition()
}

func (c *ClosePosition) CheckENUM() error {
	return nil
}
