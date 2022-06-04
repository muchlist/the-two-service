package model

type EFishData struct {
	UUID        string  `json:"uuid"`
	Commodity   string  `json:"commodity"`
	Province    string  `json:"province"`
	City        string  `json:"city"`
	Size        string  `json:"size"`
	Price       float64 `json:"price"`
	PriceUSD    float64 `json:"price_usd"`
	TimeParsing string  `json:"time_parsing"`
	Timestamp   string  `json:"timestamp"`
}
