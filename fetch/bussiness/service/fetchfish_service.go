package service

import (
	"errors"
	"fetch-api/bussiness/model"
	"fetch-api/bussiness/repository"
	"fetch-api/pkg/conv"
	"fetch-api/pkg/slicer"
	"fmt"
	"strings"
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

	// data group contains data slice by province and year-week
	dataGroup := make(map[string][]model.EFishData)

	for _, value := range dataList {
		date := conv.ParseDate(value.Time)

		// date.ISOWeek() return year and week
		year, week := date.ISOWeek()
		key := fmt.Sprintf("%s-%d-%d", value.Province, year, week)

		// grouping by province , year-week
		dataGroup[key] = append(dataGroup[key], value)
	}

	// container for result
	var result []model.AggregationData

	// loop data map (create result per key)
	for key, dataList := range dataGroup {
		keySlice := strings.Split(key, "-") // return example Banjarmasin-2020-8
		province := keySlice[0]
		year := keySlice[1]
		week := keySlice[2]

		// skip if date format incorectly
		if year == "1" {
			continue
		}

		var sizes []float64
		var prices []float64
		var usdPrices []float64

		for _, data := range dataList {
			sizes = append(sizes, data.Size)
			prices = append(prices, data.Price)
			usdPrices = append(usdPrices, data.PriceUSD)
		}

		medianSize := slicer.Median(sizes)
		maxSize := slicer.Max(sizes)
		minSize := slicer.Min(sizes)
		averageSize := slicer.Average(sizes)

		medianPrice := slicer.Median(prices)
		maxPrice := slicer.Max(prices)
		minPrice := slicer.Min(prices)
		averagePrice := slicer.Average(prices)

		medianPriceUsd := slicer.Median(usdPrices)
		maxPriceUsd := slicer.Max(usdPrices)
		minPriceUsd := slicer.Min(usdPrices)
		averagePriceUsd := slicer.Average(usdPrices)

		result = append(result, model.AggregationData{
			Year:     year,
			Week:     week,
			Province: province,
			Count:    len(dataList),
			Size: model.Compute{
				Maximal: maxSize,
				Minimal: minSize,
				Median:  medianSize,
				Average: averageSize,
			},
			Price: model.Compute{
				Maximal: maxPrice,
				Minimal: minPrice,
				Median:  medianPrice,
				Average: averagePrice,
			},
			PriceUSD: model.Compute{
				Maximal: maxPriceUsd,
				Minimal: minPriceUsd,
				Median:  medianPriceUsd,
				Average: averagePriceUsd,
			},
		})

	}

	return result, err
}
