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
