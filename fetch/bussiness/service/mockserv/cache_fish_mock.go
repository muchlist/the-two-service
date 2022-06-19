package mockserv

import (
	"fetch-api/bussiness/model"

	"github.com/stretchr/testify/mock"
)

type CacheFishMock struct {
	mock.Mock
}

// SetFish(code string, value float64) error
// GetFish(code string) (float64, error)
// ClearFish(code string) error
const SetFish = "SetFish"

func (m *CacheFishMock) SetFish(code string, value []model.EFishDTO) error {
	args := m.Called(code, value)
	var err error
	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}
	return err
}

const GetFishC = "GetFish"

func (m *CacheFishMock) GetFish(url string) ([]model.EFishDTO, error) {
	args := m.Called(url)
	var res []model.EFishDTO
	if args.Get(0) != nil {
		res = args.Get(0).([]model.EFishDTO)
	}
	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}
	return res, err
}

func (m *CacheFishMock) ClearFish(code string) error {
	return nil
}
