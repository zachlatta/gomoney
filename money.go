package gomoney

import "fmt"

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

func (m *money) Symbol() string {
	return m.currency.symbol
}

func (m *money) String() string {
	if m.currency.symbolPrecedesValue == true {
		return fmt.Sprintf("%s%g", m.currency.symbol, m.Value)
	}
	return fmt.Sprintf("%g%s", m.Value, m.currency.symbol)
}
