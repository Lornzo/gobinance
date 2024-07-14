package gobinanceenums

import "fmt"

type ReduceOnly bool

func (r *ReduceOnly) GetLabelName() string {
	return "reduceOnly"
}

func (r *ReduceOnly) GetValue() bool {
	return bool(*r)
}

func (r *ReduceOnly) GetReduceOnly() string {

	if r.GetValue() {
		return "true"
	}

	return "false"

}

func (r *ReduceOnly) SetReduceOnly(reduceOnly bool) {
	*r = ReduceOnly(reduceOnly)
}

func (r *ReduceOnly) SetReduceOnlyTRUE() {
	*r = ReduceOnly(true)
}

func (r *ReduceOnly) SetReduceOnlyFALSE() {
	*r = ReduceOnly(false)
}

func (r *ReduceOnly) GetQueryString() string {
	return fmt.Sprint(r.GetLabelName(), "=", r.GetReduceOnly())
}

func (r *ReduceOnly) CheckENUM() error {
	return nil
}
