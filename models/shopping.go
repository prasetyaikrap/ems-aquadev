package models

import "time"

type (
	CartSession struct {
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		UserUID string `json:"user_uid" gorm:"type:uuid;not null;default:null"`
		Total uint `json:"total" gorm:"not null;default:0"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User User `gorm:"foreignKey:UserUID"`
	}

	// CartItem struct {
	// 	ID uint `json:"id" gorm:"primaryKey;not null;unique"`
	// 	SessionID uint `json:"session_id" gorm:"not_null;default:null"`
	// 	ProductID uint `json:"product_id" gorm:"not null;default:null"`
	// 	Quantity uint `json:"quantity"`
	// 	CreatedAt time.Time `json:"created_at"`
	// 	UpdatedAt time.Time `json:"updated_at"`
	// 	CartSession CartSession `gorm:"foreignKey:SessionID"`
	// 	Product Product `gorm:"foreignKey:ProductID"`
	// }

	// Order struct {
	// 	ID uint `json:"id" gorm:"primaryKey;not null;unique"`
	// 	UserUID string `json:"user_uid" gorm:"type:uuid;not null;default:null"`
	// 	AddressID uint `json:"address_id" gorm:"not null;default:null"`
	// 	PaymentID uint `json:"payment_id" gorm:"not null;default:null"`
	// 	Total uint `json:"total" gorm:"not null;default:null"`
	// 	CreatedAt time.Time `json:"created_at"`
	// 	UpdatedAt time.Time `json:"updated_at"`
	// 	User User `gorm:"foreignKey:UserUID"`
	// 	UserAddress UserAddress `gorm:"foreignKey:AddressID"`
	// 	PaymentDetails PaymentDetails `gorm:"foreignKey:PaymentID"`
	// }

	// OrderItem struct {
	// 	ID uint `json:"id" gorm:"primaryKey;not null;unique"`
	// 	OrderID uint `json:"order_id" gorm:"not_null;default:null"`
	// 	ProductID uint `json:"product_id" gorm:"not null;default:null"`
	// 	Quantity uint `json:"quantity"`
	// 	CreatedAt time.Time `json:"created_at"`
	// 	UpdatedAt time.Time `json:"updated_at"`
	// 	Order Order `gorm:"foreignKey:OrderID"`
	// 	Product Product `gorm:"foreignKey:ProductID"`
	// }

	// PaymentDetails struct {
	// 	ID uint `json:"id" gorm:"primaryKey;not null;unique"`
	// 	UserPaymentID uint `json:"user_payment_id"`
	// 	Ammount uint `json:"ammount" gorm:"not null;default:null"`
	// 	ReceiptURL string `json:"receipt_url" gorm:"type:text"`
	// 	CreatedAt time.Time `json:"created_at"`
	// 	UpdatedAt time.Time `json:"updated_at"`
	// 	UserPayment UserPayment `gorm:"foreignKey:UserPaymentID"`
	// }
)