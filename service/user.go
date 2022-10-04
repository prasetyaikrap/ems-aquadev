package service

import (
	md "ems-aquadev/models"
	repo "ems-aquadev/repository"
	"ems-aquadev/utils"
)

type IService interface {
	CreateUserTransaction(userReq md.UserRegRequest) (md.UserRegResponse, error)
	LoginUser(userReq md.UserLoginRequest) (md.UserLoginResponse, error)
	GetUserProfile(uid string) (md.GetUserProfile, error)
}

type UserService struct {
	userRepository repo.IUserRepository
}

func NewUserService(userRepository repo.IUserRepository) *UserService {
	return &UserService{userRepository}
}

func (service UserService) CreateUserTransaction(userReq md.UserRegRequest) (md.UserRegResponse, error) {
	hashPassword, hashErr := utils.HashPassword(userReq.Password)
	if hashErr != nil {
		return md.UserRegResponse{}, hashErr
	}
	userReq.Password = hashPassword
	user, err := service.userRepository.StoreUserTransaction(userReq)
	if err != nil {
		return md.UserRegResponse{}, err
	}
	userRes := md.UserRegResponse{
		UID: user.UID,
		Username: user.Username,
	}

	return userRes, nil
}

func (service UserService) LoginUser(userReq md.UserLoginRequest) (md.UserLoginResponse, error) {
	user, err := service.userRepository.FindUserByUsername(userReq.Username)
	if err != nil || user.UID == ""{
		return md.UserLoginResponse{}, err
	}
	err = utils.CheckPasswordHash(userReq.Password,user.Password)
	if err != nil {
		return md.UserLoginResponse{}, err
	}

	token, _ := utils.GenerateJWTToken(user.UID,user.Username)

	return md.UserLoginResponse{
		AccessToken: token,
	}, nil
}

func (service UserService) GetUserProfile(uid string) (md.GetUserProfile, error) {
	var (
		user md.User
		profile md.UserProfile
		userProfile md.GetUserProfile
		err error
	)
	
	user, err = service.userRepository.FindUserByUID(uid)
	if err != nil {
		return md.GetUserProfile{}, err
	}
	profile, err = service.userRepository.FindProfileByUID(uid)
	if err != nil {
		return md.GetUserProfile{}, err
	}
	userProfile = md.GetUserProfile{
		UID: user.UID,
		Username: user.Username,
		Fullname: profile.Fullname,
		Email: profile.Email,
		Gender: profile.Gender,
		Phone: profile.Phone,
	}

	return userProfile, nil
}

// func (service UserService) NewUserAddress(uid string, userReq md.CreateUserAddressReq) (md.CreateUserAddressRes, error) {
// 	ua := md.UserAddress{
// 		UserUID: uid,
// 		AddressLine: userReq.AddressLine,
// 		City: userReq.City,
// 		Province: userReq.Province,
// 		PostalCode: userReq.PostalCode,
// 		Country: userReq.Country,
// 		RegionID: userReq.RegionID,
// 		IsDefault: userReq.IsDefault,
// 		CreatedAt: userReq.CreatedAt,
// 		UpdatedAt: userReq.UpdatedAt,
// 	}
// 	user, err := service.userRepository.StoreUserAddress(ua)
// 	if err != nil {
// 		return md.CreateUserAddressRes{}, err
// 	}
// 	userRes := md.CreateUserAddressRes{
		
// 	}
// }

