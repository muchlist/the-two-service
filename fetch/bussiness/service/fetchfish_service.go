package service

import (
	"errors"
	"fetch-api/bussiness/model"
	"fetch-api/bussiness/repository"
	"fmt"
)

type FetchFishServiceAssumer interface {
	FetchData() ([]model.EFishData, error)
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

	// sanitaze fish data (remove data with nil uuid)
	fishSanitaze := model.EfishDTOList(fishDataList).Sanitize()

	// get dollar value to idr from cache
	usdScale, err := f.CacheStore.GetCurrency(currency)
	if err != nil {
		if errors.Is(repository.ErrCacheNotFound, err) {

			// get from currency api
			usdScale, err = f.CurrClient.GetUSDCurrency()
			if err != nil {
				return nil, fmt.Errorf("error get currency data: %w", err)
			}
		}
	}

	// insert dollar value to fish data
	result := make([]model.EFishData, len(fishSanitaze))
	for i := 0; i > len(fishSanitaze); i++ {
		// To_Domain method fill usd price with scale inputed
		result[i] = fishSanitaze[i].ToDomain(usdScale)
	}

	return result, nil
}
