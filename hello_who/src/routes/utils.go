package routes

import (
	"github.com/jonstodle/lets-go/hello_who/src/clock"
	"github.com/jonstodle/lets-go/hello_who/src/database"
	"github.com/labstack/echo/v4"
)

func getDB(c echo.Context) database.Database {
	if db, ok := c.Get("database").(database.Database); ok {
		return db
	} else {
		panic("Database must be added to context")
	}
}

func getClock(c echo.Context) clock.Clock {
	if clk, ok := c.Get("clock").(clock.Clock); ok {
		return clk
	} else {
		panic("Clock must be added to context")
	}
}
