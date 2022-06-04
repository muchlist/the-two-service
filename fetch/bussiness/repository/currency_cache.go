package repository

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

var ErrCacheNotFound = errors.New("cache not found")

type CurrencyCache struct {
	Cache *cache.Cache
}

func NewCurrencyStorer() CurrencyStorer {
	return &CurrencyCache{
		Cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

// ConvertIDRToUSD implements CurrencyCaller
func (c *CurrencyCache) SetCurrency(code string, value float64) error {
	c.Cache.Set(code, value, 5*time.Minute)
	return nil
}

func (c *CurrencyCache) GetCurrency(code string) (float64, error) {
	value, found := c.Cache.Get(code)
	if found {
		return value.(float64), nil
	}

	return 0, ErrCacheNotFound
}

func (c *CurrencyCache) ClearCurrency(code string) error {
	c.Cache.Delete(code)
	return nil
}
