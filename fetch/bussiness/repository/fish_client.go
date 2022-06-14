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

var ErrInternalServerFish = errors.New("panic : internal server error from fish api")
var ErrServerFish = errors.New("response error from fish api")

type FishApi struct {
	URL string
}

func NewFishApiCaller(config conf.Config) FishApiCaller {
	return &FishApi{
		URL: config.ResourceURL,
	}
}

// GetFish implements FishApiCaller
func (c *FishApi) GetFish() ([]model.EFishDTO, error) {

	if c.URL == "" {
		return nil, errors.New("url for getting fish data is empty")
	}

	// DO in Hystrix Circuit Breaker
	output := make(chan []model.EFishDTO, 1)
	userErr := make(chan error, 1)
	errors := hystrix.Go("get_fish", func() error {

		// call http rest
		client := resty.New().
			SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
			SetTimeout(10*time.Second).
			SetHeader("Accept", "application/json")

		var result []model.EFishDTO

		resp, err := client.R().
			SetResult(&result).
			Get(c.URL)

		if err != nil {
			return err
		}

		if resp.StatusCode() != http.StatusOK {
			if resp.StatusCode() >= 500 {
				return fmt.Errorf("%w : %s", ErrInternalServerFish, resp.String())
			}
			// if not err 500, it is user error
			userErr <- fmt.Errorf("%w : %s", ErrServerFish, resp.String())
			return nil
		}
		output <- result
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return nil, err
	case err := <-userErr:
		return nil, err
	}

}
