package route

import (
	h "ems-aquadev/handler"
	"ems-aquadev/utils"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ProductRoutes(g *echo.Group, productHandler *h.ProductHandler) {
	//public
	public := g.Group("/public")
	public.GET("/products", productHandler.GetListProducts)
	public.GET("/products/categories",productHandler.GetListCategory)
	public.GET("/products/:productid", productHandler.GetProduct)

	//auth
	auth := g.Group("/auth")
	auth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims: &utils.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	auth.POST("/admins/products", productHandler.CreateProduct)
	auth.GET("/admins/products", productHandler.GetListProducts)
	auth.GET("/admins/products/:productid", productHandler.GetProduct)
	auth.PUT("/admins/products/:productid", productHandler.UpdateProduct)
	auth.DELETE("/admins/products/:productid", productHandler.RemoveProduct)
	auth.GET("/admins/products/categories",productHandler.GetListCategory)
}