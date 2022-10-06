package handler

import (
	md "ems-aquadev/models"
	svc "ems-aquadev/service"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *svc.UserService
}

func NewUserHandler(userService *svc.UserService) *UserHandler {
	return &UserHandler{userService}
}

//Admin handler
func (handler UserHandler) RegisterAdmin(c echo.Context) error {
	req := md.AdminRegReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}

	createdAdmin, err := handler.userService.CreateAdmin(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Sever Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, md.HttpResponse{
		Code: http.StatusCreated,
		Message: "Admin Registered Successfully",
		Data: createdAdmin,
	})
}
func (handler UserHandler) LoginAdmin(c echo.Context) error {
	req := md.AdminLoginReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Register Admin Failed",
			Data: err.Error(),
		})
	}
	adminAuth, err := handler.userService.LoginAdmin(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	if adminAuth.AccessToken == "" {
		return c.JSON(http.StatusUnauthorized, md.HttpResponse{
			Code: http.StatusUnauthorized,
			Message: "Access Denied. Unauthorized",
			Data: "Access token not found",
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
		Code: http.StatusOK,
		Message: "Login Success",
		Data: adminAuth,
	})
}

//User and Profile handler
func (handler UserHandler) CreateUser(c echo.Context) error {
	req := md.UserRegRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}

	createdUser, err := handler.userService.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Sever Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, md.HttpResponse{
		Code: http.StatusCreated,
		Message: "User registered successfully",
		Data: createdUser,
	})
}
func (handler UserHandler) LoginUser(c echo.Context) error {
	req := md.UserLoginRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Register user failed",
			Data: err.Error(),
		})
	}
	userAuth, err := handler.userService.LoginUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Sever Error",
			Data: err.Error(),
		})
	}
	if userAuth.AccessToken == "" {
		return c.JSON(http.StatusUnauthorized, md.HttpResponse{
			Code: http.StatusUnauthorized,
			Message: "Access Denied. Unauthorized",
			Data: "Access token not found",
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
		Code: http.StatusOK,
		Message: "Login Success",
		Data: userAuth,
	})
}
func (handler UserHandler) GetUserProfile(c echo.Context) error {
	userid := c.Param("userid")
	userProfile, err := handler.userService.GetUserProfile(userid)

	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Sever Error",
			Data: err.Error(),
		})
	case userProfile.UID == "":
		return c.JSON(http.StatusNotFound, md.HTTPResponseWithoutData{
			Code: http.StatusNotFound,
			Message: "User Profile Not Found",
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
		Code: http.StatusOK,
		Message: "User Profile Found",
		Data: userProfile,
	})
}
func (handler UserHandler) UpdateUserProfile(c echo.Context) error {
	userid := c.Param("userid")
	req := md.UpdateProfileReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}
	if err := handler.userService.UpdateUserProfile(req,userid); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
		Code: http.StatusOK,
		Message: "Profile Updated Successfully",
	})
}

//User Address handler
func (handler UserHandler) CreateUserAddress(c echo.Context) error {
	userid := c.Param("userid")
	req := md.UserAddressReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}
	userAddress, err := handler.userService.CreateUserAddress(req,userid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Sever Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, md.HttpResponse{
		Code: http.StatusCreated,
		Message: "Address Added Successfully",
		Data: userAddress,
	})
}
func (handler UserHandler) GetListAddress(c echo.Context) error {
	status := c.QueryParam("status")
	userid := c.Param("userid")
	listAddress, err := handler.userService.GetListAddress(userid, status)
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	case len(listAddress) <= 0:
		return c.JSON(http.StatusOK, md.HttpResponse{
			Code: http.StatusOK,
			Message: "List Address is Empty",
			Data: listAddress,
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
			Code: http.StatusOK,
			Message: "List of Address Found",
			Data: listAddress,
		})
}
func (handler UserHandler) GetAddress(c echo.Context) error {
	userid := c.Param("userid")
	addressid, _ := strconv.Atoi(c.Param("addressid"))
	address, err := handler.userService.GetAddressByID(userid, uint(addressid))
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Sever Error",
			Data: err.Error(),
		})
	case address.ID == 0:
		return c.JSON(http.StatusNotFound, md.HTTPResponseWithoutData{
			Code: http.StatusNotFound,
			Message: "Address Not Found",
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
			Code: http.StatusOK,
			Message: "Address Found",
			Data: address,
		})
}
func (handler UserHandler) UpdateAddress(c echo.Context) error {
	userid := c.Param("userid")
	addressid, _ := strconv.Atoi(c.Param("addressid"))
	req := md.UserAddressReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}

	if err := handler.userService.UpdateAddress(req, userid, uint(addressid)); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
		Code: http.StatusOK,
		Message: "Address Updated Successfully",
	})
}
func (handler UserHandler) SetDeletedAddress(c echo.Context) error {
	userid := c.Param("userid")
	addressid,_ := strconv.Atoi(c.Param("addressid"))
	if err := handler.userService.SetDeletedAddress(userid, uint(addressid)); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Delete Address Failed",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
			Code: http.StatusOK,
			Message: "Address Set To Deleted",
		})
}

//User Payment
func (handler UserHandler) CreateUserPayment(c echo.Context) error {
	userid := c.Param("userid")
	req := md.UserPaymentReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}
	userAddress, err := handler.userService.CreateUserPayment(userid,req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Sever Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, md.HttpResponse{
		Code: http.StatusCreated,
		Message: "Payment Added Successfully",
		Data: userAddress,
	})
}
func (handler UserHandler) GetListPayments(c echo.Context) error {
	userid := c.Param("userid")
	listPayments, err := handler.userService.GetListPayments(userid)
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	case len(listPayments) <= 0:
		return c.JSON(http.StatusOK, md.HttpResponse{
			Code: http.StatusOK,
			Message: "List Payments is Empty",
			Data: listPayments,
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
			Code: http.StatusOK,
			Message: "List of Payment Found",
			Data: listPayments,
		})
}
func (handler UserHandler) GetPayment(c echo.Context) error {
	userid := c.Param("userid")
	paymentid, _ := strconv.Atoi(c.Param("paymentid"))
	payment, err := handler.userService.GetPayment(userid, uint(paymentid))
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	case payment.ID == 0:
		return c.JSON(http.StatusNotFound, md.HTTPResponseWithoutData{
			Code: http.StatusNotFound,
			Message: "Address Not Found",
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
			Code: http.StatusOK,
			Message: "Address Found",
			Data: payment,
		})
}
func (handler UserHandler) UpdatePayment(c echo.Context) error {
	userid := c.Param("userid")
	paymentid, _ := strconv.Atoi(c.Param("paymentid"))
	req := md.UserPaymentReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}

	if err := handler.userService.UpdatePayment(req, userid, uint(paymentid)); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
		Code: http.StatusOK,
		Message: "Address Updated Successfully",
	})
}
func (handler UserHandler) SetDeletedPayment(c echo.Context) error {
	userid := c.Param("userid")
	paymentid,_ := strconv.Atoi(c.Param("paymentid"))
	if err := handler.userService.SetDeletedAddress(userid, uint(paymentid)); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Delete Payment Failed",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
			Code: http.StatusOK,
			Message: "Payment Set To Deleted",
		})
}

//Transaction
func (handler UserHandler) GetCartSession(c echo.Context) error {
	userid := c.Param("userid")
	cartSession, err := handler.userService.FindOrCreateCart(userid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Sever Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
		Code: http.StatusOK,
		Message: "Cart Session Found",
		Data: cartSession,
	})
}
func (handler UserHandler) AddItemToCart(c echo.Context) error {
	sessionId, _ := strconv.Atoi(c.Param("sessionid"))
	req := md.AddItemToCartReq{
		SessionID: uint(sessionId),
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}
	if err := handler.userService.AddItemToCart(req); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
		Code: http.StatusOK,
		Message: "Item Added Successfully",
	})
}
func (handler UserHandler) GetItemsCart(c echo.Context) error {
	sessionId, _ := strconv.Atoi(c.Param("sessionid"))
	cartItems, err := handler.userService.GetItemsCart(uint(sessionId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	if len(cartItems) <= 0 {
		return c.JSON(http.StatusOK, md.HttpResponse{
			Code: http.StatusOK,
			Message: "Success With Empty Data",
			Data: cartItems,
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
		Code: http.StatusOK,
		Message: "Items Found",
		Data: cartItems,
	})
}
func (handler UserHandler) DeleteItemFromCart(c echo.Context) error {
	cartItemId, _ := strconv.Atoi(c.Param("itemid"))
	sessionId, _ := strconv.Atoi(c.Param("sessionid"))
	if err := handler.userService.DeleteItemFromCart(uint(sessionId), uint(cartItemId)); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
		Code: http.StatusOK,
		Message: "Item Deleted Successfully",
	})
}
func (handler UserHandler) CreateOrder(c echo.Context) error {
	sessionId, _ := strconv.Atoi(c.Param("sessionid"))
	userid := c.Param("userid")
	req := md.CreateOrderReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}
	if err := handler.userService.CreateOrder(userid, uint(sessionId), req); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
		Code: http.StatusOK,
		Message: "Order Created Successfully",
	})
}
func (handler UserHandler) GetListOrders(c echo.Context) error {
	userid := c.Param("userid")
	orders, err := handler.userService.GetListOrders(userid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	if len(orders) <= 0 {
		return c.JSON(http.StatusOK, md.HttpResponse{
			Code: http.StatusOK,
			Message: "Success With Empty Data",
			Data: orders,
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
		Code: http.StatusOK,
		Message: "Orders Found",
		Data: orders,
	})
}
func (handler UserHandler) GetOrderById(c echo.Context) error {
	userid := c.Param("userid")
	orderid, _ := strconv.Atoi(c.Param("orderid"))

	order, err := handler.userService.GetOrder(userid, uint(orderid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
		Code: http.StatusOK,
		Message: "Orders Found",
		Data: order,
	})
}
func (handler UserHandler) UploadReceipt(c echo.Context) error {
	paymentid, _:= strconv.Atoi(c.Param("paymentid"))
	req := md.ReceiptURLReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data: err.Error(),
		})
	}
	if err := handler.userService.UploadReceipt(uint(paymentid), req.PaymentURL); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code: http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
		Code: http.StatusOK,
		Message: "Receipt Uploaded",
	})
}