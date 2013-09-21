package gomoney

type currency struct {
	isoCode             string
	name                string
	symbol              string
	symbolPrecedesValue bool
	subunit             string
	subunitToUnit       int
	separator           string
	delimiter           string
}

func NewCurrency(isoCode, name, symbol string, symbolPrecedesValue bool,
	subunit string, subunitToUnit int, separator, delimiter string) *currency {
	return &currency{isoCode, name, symbol, symbolPrecedesValue, subunit,
		subunitToUnit, separator, delimiter}
}
