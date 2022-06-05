package service

import (
	"fetch-api/bussiness/model"
	"fetch-api/pkg/conv"
	"fetch-api/pkg/slicer"
	"fmt"
	"strings"
)

func Aggregate(dataList []model.EFishData) []model.AggregationData {
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

	return result
}
