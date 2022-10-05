package repository

import (
	md "ems-aquadev/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type (
	IProductRepository interface{
		StoreProduct(product md.Product) (md.Product, error)
		FindListProducts(queries md.ProductQueries)([]md.Product, error)
	}

	ProductRepository struct {
		db *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

//Main Product
func (repo ProductRepository) StoreProduct(product md.Product) (md.Product, error) {
	if err := repo.db.Create(&product).Error; err != nil {
		return md.Product{}, err
	}
	return product,nil
}
func (repo ProductRepository) FindListProducts(queries md.ProductQueries)([]md.Product, error) {
	var (
		listProducts []md.Product
		query string
		queryValue []interface{}
	)
	switch {
	case queries.CategoryID != 0 && queries.Price[0] != 0 && queries.Price[1] != 0:
		query = "category_id = ? AND price BETWEEN ? AND ?"
		queryValue = []interface{}{queries.CategoryID, queries.Price[0], queries.Price[1]}
	case queries.CategoryID != 0 && queries.Price[0] != 0:
		query = "category_id = ? AND price >= ?"
		queryValue = []interface{}{queries.CategoryID, queries.Price[0]}
	case queries.CategoryID != 0 && queries.Price[1] != 0:
		query = "category_id = ? AND price <= ?"
		queryValue = []interface{}{queries.CategoryID, queries.Price[1]}
	case queries.CategoryID != 0:
		query = "category_id = ?"
		queryValue = []interface{}{queries.CategoryID}
	case queries.Price[0] != 0 && queries.Price[1] != 0:
		query = "price BETWEEN ? AND ?"
		queryValue = []interface{}{queries.Price[0], queries.Price[1]}
	case queries.Price[0] != 0:
		query = "price >= ?"
		queryValue = []interface{}{queries.Price[0]}
	case queries.Price[1] != 0:
		query = "price <= ?"
		queryValue = []interface{}{queries.Price[1]} 
	}
	if query == "" {
		if err := repo.db.Joins("ProductCategory").Find(&listProducts).Error; err != nil {
			return []md.Product{}, err
		}
	} else {
		if err := repo.db.Where(query,queryValue).Joins("ProductCategory").Find(&listProducts).Error; err != nil {
			return []md.Product{}, err
		}
	}
	return listProducts, nil
}
func (repo ProductRepository) UpdateProduct(product md.Product)(error) {
	if err := repo.db.Omit("deleted_at").Save(&product).Error; err != nil {
		return err
	}
	return nil
}
func (repo ProductRepository) SetDeleteProduct(productid uint) error {
	product := md.Product{}
	result := repo.db.Model(&product).Where("id = ?", productid).Update("deleted_at",time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("No rows affected. The record probably does not exist")
	}
	return nil
}

//Product Category
func (repo ProductRepository) FindOrCreateCategory(name string) (md.ProductCategory, error) {
	category := md.ProductCategory{}
	if err := repo.db.Where(md.ProductCategory{Name: name}).FirstOrCreate(&category).Error; err != nil {
		return md.ProductCategory{}, err
	}
	return category, nil
}
func (repo ProductRepository) FindListCategory()([]md.ProductCategory, error) {
	listCategory := []md.ProductCategory{}
	if err := repo.db.Find(&listCategory).Error; err != nil {
		return []md.ProductCategory{}, err
	}
	return listCategory, nil
}