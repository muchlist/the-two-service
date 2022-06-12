package expvarcollector

import "expvar"

var (
	cacheCurrencyHit  = expvar.NewInt("cache_currency_hit")
	cacheCurrencyMiss = expvar.NewInt("cache_currency_miss")
)

func AddCurrencyHit() {
	cacheCurrencyHit.Add(1)
}

func AddCurrencyMiss() {
	cacheCurrencyMiss.Add(1)
}
