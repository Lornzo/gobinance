package gobinanceenums

import "fmt"

type WorkingType string

func (w *WorkingType) GetLabelName() string {
	return "workingType"
}

func (w *WorkingType) GetValue() string {
	return string(*w)
}

func (w *WorkingType) GetWorkingType() string {
	return w.GetValue()
}

func (w *WorkingType) SetWorkingType(workingType string) {
	*w = WorkingType(workingType)
}

func (w *WorkingType) SetWorkingTypeMARKPRICE() {
	*w = WorkingType(WORKING_TYPE_MARK_PRICE)
}

func (w *WorkingType) SetWorkingTypeCONTRACTPRICE() {
	*w = WorkingType(WORKING_TYPE_CONTRACT_PRICE)
}

func (w *WorkingType) GetQueryString() string {
	return w.GetLabelName() + "=" + w.GetValue()
}

func (w *WorkingType) CheckENUM() error {

	switch w.GetValue() {
	case WORKING_TYPE_MARK_PRICE, WORKING_TYPE_CONTRACT_PRICE:
		return nil
	default:
		return fmt.Errorf("invalid %s: %s", w.GetLabelName(), w.GetValue())
	}

}
