package model

type EFishData struct {
	UUID      string  `json:"uuid"`
	Commodity string  `json:"commodity"`
	Province  string  `json:"province"`
	City      string  `json:"city"`
	Size      float64 `json:"size"`
	Price     float64 `json:"price"`
	PriceUSD  float64 `json:"price_usd"`
	Time      string  `json:"time_parsing"`
	Timestamp string  `json:"timestamp"`
}
