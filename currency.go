package gomoney

type Currency struct {
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
	subunit string, subunitToUnit int, separator, delimiter string) *Currency {
	return &Currency{isoCode, name, symbol, symbolPrecedesValue, subunit,
		subunitToUnit, separator, delimiter}
}
