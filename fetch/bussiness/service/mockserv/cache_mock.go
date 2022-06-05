package mockserv

import "github.com/stretchr/testify/mock"

type CacheMock struct {
	mock.Mock
}

// SetCurrency(code string, value float64) error
// GetCurrency(code string) (float64, error)
// ClearCurrency(code string) error
const SetCurrency = "SetCurrency"

func (m *CacheMock) SetCurrency(code string, value float64) error {
	args := m.Called(code, value)
	var err error
	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}
	return err
}

const GetCurrency = "GetCurrency"

func (m *CacheMock) GetCurrency(code string) (float64, error) {
	args := m.Called(code)
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

func (m *CacheMock) ClearCurrency(code string) error {
	return nil
}
