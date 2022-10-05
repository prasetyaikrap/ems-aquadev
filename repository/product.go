package repository

import (
	md "ems-aquadev/models"

	"gorm.io/gorm"
)

type (
	IProductRepository interface{
		StoreUser(user md.User) (md.User, error)
		FindUserByUsername(username string) (md.User, error)
		FindUserProfileByUID(uid string) (md.User, error)
		FindUserProfileID(uid string)(uint, error)
		UpdateProfile(profile md.UserProfile) (md.UserProfile, error)
		StoreUserAddress(userAddress md.UserAddress) (md.UserAddress, error)
		FindListAddress(uid string, status string) ([]md.UserAddress, error)
		FindAddressByID(uid string, id uint) (md.UserAddress, error)
		UpdateAddressByID(address md.UserAddress) (md.UserAddress, error)
		SetDeletedAddress(uid string, id uint) error
	}

	ProductRepository struct {
		db *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}