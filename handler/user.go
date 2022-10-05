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
			Message: "Internal Sever Error",
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