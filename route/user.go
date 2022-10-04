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
	public.POST("/users/register", userHandler.CreateUserTransaction)
	public.POST("/users/login", userHandler.LoginUser)

	//auth
	auth := g.Group("/auth")
	auth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims: &utils.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	auth.GET("/users/:userid/profile", userHandler.GetUserProfile)
	// auth.PUT("/users/:userid/profile", userHandler.UpdateUserProfile)
	// auth.POST("/users/:userid/address", userHandler.CreateNewAddress)
	// auth.PUT("/users/:userid/address", userHandler.UpdateAddress)
	// auth.GET("/users/:userid/address", userHandler.GetListAddress)
	// auth.GET("/users/:userid/address/:addressid", userHandler.GetAddressById)
	// auth.POST("/users/:userid/payment", userHandler.CreateNewPayment)
	// auth.PUT("/users/:userid/payment", userHandler.UpdatePayment)
	// auth.GET("/users/:userid/payment", userHandler.GetListPayment)
	// auth.GET("/users/:userid/payment/:paymentid", userHandler.GetPaymentById)
}