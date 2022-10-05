package service

import (
	md "ems-aquadev/models"
	repo "ems-aquadev/repository"
)

type ProductService struct {
	ProductRepository *repo.ProductRepository
}

func NewProductService(productRepository *repo.ProductRepository) *ProductService {
	return &ProductService{productRepository}
}

//Product Service
//Product
func (service ProductService) CreateProduct(productReq md.CreateProductReq) (md.CreateProductRes, error) {
	//Store Product
	pd := md.Product{
		SKU: productReq.SKU,
		Name: productReq.Name,
		Description: productReq.Description,
		Price: productReq.Price,
		Quantity: productReq.Quantity,
		ImageURL: productReq.ImageURL,
	}
	category, err := service.ProductRepository.FindOrCreateCategory(productReq.ProductCategory)
	if err != nil {
		return md.CreateProductRes{}, err
	}
	pd.CategoryID = category.ID
	product, err := service.ProductRepository.StoreProduct(pd)
	if err != nil {
		return md.CreateProductRes{}, err
	}
	//Response
	productRes := md.CreateProductRes {
		ID: product.ID,
		SKU: product.SKU,
		Name: product.Name,
	}
	return productRes, nil
}
func (service ProductService) GetListProducts(queries md.ProductQueries) ([]md.GetProductRes, error) {
	listProducts, err := service.ProductRepository.GetListProducts(queries)
	if err != nil {
		return []md.GetProductRes{}, nil
	}
	listProductsRes := []md.GetProductRes{}
	for _, product := range listProducts {
		dataAppend := md.GetProductRes{
			ID: product.ID,
			SKU: product.SKU,
			Name: product.Name,
			Description: product.Description,
			Price: product.Price,
			Quantity: product.Quantity,
			Category: md.GetListProductCategoryRes{
				CategoryID: product.ProductCategory.ID,
				CategoryName: product.ProductCategory.Name,
			},
			ImageURL: product.ImageURL,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
			DeletedAt: product.DeletedAt,
		}
		listProductsRes = append(listProductsRes, dataAppend)
	}
	return listProductsRes, nil
}
func (service ProductService) GetProduct(productid uint) (md.GetProductRes, error) {
	product, err := service.ProductRepository.GetProduct(productid)
	if err != nil {
		return md.GetProductRes{}, err
	}
	productRes := md.GetProductRes{
		ID: product.ID,
		SKU: product.SKU,
		Name: product.Name,
		Description: product.Description,
		Price: product.Price,
		Quantity: product.Quantity,
		Category: md.GetListProductCategoryRes{
			CategoryID: product.ProductCategory.ID,
			CategoryName: product.ProductCategory.Name,
		},
		ImageURL: product.ImageURL,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		DeletedAt: product.DeletedAt,
	}
	return productRes, nil
}
func (service ProductService) UpdateProduct(productReq md.UpdateProductReq, productid uint) error {
	//Update Product
	pd := md.Product{
		ID: productid,
		SKU: productReq.SKU,
		Name: productReq.Name,
		Description: productReq.Description,
		Price: productReq.Price,
		Quantity: productReq.Quantity,
		ImageURL: productReq.ImageURL,
	}
	category, err := service.ProductRepository.FindOrCreateCategory(productReq.ProductCategory)
	if err != nil {
		return err
	}
	pd.CategoryID = category.ID
	err = service.ProductRepository.UpdateProduct(pd)
	if err != nil {
		return err
	}
	return nil
}
func (service ProductService) SetDeleteProduct(productid uint) error {
	if err := service.ProductRepository.SetDeleteProduct(productid); err != nil {
		return err
	}
	return nil
}

//Product Category
func (service ProductService) FindListCategory() ([]md.ProductCategory, error) {
	listCategory, err := service.ProductRepository.FindListCategory()
	if err != nil {
		return []md.ProductCategory{}, err
	}
	return listCategory, nil
}