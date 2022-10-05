package service

import (
	md "ems-aquadev/models"
	repo "ems-aquadev/repository"
	"ems-aquadev/utils"
)

type UserService struct {
	userRepository *repo.UserRepository
}

func NewUserService(userRepository *repo.UserRepository) *UserService {
	return &UserService{userRepository}
}

//User and Profile Service
func (service UserService) CreateUser(userReq md.UserRegRequest) (md.UserRegResponse, error) {
	hashPassword, hashErr := utils.HashPassword(userReq.Password)
	if hashErr != nil {
		return md.UserRegResponse{}, hashErr
	}
	u := md.User{
		Username: userReq.Username,
		Password: hashPassword,
		UserProfile: md.UserProfile{
			Fullname: userReq.Fullname,
			Email: userReq.Email,
		},
	}
	user, err := service.userRepository.StoreUser(u)
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
	user, err := service.userRepository.FindUserProfileByUID(uid)
	if err != nil {
		return md.GetUserProfile{}, err
	}
	userProfile := md.GetUserProfile{
		UID: user.UID,
		Username: user.Username,
		Fullname: user.UserProfile.Fullname,
		Email: user.UserProfile.Email,
		Gender: user.UserProfile.Gender,
		Phone: user.UserProfile.Phone,
	}

	return userProfile, nil
}
func (service UserService) UpdateUserProfile(profileReq md.UpdateProfileReq, uid string) error {
	profileID, err:= service.userRepository.FindUserProfileID(uid)
	if err != nil {
		return err
	}
	updatedProfile := md.UserProfile{
		ID: profileID,
		Fullname: profileReq.Fullname,
		Email: profileReq.Email,
		Gender: profileReq.Gender,
		Phone: profileReq.Phone,
	}
	_, err = service.userRepository.UpdateProfile(updatedProfile)
	if err != nil {
		return err
	}
	return nil
}

//User Address Service
func (service UserService) CreateUserAddress(userReq md.UserAddressReq, uid string) (md.CreateUserAddressRes, error) {
	ua := md.UserAddress{
		UserUID: uid,
		AddressLabel: userReq.AddressLabel,
		AddressLine: userReq.AddressLine,
		City: userReq.City,
		Province: userReq.Province,
		PostalCode: userReq.PostalCode,
		Country: userReq.Country,
		RegionID: userReq.RegionID,
		IsDefault: userReq.IsDefault,
	}
	userAddress, err := service.userRepository.StoreUserAddress(ua)
	if err != nil {
		return md.CreateUserAddressRes{}, err
	}
	userRes := md.CreateUserAddressRes{
		AddressID: userAddress.ID,
		UserUID: userAddress.UserUID,
	}

	return userRes, nil
}
func (service UserService) GetListAddress(userid string, status string) ([]md.UserAddressRes, error) {
	addresses, err := service.userRepository.FindListAddress(userid, status)
	if err != nil {
		return []md.UserAddressRes{}, err
	}
	listAddress := []md.UserAddressRes{}
	for _, address := range addresses {
		dataAppend := md.UserAddressRes{
			ID: address.ID,
			AddressLabel: address.AddressLabel,
			AddressLine: address.AddressLine,
			City: address.City,
			Province: address.Province,
			PostalCode: address.PostalCode,
			Country: address.Country,
			RegionID: address.RegionID,
			IsDefault: address.IsDefault,
		}
		listAddress = append(listAddress, dataAppend)
	}
	return listAddress, nil
}
func (service UserService) GetAddressByID(userid string, addressid uint) (md.UserAddressRes, error) {
	address, err := service.userRepository.FindAddressByID(userid, addressid)
	if err != nil {
		return md.UserAddressRes{}, err
	}
	addressRes := md.UserAddressRes{
		ID: address.ID,
		AddressLabel: address.AddressLabel,
		AddressLine: address.AddressLine,
		City: address.City,
		Province: address.Province,
		PostalCode: address.PostalCode,
		Country: address.Country,
		RegionID: address.RegionID,
		IsDefault: address.IsDefault,
	}
	return addressRes, nil
}
func (service UserService) UpdateAddress(addressReq md.UserAddressReq, uid string, addressID uint) error {
	addr := md.UserAddress{
		ID: addressID,
		UserUID: uid,
		AddressLabel: addressReq.AddressLabel,
		AddressLine: addressReq.AddressLine,  
		City: addressReq.City,
		Province: addressReq.Province,
		PostalCode: addressReq.PostalCode, 
		Country: addressReq.Country,
		RegionID: addressReq.RegionID,
		IsDefault: addressReq.IsDefault,
	}

	_, err := service.userRepository.UpdateAddressByID(addr)
	if err != nil {
		return err
	}
	return nil
}
func (service UserService) SetDeletedAddress(userid string, addressid uint) error {
	if err := service.userRepository.SetDeletedAddress(userid, addressid); err != nil {
		return err
	}
	return nil
}

//Transaction
func (service UserService) FindOrCreateCart(userid string) (interface{}, error) {
	cartSession, err1 := service.userRepository.FindOrCreateCart(userid)
	if err1 != nil {
		return md.CartSession{}, err1
	}
	cartItems, err2 := service.userRepository.FindCartItems(cartSession.ID)
	if err2 != nil {
		return md.CartSession{}, err2
	}
	resultCartItems := []md.CartItemRes{}
	for _, item := range cartItems {
		cartSession.Total = cartSession.Total + (item.Quantity * item.Product.Price)
		dataAppend := md.CartItemRes{
			ID: item.ID,
			SessionID: item.SessionID,
			ProductID: item.ProductID,
			Quantity: item.Quantity,
			UpdatedAt: item.UpdatedAt,
			Product: item.Product, 
		}
		resultCartItems = append(resultCartItems, dataAppend)
	}

	cartResult := md.CartSessionRes{
		ID: cartSession.ID,
		UserUID: cartSession.UserUID,
		Total: cartSession.Total,
		UpdatedAt: cartSession.UpdatedAt,
		CartItems: resultCartItems,
	}
	return cartResult, nil
}
func (service UserService) AddItemToCart()

