package middleware

import (
	"github.com/jonstodle/lets-go/hello_who/src/clock"
	"github.com/jonstodle/lets-go/hello_who/src/database"
	"github.com/labstack/echo/v4"
)

func AddDependencies(db database.Database, clk clock.Clock) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("database", db)
			c.Set("clock", clk)

			return next(c)
		}
	}
}
