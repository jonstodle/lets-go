package main

import (
	"github.com/jonstodle/lets-go/hello_who/src/clock"
	"github.com/jonstodle/lets-go/hello_who/src/database"
	"github.com/jonstodle/lets-go/hello_who/src/middleware"
	"github.com/jonstodle/lets-go/hello_who/src/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Use(setUpDeps())

	routes.RegisterPersonsRoutes(app.Group("/persons"))
	routes.RegisterEpisodesRoutes(app.Group("/episodes"))
	routes.RegisterSeriesRoutes(app.Group("/series"))

	err := app.Start(":1323")
	if err != nil {
		panic(err)
	}
}

func setUpDeps() echo.MiddlewareFunc {
	db, err := database.Connect("doctor-who.sqlite")
	if err != nil {
		panic(err)
	}

	clk := clock.New()

	return middleware.AddDependencies(db, clk)
}
