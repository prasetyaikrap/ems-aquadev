package models

import "time"

//Database Table
type (
	User struct{
		UID string `json:"uid" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null;unique"`
		Username string `json:"username" gorm:"type:varchar(20);default:null;not null;unique"`
		Password string `json:"password" gorm:"type:text;default:null;not null;"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	UserProfile struct{
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		UserUID string `json:"user_uid" gorm:"type:uuid;not null"`
		Fullname string `json:"fullname" gorm:"type:varchar(100);not null;default:null"`
		Gender string `json:"gender" gorm:"type:varchar(20)"`
		Email string `json:"email" gorm:"type:varchar(100);not null;default:null"`
		Phone uint `json:"phone" gorm:"default:0"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User User `gorm:"foreignKey:UserUID"`
	}

	UserAddress struct {
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		UserUID string `json:"user_uid" gorm:"type:uuid;not null;default:null"`
		AddressLine string `json:"address_line" gorm:"type:text;not null;default:null"`
		City string `json:"city" gorm:"type:varchar(50);not null;default:null"`
		Province string `json:"province" gorm:"type:varchar(50);not null;default:null"`
		PostalCode uint `json:"postal_code"`
		Country string `json:"country" gorm:"type:varchar(50);not null;default:null"`
		RegionID string `json:"region_id" gorm:"type:varchar(10);not null;default:null"`
		IsDefault bool `json:"is_default" gorm:"not null;default:false"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User User `gorm:"foreignKey:UserUID"`
	}

	UserPayment struct {
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		UserUID string `json:"user_uid" gorm:"type:uuid;not null;default:null"`
		PaymentType string `json:"payment_type" gorm:"type:varchar(30);not null;default:null"`
		Provider string `json:"provider" gorm:"type:varchar(50);not null;default:null"`
		AccountNumber uint `json:"account_number"`
		Exp time.Time `json:"exp"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User User `gorm:"foreignKey:UserUID"`
	}
)

//User reqeust and Response
type (
	UserRegRequest struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Fullname string `json:"fullname"`
		Password string `json:"password"`
	}

	UserRegResponse struct {
		UID string `json:"uid"`
		Username string `json:"username"`
	}

	UserLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UserLoginResponse struct {
		AccessToken string `json:"access_token"`
	}
	
	GetUserProfile struct {
		UID string `json:"uid"`
		Username string `json:"username"`
		Fullname string `json:"fullname"`
		Gender string `json:"gender"`
		Email string `json:"email"`
		Phone uint `json:"phone"`
	}

	GetUserAddress struct {
		ID uint `json:"id"`
		AddressLine string `json:"address_line"`
		City string `json:"city"`
		Province string `json:"province"`
		PostalCode uint `json:"postal_code"`
		Country string `json:"country"`
		RegionID string `json:"region_id"`
		IsDefault bool `json:"is_default"`
	}

	GetUserPayment struct {
		ID uint `json:"id"`
		UserUID string `json:"user_uid"`
		PaymentType string `json:"payment_type"`
		Provider string `json:"provider"`
		AccountNumber uint `json:"account_number"`
		Exp time.Time `json:"exp"`
	}

	UpdateProfileReq struct {
		UID string `json:"uid"`
		Username string `json:"username"`
		Fullname string `json:"fullname"`
		Gender string `json:"gender"`
		Email string `json:"email"`
		Phone uint `json:"phone"`
	}
)