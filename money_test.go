package gomoney

import "testing"

func TestNewMoney(t *testing.T) {
	btc := NewCurrency(
		"BTC",
		"Bitcoin",
		"Ƀ",
		"Satoshi",
		1000000,
		",",
		".")
  m := NewMoney(btc, 0)
  if m == nil {
    t.Fail()
  }
}
