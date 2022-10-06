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
	user := g.Group("/auth")
	user.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims: &utils.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	user.GET("/users/:userid/profile", userHandler.GetUserProfile)
	user.PUT("/users/:userid/profile", userHandler.UpdateUserProfile)
	user.POST("/users/:userid/addresses", userHandler.CreateUserAddress)
	user.GET("/users/:userid/addresses", userHandler.GetListAddress)
	user.GET("/users/:userid/addresses/:addressid", userHandler.GetAddress)
	user.PUT("/users/:userid/addresses/:addressid", userHandler.UpdateAddress)
	user.DELETE("/users/:userid/addresses/:addressid", userHandler.SetDeletedAddress)
	user.POST("/users/:userid/payment", userHandler.CreateUserPayment)
	user.GET("/users/:userid/payment", userHandler.GetListPayments)
	user.GET("/users/:userid/payment/:paymentid", userHandler.GetPayment)
	user.PUT("/users/:userid/payment/:paymentid", userHandler.UpdatePayment)
	user.DELETE("/users/:userid/payment/:paymentid", userHandler.UpdatePayment)

	//Shopping
	user.GET("/users/:userid/cart", userHandler.GetCartSession)
	user.POST("/users/:userid/cart/:sessionid/item", userHandler.AddItemToCart)
	user.GET("/users/:userid/cart/:sessionid/item", userHandler.GetItemsCart)
	user.DELETE("/users/:userid/cart/:sessionid/item/:itemid", userHandler.DeleteItemFromCart)
	user.POST("/users/:userid/order/create/:sessionid", userHandler.CreateOrder)
	user.GET("/users/:userid/order", userHandler.GetListOrders)
	user.GET("/users/:userid/order/:orderid", userHandler.GetOrderById)
	user.PUT("/users/:userid/order/:orderid/payment/:paymentid", userHandler.UploadReceipt)
}