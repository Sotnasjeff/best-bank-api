package util

const (
	USD = "USD"
	EUR = "EUR"
	BRL = "BRL"
	GBP = "GBP"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, BRL, GBP:
		return true
	}
	return false
}
