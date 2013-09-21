package gomoney

import "fmt"

type Money struct {
	Value    float64
	Currency *Currency
}

func NewMoney(c *Currency) *Money {
	return NewMoneyWithValue(0, c)
}

func NewMoneyWithValue(a float64, c *Currency) *Money {
	return &Money{a, c}
}

func (m *Money) Subunits() float64 {
	if m.Value == 0 {
		return 0
	}
	return m.Value / float64(m.Currency.subunitToUnit)
}

func (m *Money) Symbol() string {
	return m.Currency.symbol
}

func (m *Money) String() string {
	if m.Currency.symbolPrecedesValue == true {
		return fmt.Sprintf("%s%g", m.Currency.symbol, m.Value)
	}
	return fmt.Sprintf("%g%s", m.Value, m.Currency.symbol)
}
