package gobinanceenums

import "fmt"

type Side string

func (s *Side) GetLabelName() string {
	return "side"
}

func (s *Side) GetValue() string {
	return string(*s)
}

func (s *Side) GetSide() string {
	return s.GetValue()
}

func (s *Side) SetSide(side string) {
	*s = Side(side)
}

func (s *Side) SetSideSELL() {
	*s = SIDE_SELL
}

func (s *Side) SetSideBUY() {
	*s = SIDE_BUY
}

func (s *Side) GetQueryString() string {
	return fmt.Sprint(s.GetLabelName(), "=", s.GetValue())
}

func (s *Side) CheckENUM() error {

	switch s.GetValue() {
	case SIDE_BUY, SIDE_SELL:
		return nil
	default:
		return fmt.Errorf("invalid %s: %s", s.GetLabelName(), s.GetValue())
	}

}
