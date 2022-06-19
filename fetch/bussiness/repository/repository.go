// package repository serve data from outside
// also data from outside which is already stored in cache
package repository

import "fetch-api/bussiness/model"

// CurrencyStorer used to cache value of USD currency
type CurrencyStorer interface {
	SetCurrency(code string, value float64) error
	GetCurrency(code string) (float64, error)
	ClearCurrency(code string) error
}

// FishStorer used to cache fish data
type FishStorer interface {
	SetFish(url string, value []model.EFishDTO) error
	GetFish(url string) ([]model.EFishDTO, error)
	ClearFish(url string) error
}

// CurrencyApiCaller used to get recent value of convert scale RP to USD
type CurrencyApiCaller interface {
	GetUSDCurrency() (float64, error)
}

type FishApiCaller interface {
	GetFish() ([]model.EFishDTO, error)
}
