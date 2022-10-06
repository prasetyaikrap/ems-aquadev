package models

import "time"

type (
	Admin struct{
		UID string `json:"uid" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null;unique"`
		Username string `json:"username" gorm:"type:varchar(100);default:null;not null;unique"`
		Email string `json:"email" gorm:"type:varchar(100);default:null;not null;unique"`
		Password string `json:"password" gorm:"type:text;default:null;not null;"`
		Fullname string `json:"fullname" gorm:"type:varchar(100);default:null;not null"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	AdminRegReq struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Fullname string `json:"fullname"`
		Password string `json:"password"`
	}
	AdminRegRes struct {
		UID string `json:"uid"`
		Username string `json:"username"`
	}
	AdminLoginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	AdminLoginRes struct {
		AccessToken string `json:"access_token"`
	}
)
