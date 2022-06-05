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

	client := resty.New().
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetTimeout(10*time.Second).
		SetHeader("Accept", "application/json").
		SetRetryCount(2).
		SetRetryWaitTime(1 * time.Second)

	var result []model.EFishDTO

	resp, err := client.R().
		SetResult(&result).
		Get(c.URL)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		if resp.StatusCode() >= 500 {
			return nil, fmt.Errorf("%w : %s", ErrInternalServerFish, resp.String())
		}
		return nil, fmt.Errorf("%w : %s", ErrServerFish, resp.String())
	}
	return result, nil
}
