package repository

import (
	"crypto/tls"
	"errors"
	"fetch-api/bussiness/model"
	"fetch-api/conf"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

var ErrInternalServerCurrency = errors.New("panic : internal server error from currency api")
var ErrServerCurrency = errors.New("response error from currency api")

type CurrencyApi struct {
	URL string
}

func NewCurrencyApiCaller(config conf.Config) CurrencyApiCaller {
	return &CurrencyApi{
		URL: config.CurrencyURL,
	}
}

// GetUSDCurrency implements CurrencyApiCaller
func (c *CurrencyApi) GetUSDCurrency() (float64, error) {

	if c.URL == "" {
		return 0, errors.New("url for get usd currency is empty")
	}

	client := resty.New().
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetTimeout(10*time.Second).
		SetHeader("Accept", "application/json").
		SetRetryCount(2).
		SetRetryWaitTime(1 * time.Second)

	var result model.CurrencyDTO

	resp, err := client.R().
		SetResult(&result).
		Get(c.URL)

	if err != nil {
		return 0, err
	}

	if resp.StatusCode() != http.StatusOK {
		if resp.StatusCode() >= 500 {
			return 0, fmt.Errorf("%w : %s", ErrInternalServerCurrency, resp.String())
		}
		return 0, fmt.Errorf("%w : %s", ErrServerCurrency, resp.String())
	}
	return result.IDRToUSD, nil
}
