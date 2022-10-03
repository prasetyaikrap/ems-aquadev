package models

import "time"

type (
	Admin struct{
		UID string `json:"uid" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null;unique"`
		Username string `json:"username" gorm:"type:varchar(100);default:null;not null;unique"`
		Email string `json:"email" gorm:"type:varchar(100);default:null;not null;unique"`
		Password string `json:"password" gorm:"type:text;default:null;not null;"`
		AccessToken string `json:"access_token" gorm:"type:text;not null"`
		RefreshToken string `json:"refresh_token" gorm:"type:text;not_null"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	AdminProfile struct{
		ID uint `json:"id" gorm:"primaryKey;not null;unique"`
		AdminUID string `json:"admin_uid" gorm:"type:uuid;not null"`
		Firstname string `json:"firstname" gorm:"type:varchar(50);not null"`
		Middlename string `json:"middlename" gorm:"type:varchar(50)"`
		Lastname string `json:"lastname" gorm:"type:varchar(50);not null"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Admin Admin `gorm:"foreignKey:AdminUID"`
	}
)
