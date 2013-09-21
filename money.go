package gomoney

type money struct {
	Value    float64
	currency *currency
}

func NewMoney(c *currency) *money {
	return NewMoneyWithAmount(0, c)
}

func NewMoneyWithAmount(a float64, c *currency) *money {
	return &money{a, c}
}

func (m *money) Subunits() float64 {
	if m.Value == 0 {
		return 0
	}
	return m.Value / float64(m.currency.subunitToUnit)
}
