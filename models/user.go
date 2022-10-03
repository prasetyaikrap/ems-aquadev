package models

import "time"

type (
	User struct{
		UID string `json:"uid" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null;unique"`
		Username string `json:"username" gorm:"type:varchar(100);default:null;not null;unique"`
		Email string `json:"email" gorm:"type:varchar(100);default:null;not null;unique"`
		Password string `json:"password" gorm:"type:text;default:null;not null;"`
		AccessToken string `json:"access_token" gorm:"type:text;not null"`
		RefreshToken string `json:"refresh_token" gorm:"type:text;not_null"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	UserProfile struct{
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		UserUID string `json:"user_uid" gorm:"type:uuid;not null"`
		Firstname string `json:"firstname" gorm:"type:varchar(50);not null"`
		Middlename string `json:"middlename" gorm:"type:varchar(50)"`
		Lastname string `json:"lastname" gorm:"type:varchar(50);not null"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User User `gorm:"foreignKey:UserUID"`
	}

	UserAddress struct {
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		UserUID string `json:"user_uid" gorm:"type:uuid;not null;default:null"`
		AddressLine string `json:"address_line" gorm:"type:text;not null;default:null"`
		City string `json:"city" gorm:"type:varchar(100);not null;default:null"`
		Province string `json:"province" gorm:"type:varchar(100);not null;default:null"`
		PostalCode uint `json:"postal_code"`
		Country string `json:"country" gorm:"type:varchar(100);not null;default:null"`
		RegionID string `json:"region_id" gorm:"type:varchar(20)"`
		IsDefault bool `json:"is_default" gorm:"not null"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User User `gorm:"foreignKey:UserUID"`
	}

	UserPayment struct {
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		UserUID string `json:"user_uid" gorm:"type:uuid;not null;default:null"`
		PaymentType string `json:"payment_type" gorm:"type:varchar(100);not null;default:null"`
		Provider string `json:"provider" gorm:"type:varchar(100);not null;default:null"`
		AccountNumber uint `json:"account_number"`
		Exp time.Time `json:"exp"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User User `gorm:"foreignKey:UserUID"`
	}
)