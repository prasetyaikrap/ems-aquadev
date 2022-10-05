package models

import "time"

type (
	Product struct {
		ID uint `json:"id" gorm:"primaryKey;not null;;unique"`
		SKU string `json:"sku" gorm:"type:varchar(100);not null;default:null"`
		Name string `json:"name" gorm:"type:varchar(100);not null;default:null"`
		Description string `json:"description" gorm:"type:text"`
		ImageURL string `json:"image_url" gorm:"type:text;"`
		Price uint `json:"price" gorm:"not null;default:null"`
		CategoryID uint `json:"category_id" gorm:"not null;default:null"`
		Quantity uint `json:"quantity" gorm:"not null; default:null"`
		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at" gorm:"default:null"`
		ProductCategory ProductCategory `gorm:"foreignKey:CategoryID"`
	}

	ProductCategory struct{
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		Name string `json:"name" gorm:"type:varchar(50);not null;default:null;unique"`
		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
	}
)

//Product Request and Response
type (
	CreateProductReq struct {
		SKU string `json:"sku"`
		Name string `json:"name"`
		Description string `json:"description"`
		Price uint `json:"price"`
		Quantity uint `json:"quantity"`
		ProductCategory string `json:"product_category"`
		ImageURL string `json:"image_url"`
	}
	CreateProductRes struct {
		ID uint `json:"product_id"`
		SKU string `json:"product_sku"`
		Name string `json:"product_name"`
	}
	GetListProductRes struct {
		ID uint `json:"id"`
		SKU string `json:"sku"`
		Name string `json:"name"`
		Description string `json:"description"`
		Price uint `json:"price"`
		Quantity uint `json:"quantity"`
		Category GetListProductCategoryRes `json:"category"`
		ImageURL string `json:"image_url"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}
	GetListProductCategoryRes struct {
		CategoryID uint `json:"category_id"`
		CategoryName string `json:"category_name"`
	}
	UpdateProductReq struct {
		SKU string `json:"sku"`
		Name string `json:"name"`
		Description string `json:"description"`
		Price uint `json:"price"`
		Quantity uint `json:"quantity"`
		ProductCategory string `json:"product_category"`
		ImageURL string `json:"image_url"`
	}
)

//Product Process
type (
	ProductQueries struct {
		CategoryID uint
		Price []uint
	}
)