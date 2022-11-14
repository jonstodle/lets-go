package dbtest

import (
	"github.com/stretchr/testify/mock"
)

type Database struct {
	mock.Mock
}

func (d *Database) Get(dest any, query string, args ...any) error {
	return d.Called(dest, query, args).Error(0)
}

func (d *Database) Select(dest any, query string, args ...any) error {
	return d.Called(dest, query, args).Error(0)
}
