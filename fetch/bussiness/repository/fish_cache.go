package repository

import (
	"fetch-api/bussiness/model"
	"time"

	"github.com/patrickmn/go-cache"
)

type FishCache struct {
	Cache *cache.Cache
}

func NewFishStorer() FishStorer {
	return &FishCache{
		Cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (f *FishCache) SetFish(url string, data []model.EFishDTO) error {
	f.Cache.Set(url, data, 5*time.Minute)
	return nil
}

func (f *FishCache) GetFish(url string) ([]model.EFishDTO, error) {
	value, found := f.Cache.Get(url)
	if found {
		return value.([]model.EFishDTO), nil
	}

	return nil, ErrCacheNotFound
}

func (f *FishCache) ClearFish(url string) error {
	f.Cache.Delete(url)
	return nil
}
