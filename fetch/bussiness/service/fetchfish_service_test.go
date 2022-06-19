package service

import (
	"fetch-api/bussiness/model"
	"fetch-api/bussiness/repository"
	"fetch-api/bussiness/service/mockserv"
	"fetch-api/conf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchData_Success(t *testing.T) {

	type test struct {
		input []model.EFishDTO
		want  []model.EFishData
	}

	tests := test{
		input: []model.EFishDTO{
			{
				UUID:        "123",
				Commodity:   "Kelatau",
				Province:    "KALIMANTAN SELATAN",
				City:        "BANJARMASIN",
				Size:        "20",
				Price:       "10000",
				TimeParsing: "2022/05/16 19:42:29",
			},
			{
				UUID:        "124",
				Commodity:   "Haruan",
				Province:    "KALIMANTAN SELATAN",
				City:        "BANJARMASIN",
				Size:        "10",
				Price:       "20000",
				TimeParsing: "2022/05/16 19:42:29",
			},
		},
		want: []model.EFishData{
			{
				UUID:      "123",
				Commodity: "Kelatau",
				Province:  "KALIMANTAN SELATAN",
				City:      "BANJARMASIN",
				Size:      20,
				Price:     10000,
				PriceUSD:  5000,
				Time:      "2022/05/16 19:42:29",
			},
			{
				UUID:      "124",
				Commodity: "Haruan",
				Province:  "KALIMANTAN SELATAN",
				City:      "BANJARMASIN",
				Size:      10,
				Price:     20000,
				PriceUSD:  10000,
				Time:      "2022/05/16 19:42:29",
			},
		},
	}

	// dependency
	// fish caller mock
	fc := new(mockserv.FishClientMock)
	fc.On(mockserv.GetFish).Return(
		tests.input, nil,
	)

	// currency caller mock
	cc := new(mockserv.CurrencyClientMock)
	cc.On(mockserv.GetUSD).Return(0.5, nil)

	// cache mock
	rc := new(mockserv.CacheMock)
	rc.On(mockserv.GetCurrency, "USD").Return(0.0, repository.ErrCacheNotFound)
	rc.On(mockserv.SetCurrency, "USD", 0.5).Return(nil)

	// fake config
	cfg := conf.Config{}

	// cache fish mock
	fcache := new(mockserv.CacheFishMock)
	fcache.On(mockserv.GetFishC, cfg.ResourceURL).Return(nil, repository.ErrCacheNotFound)
	fcache.On(mockserv.SetFish, cfg.ResourceURL, tests.input).Return(nil)

	fetchService := NewFetchFishServiceAssumer(cfg, fc, cc, rc, fcache)

	got, err := fetchService.FetchData()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(got), "size returned data")
	assert.Equal(t, tests.want, got, "result")
}

func TestFetchData_CachedCurrency(t *testing.T) {

	type test struct {
		input []model.EFishDTO
		want  []model.EFishData
	}

	tests := test{
		input: []model.EFishDTO{
			{
				UUID:        "123",
				Commodity:   "Kelatau",
				Province:    "KALIMANTAN SELATAN",
				City:        "BANJARMASIN",
				Size:        "20",
				Price:       "10000",
				TimeParsing: "2022/05/16 19:42:29",
			},
			{
				UUID:        "124",
				Commodity:   "Haruan",
				Province:    "KALIMANTAN SELATAN",
				City:        "BANJARMASIN",
				Size:        "10",
				Price:       "20000",
				TimeParsing: "2022/05/16 19:42:29",
			},
		},
		want: []model.EFishData{
			{
				UUID:      "123",
				Commodity: "Kelatau",
				Province:  "KALIMANTAN SELATAN",
				City:      "BANJARMASIN",
				Size:      20,
				Price:     10000,
				PriceUSD:  5000,
				Time:      "2022/05/16 19:42:29",
			},
			{
				UUID:      "124",
				Commodity: "Haruan",
				Province:  "KALIMANTAN SELATAN",
				City:      "BANJARMASIN",
				Size:      10,
				Price:     20000,
				PriceUSD:  10000,
				Time:      "2022/05/16 19:42:29",
			},
		},
	}

	// dependency
	// fish caller mock
	fc := new(mockserv.FishClientMock)
	fc.On(mockserv.GetFish).Return(
		tests.input, nil,
	)

	// currency caller mock
	cc := new(mockserv.CurrencyClientMock)
	cc.On(mockserv.GetUSD).Return(0.5, nil)

	// cache mock
	rc := new(mockserv.CacheMock)
	rc.On(mockserv.GetCurrency, "USD").Return(0.5, nil)
	// not called because cache hit
	// rc.On(mockserv.SetCurrency, "USD", 0.5).Return(nil)

	// fake config
	cfg := conf.Config{}

	// cache fish mock
	fcache := new(mockserv.CacheFishMock)
	fcache.On(mockserv.GetFishC, "").Return(nil, repository.ErrCacheNotFound)
	fcache.On(mockserv.SetFish, "", tests.input).Return(nil)

	fetchService := NewFetchFishServiceAssumer(cfg, fc, cc, rc, fcache)

	got, err := fetchService.FetchData()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(got), "size returned data")
	assert.Equal(t, tests.want, got, "result")
}

func TestAggregateData_Success(t *testing.T) {

	type test struct {
		rawData []model.EFishDTO
		want    []model.AggregationData
	}

	tests := test{
		rawData: []model.EFishDTO{
			{
				UUID:        "123",
				Commodity:   "Kelatau",
				Province:    "KALIMANTAN SELATAN",
				City:        "BANJARMASIN",
				Size:        "20",
				Price:       "10000",
				TimeParsing: "2022/05/16 19:42:29",
			},
			{
				UUID:        "124",
				Commodity:   "Haruan",
				Province:    "KALIMANTAN SELATAN",
				City:        "BANJARMASIN",
				Size:        "10",
				Price:       "20000",
				TimeParsing: "2022/05/16 19:42:29",
			},
		},
		want: []model.AggregationData{
			{
				Year:     "2022",
				Week:     "20",
				Province: "KALIMANTAN SELATAN",
				Count:    2,
				Size: model.Compute{
					Maximal: 20,
					Minimal: 10,
					Median:  15,
					Average: 15,
				},
				Price: model.Compute{
					Maximal: 20000,
					Minimal: 10000,
					Median:  15000,
					Average: 15000,
				},
				PriceUSD: model.Compute{
					Maximal: 10000,
					Minimal: 5000,
					Median:  7500,
					Average: 7500,
				},
			},
		},
	}

	// dependency
	// fish caller mock
	fc := new(mockserv.FishClientMock)
	fc.On(mockserv.GetFish).Return(
		tests.rawData, nil,
	)

	// currency caller mock
	cc := new(mockserv.CurrencyClientMock)
	cc.On(mockserv.GetUSD).Return(0.5, nil)

	// cache mock
	rc := new(mockserv.CacheMock)
	rc.On(mockserv.GetCurrency, "USD").Return(0.0, repository.ErrCacheNotFound)
	rc.On(mockserv.SetCurrency, "USD", 0.5).Return(nil)

	// fake config
	cfg := conf.Config{}

	// cache fish mock
	fcache := new(mockserv.CacheFishMock)
	fcache.On(mockserv.GetFishC, cfg.ResourceURL).Return(nil, repository.ErrCacheNotFound)
	fcache.On(mockserv.SetFish, cfg.ResourceURL, tests.rawData).Return(nil)

	fetchService := NewFetchFishServiceAssumer(cfg, fc, cc, rc, fcache)

	got, err := fetchService.GetAggregatedData()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(got), "size returned data")
	assert.Equal(t, tests.want, got, "result")
}
