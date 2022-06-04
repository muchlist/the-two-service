package model

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

type EfishDTOList []EFishDTO

type EFishDTO struct {
	UUID        *string `json:"uuid"`
	Commodity   *string `json:"commodity"`
	Province    *string `json:"province"`
	City        *string `json:"city"`
	Size        *string `json:"size"`
	Price       *string `json:"price"`
	TimeParsing *string `json:"time_parsing"`
	Timestamp   *string `json:"timestamp"`
}

// Sanitize removing data with uuid null
func (e EfishDTOList) Sanitize() []EFishDTO {
	result := make([]EFishDTO, 0, len(e))

	for _, v := range e {
		if v.UUID == nil {
			continue
		}
		result = append(result, v)
	}

	return result
}
