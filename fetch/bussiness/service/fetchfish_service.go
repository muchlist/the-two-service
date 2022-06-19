package service

import (
	"errors"
	"fetch-api/bussiness/model"
	"fetch-api/bussiness/repository"
	"fetch-api/conf"
	expvarcollector "fetch-api/pkg/expvar_collector"
	"fmt"
)

type FetchFishServiceAssumer interface {
	FetchData() ([]model.EFishData, error)
	GetAggregatedData() ([]model.AggregationData, error)
}

func NewFetchFishServiceAssumer(
	cfg conf.Config,
	fishClient repository.FishApiCaller,
	currClient repository.CurrencyApiCaller,
	cacheStore repository.CurrencyStorer,
	cacheFishStore repository.FishStorer,
) FetchFishServiceAssumer {
	return &FetchService{
		FishClient:     fishClient,
		CurrClient:     currClient,
		CacheStore:     cacheStore,
		CacheFishStore: cacheFishStore,
	}
}

type FetchService struct {
	Cfg            conf.Config
	FishClient     repository.FishApiCaller
	CurrClient     repository.CurrencyApiCaller
	CacheStore     repository.CurrencyStorer
	CacheFishStore repository.FishStorer
}

const currency = "USD"

func (f *FetchService) FetchData() ([]model.EFishData, error) {
	// get all fish commodity data from cache then if not exist
	// get from call api
	fishDataList, err := f.CacheFishStore.GetFish(f.Cfg.ResourceURL)
	if err != nil {
		if errors.Is(repository.ErrCacheNotFound, err) {
			// cache miss
			expvarcollector.AddFishMiss()

			// get from fish api
			fishDataList, err = f.FishClient.GetFish()
			if err != nil {
				return nil, fmt.Errorf("error get fish data: %w", err)
			}
			// set cache
			_ = f.CacheFishStore.SetFish(f.Cfg.ResourceURL, fishDataList)
		} else {
			return nil, fmt.Errorf("error get fish data from cache: %w", err)
		}
	} else {
		// cache hit
		expvarcollector.AddFishHit()
	}

	// sanitaze data, remove item without uuid
	fishdataFiltered := model.Sanitize(fishDataList)

	// get dollar value to idr from cache
	usdScale, err := f.CacheStore.GetCurrency(currency)
	if err != nil {
		if errors.Is(repository.ErrCacheNotFound, err) {
			// cache miss
			expvarcollector.AddCurrencyMiss()

			// get from currency api
			usdScale, err = f.CurrClient.GetUSDCurrency()
			if err != nil {
				return nil, fmt.Errorf("error get currency data: %w", err)
			}
			// set cache
			_ = f.CacheStore.SetCurrency(currency, usdScale)
		} else {
			return nil, fmt.Errorf("error get currency data from cache: %w", err)
		}
	} else {
		// cache hit
		expvarcollector.AddCurrencyHit()
	}

	// insert dollar value to fish data
	result := make([]model.EFishData, 0, len(fishdataFiltered))
	for _, v := range fishdataFiltered {
		// To_Domain method fill usd price with scale inputed
		result = append(result, v.ToDomain(usdScale))
	}

	return result, nil
}

// GetAggregatedData will call fetchdata() to get data with same logic
// then data will be proceced
func (f *FetchService) GetAggregatedData() ([]model.AggregationData, error) {
	dataList, err := f.FetchData()
	if err != nil {
		return nil, fmt.Errorf("error fetch fish data: %w", err)
	}

	aggregatedData := Aggregate(dataList)

	return aggregatedData, err
}
