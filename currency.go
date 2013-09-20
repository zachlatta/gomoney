package gomoney

type currency struct {
	isoCode       string
	name          string
	symbol        string
	subunit       string
	subunitToUnit int
	separator     string
	delimiter     string
}

func NewCurrency(isoCode, name, symbol, subunit string, subunitToUnit int,
	separator, delimiter string) *currency {
	return &currency{isoCode, name, symbol, subunit, subunitToUnit, separator,
		delimiter}
}
