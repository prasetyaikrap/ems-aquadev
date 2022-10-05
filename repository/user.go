package repository

import (
	md "ems-aquadev/models"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
		db *gorm.DB
	}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}
//User and Profile Repo
func (repo UserRepository) StoreUser(user md.User) (md.User, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		return md.User{}, err
	}
	return user, nil
}
func (repo UserRepository) FindUserByUsername(username string) (md.User, error) {
	user := md.User{}
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return md.User{}, err
	}
	return user, nil
}
func (repo UserRepository) FindUserProfileByUID(uid string) (md.User, error) {
	user := md.User{}
	if err := repo.db.Joins("UserProfile").Where("uid = ?", uid).First(&user).Error; err != nil {
		return md.User{}, err
	}
	return user, nil
}
func (repo UserRepository) FindUserProfileID(uid string)(uint, error) {
	user := md.User{}
	if err := repo.db.Select("profile_id").Where("uid = ?", uid).First(&user).Error; err != nil {
		return 0, err
	}
	return user.ProfileID, nil
}
func (repo UserRepository) UpdateProfile(profile md.UserProfile) (md.UserProfile, error) {
	if err := repo.db.Save(&profile).Error; err != nil {
		return md.UserProfile{}, err
	}
	return profile, nil
}

//User Address Repo
func (repo UserRepository) StoreUserAddress(userAddress md.UserAddress) (md.UserAddress, error) {
	if err := repo.db.Create(&userAddress).Error; err != nil {
		return md.UserAddress{}, err
	}
	return userAddress, nil
}
func (repo UserRepository) FindListAddress(userid string, status string) ([]md.UserAddress, error) {
	var (
		addresses []md.UserAddress
		query string
		queryValue []string
	)
	
	switch {
	case status == "":
		query = "user_uid = ?"
		queryValue = []string{userid}
	case status == "active":
		query = "user_uid = ? AND deleted_at IS NULL"
		queryValue = []string{userid}
	case status == "inactive":
		query = "user_uid = ? AND deleted_at IS NOT NULL"
		queryValue = []string{userid}
	}
	if err := repo.db.Where(query, queryValue).Find(&addresses).Error; err != nil {
		return []md.UserAddress{}, err
	}
	return addresses, nil
}
func (repo UserRepository) FindAddressByID(userid string, id uint) (md.UserAddress, error) {
	address := md.UserAddress{}
	if err := repo.db.Debug().Where("user_uid = ? AND id = ?",userid,id).First(&address).Error; err != nil {
		return md.UserAddress{}, err
	}
	return address, nil
}
func (repo UserRepository) UpdateAddressByID(address md.UserAddress) (md.UserAddress, error) {
	if err := repo.db.Omit("deleted_at").Save(&address).Error; err != nil {
		return md.UserAddress{}, err
	}
	return address, nil
}
func (repo UserRepository) SetDeletedAddress(userid string, addressid uint) error {
	address := md.UserAddress{}
	result := repo.db.Model(&address).Where("user_uid = ? AND id = ?", userid, addressid).Update("deleted_at",time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("No rows affected. The record probably does not exist")
	}
	return nil
}

//User Payment Repo

//User Transaction
func (repo UserRepository) FindOrCreateCart(userid string) (md.CartSession, error) {
	cartSession := md.CartSession{}
	if err := repo.db.FirstOrCreate(&cartSession,md.CartSession{UserUID: userid}).Error; err != nil {
		return md.CartSession{}, err
	}
	return cartSession, nil
}
func (repo UserRepository) FindCartItems(sessionid uint) ([]md.CartItem, error) {
	cartItems := []md.CartItem{}
	if err := repo.db.Preload("Product").Preload(clause.Associations).Where("session_id = ?", sessionid).Find(&cartItems).Error; err != nil {
		return []md.CartItem{}, err
	}
	return cartItems, nil
}
func (repo UserRepository) AddItemToCart(cartItem md.CartItem) error {
	result := repo.db.Model(md.CartItem{}).Where("session_id = ? AND product_id = ?", cartItem.CartSession,cartItem.ProductID).Updates(cartItem)
	if result.RowsAffected != 0 {
		return nil
	}
	if err := repo.db.Create(&cartItem).Error; err != nil {
		return err
	}
	return nil
}
func (repo UserRepository) DeleteItemFromCart(cartItemId uint) error {
	if err := repo.db.Delete(&md.CartItem{}, cartItemId).Error; err != nil {
		return err
	}
	return nil
}
func (repo UserRepository) CreateOrder()