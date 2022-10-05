package route

import (
	h "ems-aquadev/handler"
	"ems-aquadev/utils"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRoutes(g *echo.Group, userHandler *h.UserHandler) {
	//public
	public := g.Group("/public")
	public.POST("/users/register", userHandler.CreateUser)
	public.POST("/users/login", userHandler.LoginUser)

	//auth
	auth := g.Group("/auth")
	auth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims: &utils.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	auth.GET("/users/:userid/profile", userHandler.GetUserProfile)
	auth.PUT("/users/:userid/profile", userHandler.UpdateUserProfile)
	auth.POST("/users/:userid/addresses", userHandler.CreateUserAddress)
	auth.GET("/users/:userid/addresses", userHandler.GetListAddress)
	auth.GET("/users/:userid/addresses/:addressid", userHandler.GetAddress)
	auth.PUT("/users/:userid/addresses/:addressid", userHandler.UpdateAddress)
	auth.DELETE("/users/:userid/addresses/:addressid", userHandler.SetDeletedAddress)
	// auth.POST("/users/:userid/payment", userHandler.CreateNewPayment)
	// auth.PUT("/users/:userid/payment", userHandler.UpdatePayment)
	// auth.GET("/users/:userid/payment", userHandler.GetListPayment)
	// auth.GET("/users/:userid/payment/:paymentid", userHandler.GetPaymentById)
}