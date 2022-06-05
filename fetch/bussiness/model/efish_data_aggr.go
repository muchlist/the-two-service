package model

type AggregationData struct {
	Year     string  `json:"year"`
	Week     string  `json:"week"`
	Province string  `json:"province"`
	Count    int     `json:"data_count"`
	Size     Compute `json:"size"`
	Price    Compute `json:"price"`
	PriceUSD Compute `json:"price_usd"`
}

type Compute struct {
	Maximal float64 `json:"maximal"`
	Minimal float64 `json:"minimal"`
	Median  float64 `json:"median"`
	Average float64 `json:"average"`
}
