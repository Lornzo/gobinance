package gobinanceenums

type PriceProtect bool

func (p *PriceProtect) GetLabelName() string {
	return "priceProtect"
}

func (p *PriceProtect) GetValue() bool {
	return bool(*p)
}

func (p *PriceProtect) GetPriceProtect() string {

	if p.GetValue() {
		return "TRUE"
	}

	return "FALSE"

}

func (p *PriceProtect) SetPriceProtect(priceProtect bool) {
	*p = PriceProtect(priceProtect)
}

func (p *PriceProtect) SetPriceProtectTRUE() {
	*p = PriceProtect(true)
}

func (p *PriceProtect) SetPriceProtectFALSE() {
	*p = PriceProtect(false)
}

func (p *PriceProtect) GetQueryString() string {
	return p.GetLabelName() + "=" + p.GetPriceProtect()
}

func (p *PriceProtect) CheckENUM() error {
	return nil
}
