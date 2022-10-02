package main

import (
	"ems-aquadev/config"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	//Initialize Database
	config.Database()

	//echo instances
	e := echo.New()

	e.GET("/", HomeAPI)

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}

func HomeAPI(c echo.Context) error{
	return c.String(http.StatusOK, "HomeAPI is Active")
}