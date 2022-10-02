package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", HomeAPI)

	e.Logger.Fatal(e.Start(":1323"))
}

func HomeAPI(c echo.Context) error{
	return c.String(http.StatusOK, "HomeAPI is Active")
}