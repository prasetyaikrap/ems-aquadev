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
	public.POST("/admins/register", userHandler.RegisterAdmin)
	public.POST("/admins/login", userHandler.LoginAdmin)

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

	auth.POST("/users/:userid/payment", userHandler.CreateUserPayment)
	auth.GET("/users/:userid/payment", userHandler.GetListPayments)
	auth.GET("/users/:userid/payment/:paymentid", userHandler.GetPayment)
	auth.PUT("/users/:userid/payment/:paymentid", userHandler.UpdatePayment)
	auth.DELETE("/users/:userid/payment/:paymentid", userHandler.SetDeletedPayment)

	//Shopping
	auth.GET("/users/:userid/cart", userHandler.GetCartSession)
	auth.POST("/users/:userid/cart/:sessionid/item", userHandler.AddItemToCart)
	auth.GET("/users/:userid/cart/:sessionid/item", userHandler.GetItemsCart)
	auth.DELETE("/users/:userid/cart/:sessionid/item/:itemid", userHandler.DeleteItemFromCart)
	auth.POST("/users/:userid/order/create/:sessionid", userHandler.CreateOrder)
	auth.GET("/users/:userid/order", userHandler.GetListOrders)
	auth.GET("/users/:userid/order/:orderid", userHandler.GetOrderById)
	auth.PUT("/users/:userid/order/:orderid/payment/:paymentid", userHandler.UploadReceipt)
}