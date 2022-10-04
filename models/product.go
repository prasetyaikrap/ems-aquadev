package models

type (
	// Product struct {
	// 	ID uint `json:"id" gorm:"primaryKey;not null;;unique"`
	// 	SKU string `json:"sku" gorm:"type:varchar(100);not null;default:null"`
	// 	Name string `json:"name" gorm:"type:varchar(100);not null;default:null"`
	// 	Description string `json:"description" gorm:"type:text"`
	// 	Price uint `json:"price" gorm:"not null;default:null"`
	// 	IsPublished bool `json:"is_published" gorm:"not null;default:null"`
	// 	CategoryID uint `json:"category_id"`
	// 	InventoryID uint `json:"inventory_id"`
	// 	CreatedAt time.Time `json:"created_at"`
	// 	UpdatedAt time.Time `json:"updated_at"`
	// 	ProductCategory ProductCategory `gorm:"foreignKey:CategoryID"`
	// 	ProductInventory ProductInventory `gorm:"foreignKey:InventoryID"`
	// }

	// ProductCategory struct{
	// 	ID uint `json:"id" gorm:"primaryKey;not null;unique"`
	// 	Name string `json:"name" gorm:"type:varchar(50);not null"`
	// 	Description string `json:"description" gorm:"type:text"`
	// 	CreatedAt time.Time `json:"created_at"`
	// 	UpdatedAt time.Time `json:"updated_at"`
	// }

	// ProductInventory struct{
	// 	ID uint `json:"id" gorm:"primaryKey;not null;unique"`
	// 	Quantity uint `json:"quantity" gorm:"not null; default:null"`
	// 	CreatedAt time.Time `json:"created_at"`
	// 	UpdatedAt time.Time `json:"updated_at"`
	// }

	// ProductImage struct {
	// 	ID uint `json:"id" gorm:"primaryKey;not null;unique"`
	// 	ProductID uint `json:"product_id"`
	// 	URL string
	// 	CreatedAt time.Time `json:"created_at"`
	// 	UpdatedAt time.Time `json:"updated_at"`
	// 	Product Product `gorm:"foreignKey:ProductID"`
	// }
)