package clktest

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type Clock struct {
	mock.Mock
}

func (c *Clock) Today(timezone string) (time.Time, error) {
	args := c.Called(timezone)
	return args.Get(0).(time.Time), args.Error(1)
}
