package expvarcollector

import "expvar"

var (
	cacheCurrencyHit  = expvar.NewInt("cache_currency_hit")
	cacheCurrencyMiss = expvar.NewInt("cache_currency_miss")
	cacheFishHit      = expvar.NewInt("cache_fish_hit")
	cacheFishMiss     = expvar.NewInt("cache_fish_miss")
)

func AddCurrencyHit() {
	cacheCurrencyHit.Add(1)
}

func AddCurrencyMiss() {
	cacheCurrencyMiss.Add(1)
}

func AddFishHit() {
	cacheFishHit.Add(1)
}

func AddFishMiss() {
	cacheFishMiss.Add(1)
}
