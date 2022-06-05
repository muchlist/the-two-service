package model

import (
	"strconv"
)

// {
// 	"uuid": "1653120347900",
// 	"komoditas": "Bobara",
// 	"area_provinsi": "SULAWESI BARAT",
// 	"area_kota": "MAMUJU UTARA",
// 	"size": "30",
// 	"price": "100000",
// 	"tgl_parsed": "2022-05-21T08:05:47.900Z",
// 	"timestamp": "1653120347900"
// },

type EFishDTO struct {
	UUID        string `json:"uuid"`
	Commodity   string `json:"komoditas"`
	Province    string `json:"area_provinsi"`
	City        string `json:"area_kota"`
	Size        string `json:"size"`
	Price       string `json:"price"`
	TimeParsing string `json:"tgl_parsed"`
	Timestamp   string `json:"timestamp"`
}

func (ef *EFishDTO) ToDomain(usdScale float64) EFishData {
	priceNum, err := strconv.ParseFloat(ef.Price, 64)
	if err != nil {
		priceNum = 0
	}

	sizeNum, err := strconv.ParseFloat(ef.Size, 64)
	if err != nil {
		sizeNum = 0
	}

	priceDollar := priceNum * usdScale

	return EFishData{
		UUID:      ef.UUID,
		Commodity: ef.Commodity,
		Province:  ef.Province,
		City:      ef.City,
		Size:      sizeNum,
		Price:     priceNum,
		PriceUSD:  priceDollar,
		Time:      ef.TimeParsing,
		Timestamp: ef.Timestamp,
	}
}

// Sanitize removing data with uuid null
func Sanitize(e []EFishDTO) []EFishDTO {
	result := make([]EFishDTO, 0, len(e))

	for _, v := range e {
		if v.UUID == "" {
			continue
		}
		result = append(result, v)
	}

	return result
}
