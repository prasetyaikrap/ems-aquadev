package route

import (
	h "ems-aquadev/handler"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(g *echo.Group, productHandler *h.ProductHandler) {
	//public
	// public := g.Group("/public")

	//auth
	admin := g.Group("/auth")
	admin.POST("/admins/products", productHandler.CreateProduct)
	admin.GET("/admins/products", productHandler.FindListProducts)
	admin.PUT("/admins/products/:productid", productHandler.UpdateProduct)
	admin.DELETE("/admins/products/:productid", productHandler.RemoveProduct)
	admin.GET("/admins/products/categories",productHandler.FindListCategory)
	
}