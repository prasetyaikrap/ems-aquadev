package repository

import (
	md "ems-aquadev/models"

	"gorm.io/gorm"
)

type (
	IUserRepository interface{
		StoreUserTransaction(uReg md.UserRegRequest) (md.User, error)
		FindUserByUsername(username string) (md.User, error)
		FindUserByUID(uid string) (md.User, error)
		FindProfileByUID(uid string) (md.UserProfile, error)
		ListAddressByUID(uid string) ([]md.UserAddress, error)
		ListPaymentByUID(uid string) ([]md.UserPayment, error)
		StoreUserAddress(userAddress md.UserAddress) (md.UserAddress, error)
		StoreUserPayment(userPayment md.UserPayment) (md.UserPayment, error)
	}

	UserRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo UserRepository) StoreUserTransaction(uReg md.UserRegRequest) (md.User, error) {
	userReg := md.User{
		Username: uReg.Username,
		Password: uReg.Password,
	}
	userProfile := md.UserProfile{
		Email: uReg.Email,
		Fullname: uReg.Fullname,
	}
	userCartSession := md.CartSession{}
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		//Create New User to table user
		if err := tx.Debug().Create(&userReg).Error; err != nil {
			return err
		}
		userProfile.UserUID = userReg.UID
		userCartSession.UserUID = userReg.UID
		//Create user profile
		if err := tx.Debug().Create(&userProfile).Error; err != nil {
			return err
		}
		// Create User Cart Session
		if err := tx.Debug().Create(&userCartSession).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return md.User{}, err
	}
	return userReg, nil
}

func (repo UserRepository) FindUserByUsername(username string) (md.User, error) {
	user := md.User{}
	if err := repo.db.Debug().Where("username = ?", username).First(&user).Error; err != nil {
		return md.User{}, err
	}
	return user, nil
}

func (repo UserRepository) FindUserByUID(uid string) (md.User, error) {
	user := md.User{}
	if err := repo.db.Debug().Where("uid = ?", uid).First(&user).Error; err != nil {
		return md.User{}, err
	}
	return user, nil
}

func (repo UserRepository) FindProfileByUID(uid string) (md.UserProfile, error) {
	profile := md.UserProfile{}
	if err := repo.db.Debug().Where("user_uid = ?", uid).First(&profile).Error; err != nil {
		return md.UserProfile{}, err
	}
	return profile, nil
}

func (repo UserRepository) ListAddressByUID(uid string) ([]md.UserAddress, error) {
	addresses := []md.UserAddress{}
	if err := repo.db.Debug().Where("user_uid = ?", uid).Find(&addresses).Error; err != nil {
		return []md.UserAddress{}, err
	}
	return addresses, nil
}

func (repo UserRepository) ListPaymentByUID(uid string) ([]md.UserPayment, error) {
	payments := []md.UserPayment{}
	if err := repo.db.Debug().Where("user_uid = ?", uid).Find(&payments).Error; err != nil {
		return []md.UserPayment{}, err
	}
	return payments, nil
}

func (repo UserRepository) StoreUserAddress(userAddress md.UserAddress) (md.UserAddress, error) {
	if err := repo.db.Debug().Create(&userAddress).Error; err != nil {
		return md.UserAddress{}, err
	}
	return userAddress, nil
}

func (repo UserRepository) StoreUserPayment(userPayment md.UserPayment) (md.UserPayment, error) {
	if err := repo.db.Debug().Create(&userPayment).Error; err != nil {
		return md.UserPayment{}, err
	}
	return userPayment, nil
}