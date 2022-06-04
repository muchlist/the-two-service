package service

import (
	"errors"
	"fetch-api/bussiness/model"
	"fetch-api/bussiness/repository"
	"fmt"
)

type FetchFishServiceAssumer interface {
	FetchData() ([]model.EFishData, error)
	GetAggregatedData() ([]model.AggregationData, error)
}

func NewFetchFishServiceAssumer(
	fishClient repository.FishApiCaller,
	currClient repository.CurrencyApiCaller,
	cacheStore repository.CurrencyStorer,
) FetchFishServiceAssumer {
	return &FetchService{
		FishClient: fishClient,
		CurrClient: currClient,
		CacheStore: cacheStore,
	}
}

type FetchService struct {
	FishClient repository.FishApiCaller
	CurrClient repository.CurrencyApiCaller
	CacheStore repository.CurrencyStorer
}

const currency = "USD"

func (f *FetchService) FetchData() ([]model.EFishData, error) {
	// get all fish commodity data
	fishDataList, err := f.FishClient.GetFish()
	if err != nil {
		return nil, fmt.Errorf("error get fish data: %w", err)
	}

	// sanitaze data, remove item without uuid
	fishdataFiltered := model.Sanitize(fishDataList)

	// get dollar value to idr from cache
	usdScale, err := f.CacheStore.GetCurrency(currency)
	if err != nil {
		if errors.Is(repository.ErrCacheNotFound, err) {

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
