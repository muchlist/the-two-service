package mockserv

import (
	"fetch-api/bussiness/model"

	"github.com/stretchr/testify/mock"
)

// FishClientMock
type FishClientMock struct {
	mock.Mock
}

const GetFish = "GetFish"

// Our mocked FishClient method
func (m *FishClientMock) GetFish() ([]model.EFishDTO, error) {
	args := m.Called()
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
