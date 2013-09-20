package gomoney

import "testing"

func TestNewMoney(t *testing.T) {
	m := NewMoney(btc())
	if m == nil {
		t.Fail()
	}
}

func TestNewMoneyWithAmount(t *testing.T) {
	m := NewMoneyWithAmount(1.337, btc())
	if m == nil {
		t.Fail()
	}
}

func btc() *currency {
	return NewCurrency(
		"BTC",
		"Bitcoin",
		"Ƀ",
		"Satoshi",
		1000000,
		",",
		".")
}
