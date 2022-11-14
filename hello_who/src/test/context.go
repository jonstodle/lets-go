package test

import (
	"github.com/jonstodle/lets-go/hello_who/src/clock/clktest"
	"github.com/jonstodle/lets-go/hello_who/src/database/dbtest"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
)

func NewContext(req *http.Request, rec *httptest.ResponseRecorder) (ctx echo.Context, db *dbtest.Database, clk *clktest.Clock) {
	db = &dbtest.Database{}
	clk = &clktest.Clock{}
	ctx = echo.New().NewContext(req, rec)
	ctx.Set("database", db)
	ctx.Set("clock", clk)

	return ctx, db, clk
}
