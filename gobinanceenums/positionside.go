package gobinanceenums

import "fmt"

type PositionSide string

func (p *PositionSide) GetLabelName() string {
	return "positionSide"
}

func (p *PositionSide) GetValue() string {
	return string(*p)
}

func (p *PositionSide) GetPositionSide() string {
	return p.GetValue()
}

func (p *PositionSide) SetPositionSide(positionSide string) {
	*p = PositionSide(positionSide)
}

func (p *PositionSide) SetPositionSideBOTH() {
	*p = PositionSide(POSITION_SIDE_BOTH)
}

func (p *PositionSide) SetPositionSideLONG() {
	*p = PositionSide(POSITION_SIDE_LONG)
}

func (p *PositionSide) SetPositionSideSHORT() {
	*p = PositionSide(POSITION_SIDE_SHORT)
}

func (p *PositionSide) GetQueryString() string {
	return fmt.Sprint(p.GetLabelName(), "=", p.GetValue())
}

func (p *PositionSide) CheckENUM() error {

	switch p.GetValue() {
	case POSITION_SIDE_BOTH, POSITION_SIDE_LONG, POSITION_SIDE_SHORT:
		return nil
	default:
		return fmt.Errorf("invalid %s: %s", p.GetLabelName(), p.GetValue())
	}

}
