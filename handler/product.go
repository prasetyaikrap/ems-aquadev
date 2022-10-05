package handler

import (
	md "ems-aquadev/models"
	svc "ems-aquadev/service"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService *svc.ProductService
}

func NewProductHandler(productService *svc.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

//Product
func (handler ProductHandler) CreateProduct(c echo.Context) error {
	pd := md.CreateProductReq{}
	if err := c.Bind(&pd); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}
	productRes, err := handler.productService.CreateProduct(pd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, md.HttpResponse{
			Code: http.StatusCreated,
			Message: "Product Created Successfully",
			Data: productRes,
		})
}
func (handler ProductHandler) GetListProducts(c echo.Context) error {
	priceMin, _ := strconv.Atoi(c.QueryParam("price-min"))
	priceMax, _ := strconv.Atoi(c.QueryParam("price-max"))
	categoryid, _ := strconv.Atoi(c.QueryParam("category"))
	queries := md.ProductQueries{
		CategoryID: uint(categoryid),
		Price: []uint{
			uint(priceMin),
			uint(priceMax),
		},
	}
	listProducts, err := handler.productService.GetListProducts(queries)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	if len(listProducts) <= 0 {
		return c.JSON(http.StatusOK, md.HttpResponse{
			Code: http.StatusOK,
			Message: "Success With Empty Data",
			Data: listProducts,
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
		Code: http.StatusOK,
		Message: "List Products Found",
		Data: listProducts,
	})
}
func (handler ProductHandler) GetProduct(c echo.Context) error {
	productid, _ := strconv.Atoi(c.Param("productid"))
	product, err := handler.productService.GetProduct(uint(productid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	if product.ID == 0 {
		return c.JSON(http.StatusOK, md.HttpResponse{
			Code: http.StatusOK,
			Message: "Success With Empty Data",
			Data: product,
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
		Code: http.StatusOK,
		Message: "Product Found",
		Data: product,
	})
}
func (handler ProductHandler) UpdateProduct(c echo.Context) error {
	productid, _ := strconv.Atoi(c.Param("productid"))
	pd := md.UpdateProductReq{}
	if err := c.Bind(&pd); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}
	if err := handler.productService.UpdateProduct(pd, uint(productid)); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
		Code: http.StatusOK,
		Message: "Product Updated Successfully",
	})
}
func (handler ProductHandler) RemoveProduct(c echo.Context) error {
	productid, _ := strconv.Atoi(c.Param("productid"))
	if err := handler.productService.SetDeleteProduct(uint(productid)); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Delete Product Failed",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
		Code: http.StatusOK,
		Message: "Product Removed From Catalog",
	})
}

//Product Category
func (handler ProductHandler) GetListCategory(c echo.Context) error {
	listCategory, err := handler.productService.FindListCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	if len(listCategory) <= 0 {
		return c.JSON(http.StatusOK, md.HttpResponse{
			Code: http.StatusOK,
			Message: "Success With Empty Data",
			Data: listCategory,
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
		Code: http.StatusOK,
		Message: "List Categories Found",
		Data: listCategory,
	})
}