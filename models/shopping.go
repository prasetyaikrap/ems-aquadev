package models

import "time"

type (
	CartSession struct {
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		UserUID string `json:"user_uid" gorm:"type:uuid;not null;default:null"`
		Total uint `json:"total" gorm:"not null;default:0"`
		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt time.Time `json:"updated_at"`
		User User `gorm:"foreignKey:UserUID"`
	}

	CartItem struct {
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		SessionID uint `json:"session_id" gorm:"not_null;default:null"`
		ProductID uint `json:"product_id" gorm:"not null;default:null;unique"`
		Quantity uint `json:"quantity"`
		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt time.Time `json:"updated_at"`
		CartSession CartSession `gorm:"foreignKey:SessionID"`
		Product Product `gorm:"foreignKey:ProductID"`
	}

	Order struct {
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		UserUID string `json:"user_uid" gorm:"type:uuid;not null;default:null"`
		AddressID uint `json:"address_id" gorm:"not null;default:null"`
		PaymentID uint `json:"payment_id" gorm:"not null;default:null"`
		Total uint `json:"total" gorm:"not null;default:null"`
		Status string `json:"status" gorm:"not null;default:null"`
		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt time.Time `json:"updated_at"`
		User User `gorm:"foreignKey:UserUID"`
		UserAddress UserAddress `gorm:"foreignKey:AddressID"`
		PaymentDetails PaymentDetails `gorm:"foreignKey:PaymentID"`
	}

	OrderItem struct {
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		OrderID uint `json:"order_id" gorm:"not_null;default:null"`
		ProductID uint `json:"product_id" gorm:"not null;default:null"`
		Quantity uint `json:"quantity"`
		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt time.Time `json:"updated_at"`
		Order Order `gorm:"foreignKey:OrderID"`
		Product Product `gorm:"foreignKey:ProductID"`
	}

	PaymentDetails struct {
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		UserPaymentID uint `json:"user_payment_id"`
		Ammount uint `json:"ammount" gorm:"not null;default:null"`
		ReceiptURL string `json:"receipt_url" gorm:"type:text"`
		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt time.Time `json:"updated_at"`
		UserPayment UserPayment `gorm:"foreignKey:UserPaymentID"`
	}
)

type (
	CartSessionRes struct {
		ID uint `json:"id"`
		UserUID string `json:"user_uid"`
		Total uint `json:"total"`
		UpdatedAt time.Time `json:"updated_at"`
		CartItems []CartItemRes `json:"cart_items"`
	}
	CartItemRes struct {
		ID uint `json:"id"`
		SessionID uint `json:"session_id"`
		ProductID uint `json:"product_id"`
		Quantity uint `json:"quantity"`
		UpdatedAt time.Time `json:"updated_at"`
		Product Product `json:"products"` 
	}
	AddItemToCartReq struct {
		SessionID uint `json:"session_id"`
		ProductID uint `json:"product_id"`
		Quantitty uint `json:"quantity"`
	}
	CreateOrderReq struct {
		AddressID uint `json:"address_id"`
		UserPaymentID uint `json:"user_payment_id"`
	}
	CreateOrderItemsReq struct {
		ID uint `json:"id"`
		ProductID uint `json:"product_id"`
		Quantitty uint `json:"quantity"`
	}
	GetOrdersRes struct {
		ID uint `json:"id"`
		UserUID string `json:"user_uid"`
		AddressID uint `json:"address_id"`
		PaymentID uint `json:"payment_id"`
		Total uint `json:"total"`
		Status string `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	GetOrderRes struct {
		ID uint `json:"id"`
		UserUID string `json:"user_uid"`
		AddressID uint `json:"address_id"`
		PaymentID uint `json:"payment_id"`
		Total uint `json:"total"`
		Status string `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User GetOrderUserRes `json:"user_details"`
		UserAddress GetOrderUserAddressRes `json:"shipment_address"`
		PaymentDetails GetOrderUserPaymentRes `json:"payment_details"`
		OrderItems []GetOrderItemRes
	}
	GetOrderItemRes struct {
		ID uint `json:"id"`
		OrderID uint `json:"order_id"`
		ProductID uint `json:"product_id"`
		Quantity uint `json:"quantity"`
		UpdatedAt time.Time `json:"updated_at"`
		Product Product `json:"products"` 
	}
	GetOrderUserRes struct {
		Fullname string `json:"fullname"`
		Email string `json:"email"`
		Phone uint `json:"phone"`
	}
	GetOrderUserAddressRes struct {
		ID uint `json:"id"`
		AddressLabel string `json:"address_label"`
		AddressLine string `json:"address_line"`
		City string `json:"city"`
		Province string `json:"province"`
		PostalCode uint `json:"postal_code"`
		Country string `json:"country"`
		RegionID string `json:"region_id"`
	}
	GetOrderUserPaymentRes struct {
		ID uint `json:"id"`
		PaymentType string `json:"payment_type"`
		Provider string `json:"provider"`
		AccountNumber uint `json:"account_number"`
		Exp time.Time `json:"exp"`
	}
	ReceiptURLReq struct {
		PaymentURL string `json:"payment_url"`
	}
)