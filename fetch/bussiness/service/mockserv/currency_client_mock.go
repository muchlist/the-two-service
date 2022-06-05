package mockserv

import "github.com/stretchr/testify/mock"

type CurrencyClientMock struct {
	mock.Mock
}

const GetUSD = "GetUSDCurrency"

func (m *CurrencyClientMock) GetUSDCurrency() (float64, error) {
	args := m.Called()
	var res float64
	if args.Get(0) != nil {
		res = args.Get(0).(float64)
	}
	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}
	return res, err
}
