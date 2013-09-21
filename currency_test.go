package gomoney

import "testing"

func TestNewCurrency(t *testing.T) {
	btc := NewCurrency(
		"BTC",
		"Bitcoin",
		"Ƀ",
		false,
		"Satoshi",
		1000000,
		",",
		".")
	if btc == nil {
		t.Fail()
	}
}
