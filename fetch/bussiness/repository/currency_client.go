package repository

import (
	"crypto/tls"
	"errors"
	"fetch-api/bussiness/model"
	"fetch-api/conf"
	"fmt"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
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

	// DO in Hystrix Circuit Breaker
	output := make(chan float64, 1)
	userErr := make(chan error, 1)
	errors := hystrix.Go("get_currency", func() error {

		// call http rest
		client := resty.New().
			SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
			SetTimeout(10*time.Second).
			SetHeader("Accept", "application/json")

		var result model.CurrencyDTO

		resp, err := client.R().
			SetResult(&result).
			Get(c.URL)

		if err != nil {
			return err
		}

		if resp.StatusCode() != http.StatusOK {
			if resp.StatusCode() >= 500 {
				return fmt.Errorf("%w : %s", ErrInternalServerCurrency, resp.String())
			}
			// if not err 500, it is user error
			userErr <- fmt.Errorf("%w : %s", ErrServerCurrency, resp.String())
			return nil
		}

		output <- result.IDRToUSD
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return 0, err
	case err := <-userErr:
		return 0, err
	}
}
