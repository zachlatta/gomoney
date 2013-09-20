package gomoney

type money struct {
	Value    float64
	currency *currency
}

func NewMoney(c *currency, v float64) *money {
  return &money{v, c}
}
