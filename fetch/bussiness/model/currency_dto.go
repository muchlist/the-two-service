package model

// {
// 	"meta": {
// 				"last_updated_at": "2022-06-03T23:59:59Z"
// 			},
// 	"data": {
// 		"USD": {
// 				"code": "USD",
// 				"value": 0.000069
// 				}
// 			}
// 	}

type CurrencyDTO struct {
	Data     CurrencyUSD `json:"data"`
	Metadata Meta        `json:"meta"`
}

func (c CurrencyDTO) GetValue() float64 {
	return c.Data.USD.Value
}

// ================================== CurrencyDTO proferties
type CurrencyUSD struct {
	USD Currency `json:"USD"`
}

type Meta struct {
	LastUpdate string `json:"last_updated_at"`
}

type Currency struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

// ==================================
