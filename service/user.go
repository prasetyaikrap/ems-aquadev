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
//Admin and Profile
func (service UserService) CreateAdmin(adminReq md.AdminRegReq) (md.AdminRegRes, error) {
	hashPassword, hashErr := utils.HashPassword(adminReq.Password)
	if hashErr != nil {
		return md.AdminRegRes{}, hashErr
	}
	adm := md.Admin{
		Username: adminReq.Username,
		Email: adminReq.Email,
		Password: hashPassword,
		Fullname: adminReq.Fullname,
	}
	admin, err := service.userRepository.StoreAdmin(adm)
	if err != nil {
		return md.AdminRegRes{}, err
	}
	adminRes := md.AdminRegRes{
		UID: admin.UID,
		Username: admin.Username,
	}

	return adminRes, nil
}
func (service UserService) LoginAdmin(adminReq md.AdminLoginReq) (md.AdminLoginRes, error) {
	admin, err := service.userRepository.FindAdminByUsername(adminReq.Username)
	if err != nil || admin.UID == ""{
		return md.AdminLoginRes{}, err
	}
	err = utils.CheckPasswordHash(adminReq.Password, admin.Password)
	if err != nil {
		return md.AdminLoginRes{}, err
	}

	token, _ := utils.GenerateJWTToken(admin.UID,admin.Username, "admin")

	return md.AdminLoginRes{
		AccessToken: token,
	}, nil
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

	token, _ := utils.GenerateJWTToken(user.UID,user.Username, "user")

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

//User Payments
func (service UserService) CreateUserPayment(userid string, paymentReq md.UserPaymentReq) (md.CreateUserPaymentRes, error) {
	up := md.UserPayment{
		UserUID: userid,
		PaymentType: paymentReq.PaymentType,
		Provider: paymentReq.Provider,
		AccountNumber: paymentReq.AccountNumber,
		Exp: paymentReq.Exp,
	}
	user, err := service.userRepository.StoreUserPayment(up)
	userPaymentID := md.CreateUserPaymentRes{
		ID: user.ID,
		UserUID: user.UserUID,
	}
	return userPaymentID, err
}
func (service UserService) GetListPayments(userid string) ([]md.GetUserPaymentsRes, error) {
	userPayments, err := service.userRepository.FindListPayments(userid)
	if err != nil {
		return []md.GetUserPaymentsRes{}, err
	}
	listPayments := []md.GetUserPaymentsRes{}
	for _, payment := range userPayments {
		dataAppend := md.GetUserPaymentsRes{
			ID: payment.ID,
			UserUID: payment.UserUID,
			PaymentType: payment.PaymentType,
			Provider: payment.Provider,
		}
		listPayments = append(listPayments, dataAppend)
	}
	return listPayments, nil
}
func (service UserService) GetPayment(userid string, paymentid uint) (md.GetUserPaymentRes, error) {
	payment, err := service.userRepository.FindPayment(userid, paymentid)
	if err != nil {
		return md.GetUserPaymentRes{}, err
	}
	paymentRes := md.GetUserPaymentRes{
		ID: payment.ID,
		UserUID: payment.UserUID,
		PaymentType: payment.PaymentType,
		Provider: payment.Provider,
		AccountNumber: payment.AccountNumber,
		Exp: payment.Exp,
	}
	return paymentRes, nil
}
func (service UserService) UpdatePayment(paymentReq md.UserPaymentReq, userid string, paymentid uint) error {
	userPayment := md.UserPayment{
		ID: paymentid,
		UserUID: userid,
		PaymentType: paymentReq.PaymentType,
		Provider: paymentReq.Provider,
		AccountNumber: paymentReq.AccountNumber,
		Exp: paymentReq.Exp,
	}

	if err := service.userRepository.UpdatePayment(userPayment); err != nil {
		return err
	}
	return nil
}
func (service UserService) SetDeletedPayment(userid string, paymentid uint) error {
	if err := service.userRepository.SetDeletedPayment(userid, paymentid); err != nil {
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
func (service UserService) AddItemToCart(cartItemReq md.AddItemToCartReq) (error) {
	cartItem := md.CartItem{
		SessionID: cartItemReq.SessionID,
		ProductID: cartItemReq.ProductID,
		Quantity: cartItemReq.Quantitty,
	}
	err := service.userRepository.AddItemToCart(cartItem)
	if err != nil {
		return err
	}
	return nil
}
func (service UserService) GetItemsCart(sessionId uint) ([]md.CartItemRes, error) {
	cartItems, err := service.userRepository.FindCartItems(sessionId)
	if err != nil {
		return []md.CartItemRes{}, err
	}
	resultCartItems := []md.CartItemRes{}
	for _, item := range cartItems {
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
	return resultCartItems, nil
}
func (service UserService) DeleteItemFromCart(sessionId, cartItemId uint) error {
	if err := service.userRepository.DeleteItemFromCart(sessionId, []uint{cartItemId}); err != nil {
		return err
	}
	return nil
}
func (service UserService) CreateOrder(userid string, sessionid uint, orderReq md.CreateOrderReq) error {
	od := md.Order{
		UserUID: userid,
		AddressID: orderReq.AddressID,
		Status: "PAYMENT",
		PaymentDetails: md.PaymentDetails{
			UserPaymentID: orderReq.UserPaymentID,
		},
	}
	//Find Cart Items To Add
	cartItems, err := service.userRepository.FindCartItems(sessionid)
	if err != nil {
		return err
	}
	//Calculate Total Payment and Ammount and Popule itemCart ID
	var itemCartId []uint
	for _, item := range cartItems {
		itemCartId = append(itemCartId, item.ID)
		od.Total = od.Total + (item.Quantity * item.Product.Price)
	}
	od.PaymentDetails.Ammount = od.Total
	//Create Order
	order, err := service.userRepository.CreateOrder(od)
	if err != nil {
		return err
	}
	//Conver CartItems to OrderItems and populate OrderItems ID for Deletion
	var orderItems []md.OrderItem
	for _, item := range cartItems {
		appendData := md.OrderItem {
			OrderID: order.ID,
			ProductID: item.ProductID,
			Quantity: item.Quantity,
		}
		orderItems = append(orderItems, appendData)
	}
	//Create Order Items
	if err := service.userRepository.CreateOrderItem(orderItems); err != nil {
		return err
	}
	//Delete Cart Items from Cart Session
	if err := service.userRepository.DeleteItemFromCart(sessionid, itemCartId); err != nil {
		return err
	}
	return nil
}
func (service UserService) GetListOrders(userid string) ([]md.GetOrdersRes, error) {
	orders, err := service.userRepository.FindListOrders(userid)
	if err != nil {
		return []md.GetOrdersRes{}, err
	}
	ordersRes := []md.GetOrdersRes{}
	for _, order := range orders {
		dataAppend := md.GetOrdersRes{
			ID: order.ID,
			UserUID: order.UserUID,
			AddressID: order.AddressID,
			PaymentID: order.PaymentID,
			Total: order.Total,
			Status: order.Status,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		}
		ordersRes = append(ordersRes, dataAppend)
	}
	return ordersRes, nil
}
func (service UserService) GetOrder(userid string, orderid uint) (md.GetOrderRes, error) {
	order, err := service.userRepository.FindOrder(userid, orderid)
	if err != nil {
		return md.GetOrderRes{}, err
	}
	orderItems, err := service.userRepository.FindOrderItems(order.ID)
	if err != nil {
		return md.GetOrderRes{}, err
	}
	orderItemsRes := []md.GetOrderItemRes{}
	for _, item := range orderItems {
		dataAppend := md.GetOrderItemRes{
			ID: item.ID,
			OrderID: item.OrderID,
			ProductID: item.ProductID,
			Quantity: item.Quantity,
			UpdatedAt: item.UpdatedAt,
			Product: item.Product,
		}
		orderItemsRes = append(orderItemsRes, dataAppend)
	}
	orderRes := md.GetOrderRes{
			ID: order.ID,
			UserUID: order.UserUID,
			AddressID: order.AddressID,
			PaymentID: order.PaymentID,
			Total: order.Total,
			Status: order.Status,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			User: md.GetOrderUserRes{
				Fullname: order.User.UserProfile.Fullname,
				Email: order.User.UserProfile.Email,
				Phone: order.User.UserProfile.Phone,
			},
			UserAddress: md.GetOrderUserAddressRes{
				ID: order.UserAddress.ID,
				AddressLabel: order.UserAddress.AddressLabel,
				AddressLine: order.UserAddress.AddressLine,
				City: order.UserAddress.City,
				Province: order.UserAddress.Province,
				PostalCode: order.UserAddress.PostalCode,
				Country: order.UserAddress.Country,
				RegionID: order.UserAddress.RegionID,
			},
			PaymentDetails: md.GetOrderUserPaymentRes{
				ID: order.PaymentDetails.ID,
				PaymentType: order.PaymentDetails.UserPayment.PaymentType,
				Provider: order.PaymentDetails.UserPayment.Provider,
				AccountNumber: order.PaymentDetails.UserPayment.AccountNumber,
				Exp: order.PaymentDetails.UserPayment.Exp,
			},
			OrderItems: orderItemsRes,
		}

		return orderRes, nil
}
func (service UserService) UploadReceipt(paymentid uint, paymentURL string) error {
	if err := service.userRepository.UpdateReceipt(paymentid, paymentURL); err != nil {
		return err
	}
	return nil
}