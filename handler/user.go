package handler

import (
	md "ems-aquadev/models"
	svc "ems-aquadev/service"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *svc.UserService
}

func NewUserHandler(userService *svc.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (handler UserHandler) CreateUserTransaction(c echo.Context) error {
	req := md.UserRegRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code: http.StatusBadRequest,
			Message: "Register user failed",
			Data: err.Error(),
		})
	}

	createdUser, err := handler.userService.CreateUserTransaction(req)
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