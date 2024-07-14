package gobinanceenums

import "fmt"

type TimeInForce string

func (t *TimeInForce) GetLabelName() string {
	return "timeInForce"
}

func (t *TimeInForce) GetValue() string {
	return string(*t)
}

func (t *TimeInForce) GetTimeInForce() string {
	return t.GetValue()
}

func (t *TimeInForce) SetTimeInForce(timeInForce string) {
	*t = TimeInForce(timeInForce)
}

func (t *TimeInForce) SetTimeInForceGTC() {
	*t = TimeInForce(TIME_IN_FORCE_GTC)
}

func (t *TimeInForce) SetTimeInForceIOC() {
	*t = TimeInForce(TIME_IN_FORCE_IOC)
}

func (t *TimeInForce) SetTimeInForceFOK() {
	*t = TimeInForce(TIME_IN_FORCE_FOK)
}

func (t *TimeInForce) SetTimeInForceGTX() {
	*t = TimeInForce(TIME_IN_FORCE_GTX)
}

func (t *TimeInForce) SetTimeInForceGTD() {
	*t = TimeInForce(TIME_IN_FORCE_GTD)
}

func (t *TimeInForce) GetQueryString() string {
	return t.GetLabelName() + "=" + t.GetValue()
}

func (t *TimeInForce) CheckENUM() error {

	switch t.GetValue() {
	case TIME_IN_FORCE_GTC, TIME_IN_FORCE_IOC, TIME_IN_FORCE_FOK, TIME_IN_FORCE_GTX, TIME_IN_FORCE_GTD:
		return nil
	default:
		return fmt.Errorf("invalid %s: %s", t.GetLabelName(), t.GetValue())
	}

}
