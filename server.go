package main

import (
	"ems-aquadev/config"
	"ems-aquadev/handler"
	"ems-aquadev/repository"
	"ems-aquadev/route"
	"ems-aquadev/service"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	//Initialize Database
	config.Database()
	config.Migrate()

	//echo instances
	e := echo.New()

	//Route Groups
	g := e.Group("/api/v1")

	//Initialize Handler
	userRepository := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	route.UserRoutes(g, userHandler)
	

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}

func HomeAPI(c echo.Context) error{
	return c.String(http.StatusOK, "HomeAPI is Active")
}