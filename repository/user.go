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

//Admin Repository
func (repo UserRepository) StoreAdmin(admin md.Admin) (md.Admin, error) {
	if err := repo.db.Create(&admin).Error; err != nil {
		return md.Admin{}, err
	}
	return admin, nil
}
func (repo UserRepository) FindAdminByUsername(username string) (md.Admin, error) {
	admin := md.Admin{}
	if err := repo.db.Where("username = ?", username).First(&admin).Error; err != nil {
		return md.Admin{}, err
	}
	return admin, nil
}

//User and Profile Repository
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

//User Address Repository
func (repo UserRepository) StoreUserAddress(userAddress md.UserAddress) (md.UserAddress, error) {
	if err := repo.db.Create(&userAddress).Error; err != nil {
		return md.UserAddress{}, err
	}
	return userAddress, nil
}
func (repo UserRepository) FindListAddress(userid string, status string) ([]md.UserAddress, error) {
	var (
		addresses []md.UserAddress
	)
	if err := repo.db.Where("user_uid = ? AND deleted_at IS NULL", userid).Find(&addresses).Error; err != nil {
		return []md.UserAddress{}, err
	}
	return addresses, nil
}
func (repo UserRepository) FindAddressByID(userid string, id uint, mode string) (md.UserAddress, error) {
	query := ""
	if mode == "ALL" {
		query = "user_uid = ? AND id = ?"
	} else {
		query = "user_uid = ? AND id = ? AND deleted_at IS NULL"
	}
	address := md.UserAddress{}
	if err := repo.db.Where(query,userid,id).First(&address).Error; err != nil {
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

//User Payment
func (repo UserRepository) StoreUserPayment(userData md.UserPayment) (md.UserPayment, error){
	result := repo.db.Create(&userData)
	return userData, result.Error
}
func (repo UserRepository) FindListPayments(userid string) ([]md.UserPayment, error) {
	payments := []md.UserPayment{}
	if err := repo.db.Where("deleted_at IS NULL").Find(&payments).Error; err != nil {
		return []md.UserPayment{}, err
	}
	return payments, nil
}
func (repo UserRepository) FindPayment(userid string, paymentid uint, mode string) (md.UserPayment, error) {
	query := ""
	if mode == "ALL" {
		query = "user_uid = ? AND id = ?"
	} else {
		query = "user_uid = ? AND id = ? AND deleted_at IS NULL"
	}
	payment := md.UserPayment{}
	if err := repo.db.Where(query,userid,paymentid).First(&payment).Error; err != nil {
		return md.UserPayment{}, err
	}
	return payment, nil
}
func (repo UserRepository) UpdatePayment(address md.UserPayment) error {
	if err := repo.db.Omit("deleted_at").Save(&address).Error; err != nil {
		return err
	}
	return nil
}
func (repo UserRepository) SetDeletedPayment(userid string, paymentid uint) error {
	payment := md.UserPayment{}
	result := repo.db.Model(&payment).Where("user_uid = ? AND id = ?", userid, paymentid).Update("deleted_at",time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("No rows affected. The record probably does not exist")
	}
	return nil
}

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
func (repo UserRepository) AddItemToCart(cartItem md.CartItem) (error) {
	result := repo.db.Debug().Model(md.CartItem{}).Where("session_id = ? AND product_id = ?", cartItem.SessionID, cartItem.ProductID).Updates(&cartItem)
	if result.RowsAffected != 0 {
		return nil
	}
	if err := repo.db.Debug().Create(&cartItem).Error; err != nil {
		return err
	}
	return nil
}
func (repo UserRepository) DeleteItemFromCart(sessionId uint, cartItemId []uint) error {
	if err := repo.db.Where("session_id = ? AND id IN (?)",sessionId, cartItemId).Delete(&md.CartItem{}).Error; err != nil {
		return err
	}
	return nil
}
func (repo UserRepository) CreateOrder(order md.Order) (md.Order, error) {
	if err := repo.db.Create(&order).Error; err != nil {
		return md.Order{}, err
	}
	return order, nil
}
func (repo UserRepository) CreateOrderItem(orderItems []md.OrderItem) (error) {
	if err := repo.db.Create(&orderItems).Error; err != nil {
		return err
	}
	return nil
}
func (repo UserRepository) FindListOrders(userid string) ([]md.Order, error) {
	orders := []md.Order{}
	if err := repo.db.Preload(clause.Associations).Where("user_uid = ?", userid).Find(&orders).Error; err != nil {
		return []md.Order{}, err
	}
	return orders, nil
}
func (repo UserRepository) FindOrder(userid string, orderid uint) (md.Order, error) {
	order := md.Order{}
	if err := repo.db.Preload("User").Preload("UserAddress").Preload("PaymentDetails").Where("user_uid = ? AND id = ?",userid,orderid).First(&order).Error; err != nil {
		return md.Order{}, err
	}
	return order, nil
}
func (repo UserRepository) FindOrderItems(orderid uint) ([]md.OrderItem, error) {
	orderItems := []md.OrderItem{}
	if err := repo.db.Preload("Product").Preload(clause.Associations).Where("order_id = ?", orderid).Find(&orderItems).Error; err != nil {
		return []md.OrderItem{}, err
	}
	return orderItems, nil
}
func (repo UserRepository) UpdateReceipt(paymentid uint, receiptURL string) error {
	if err := repo.db.Debug().Where("id = ?", paymentid).Updates(md.PaymentDetails{ReceiptURL: receiptURL}).Error; err != nil {
		return err
	}
	return nil
}