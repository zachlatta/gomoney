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

func TestSubunits(t *testing.T) {
	m := NewMoneyWithAmount(100000000, btc())
	if m.Subunits() != float64(1) {
		t.Fail()
	}
}

func TestSymbol(t *testing.T) {
	m := NewMoneyWithAmount(14, btc())
	if m.Symbol() != "Ƀ" {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	m := NewMoneyWithAmount(42, btc())
	if m.String() != "42Ƀ" {
		t.Fail()
	}
}

func btc() *Currency {
	return NewCurrency(
		"BTC",
		"Bitcoin",
		"Ƀ",
		false,
		"Satoshi",
		100000000,
		",",
		".")
}
