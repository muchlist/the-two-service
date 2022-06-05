package service

import (
	"fetch-api/bussiness/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregatorHelper(t *testing.T) {
	data := []model.EFishData{
		{
			Province: "KALIMANTAN",
			Size:     100,
			Price:    10000,
			Time:     "2021/05/16 19:42:29",
		},
		{
			Province: "KALIMANTAN",
			Size:     10,
			Price:    20000,
			Time:     "2021/05/16 19:42:29",
		},
		{
			Province: "KALIMANTAN",
			Size:     10,
			Price:    50,
			Time:     "2022/07/16 19:42:29",
		},
		{
			Province: "BANTEN",
			Size:     15,
			Price:    15000,
			Time:     "2022/05/16 19:42:29",
		},
		{
			Province: "BANTEN",
			Size:     10,
			Price:    10,
			PriceUSD: 20,
			Time:     "2022/05/17 19:42:29",
		},
	}

	want := []model.AggregationData{
		{
			Year:     "2022",
			Week:     "20",
			Province: "BANTEN",
			Count:    2,
			Size: model.Compute{
				Maximal: 15,
				Minimal: 10,
				Median:  12.5,
				Average: 12.5,
			},
			Price: model.Compute{
				Maximal: 15000,
				Minimal: 10,
				Median:  7505,
				Average: 7505,
			},
		},
		{
			Year:     "2021",
			Week:     "19",
			Province: "KALIMANTAN",
			Count:    2,
			Size: model.Compute{
				Maximal: 100,
				Minimal: 10,
				Median:  55,
				Average: 55,
			},
			Price: model.Compute{
				Maximal: 20000,
				Minimal: 10000,
				Median:  15000,
				Average: 15000,
			},
		},
		{
			Year:     "2022",
			Week:     "28",
			Province: "KALIMANTAN",
			Count:    1,
			Size: model.Compute{
				Maximal: 10,
				Minimal: 10,
				Median:  10,
				Average: 10,
			},
			Price: model.Compute{
				Maximal: 50,
				Minimal: 50,
				Median:  50,
				Average: 50,
			},
		},
	}

	got := Aggregate(data)
	for i := range got {
		// disable usd check (because same as price)
		got[i].PriceUSD = model.Compute{}
	}

	assert.Equal(t, want, got)
}
