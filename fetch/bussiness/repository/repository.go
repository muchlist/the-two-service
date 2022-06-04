// package repository serve data from outside
// also data from outside which is already stored in cache
package repository

type CurrencyStorer interface {
	SetCurrency(code string, value float64) error
	GetCurrency(code string) (float64, error)
	ClearCurrency(code string) error
}
