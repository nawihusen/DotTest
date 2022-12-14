package controller

import (
	user "dot/features/users"
	"dot/middlewares"
	helpers "dot/utils/helpers"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	control user.ServiceInterface
}

func New(e *echo.Echo, data user.ServiceInterface) {
	handler := &Controller{
		control: data,
	}

	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)
	e.GET("/profile", handler.Profile, middlewares.JWTMiddleware())
	e.PUT("/profile", handler.EditProfilePut, middlewares.JWTMiddleware())
	e.PATCH("/profile", handler.EditProfilePatch, middlewares.JWTMiddleware())
	e.DELETE("/profile", handler.Delete, middlewares.JWTMiddleware())
}

func (client *Controller) Register(c echo.Context) error {
	var user UserRequest
	erb := c.Bind(&user)
	if erb != nil {
		return helpers.ErrorBind(c)
	}

	if user.Email == "" || user.Username == "" || user.Password == "" {
		return c.JSON(400, helpers.FailedResponseHelper("Data Can't Blank"))
	}

	msg, err := client.control.Register(user.ReqToCore())
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}

	return c.JSON(201, helpers.SuccessResponseHelper(msg))
}

func (client *Controller) Login(c echo.Context) error {
	var user LoginRequest
	erb := c.Bind(&user)

	if erb != nil {
		return helpers.ErrorBind(c)
	}

	msg, err := client.control.Login(user.ReqToCoreLogin())
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}

	return c.JSON(200, helpers.SuccessDataResponseHelper("Login Success", msg))
}

func (client *Controller) Profile(c echo.Context) error {
	userID := middlewares.ExtractToken(c)

	data, msg, err := client.control.GetProfile(userID)
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}
	// tambahkan ke respon terlebih dahulu
	return c.JSON(200, helpers.SuccessDataResponseHelper(msg, CoreToResUser(data)))
}

func (client *Controller) EditProfilePut(c echo.Context) error {
	userID := middlewares.ExtractToken(c)

	var user UpdateRequest
	erb := c.Bind(&user)
	if erb != nil {
		return helpers.ErrorBind(c)
	}

	msg, err := client.control.PutUpdate(userID, user.ReqToCoreUpdate())
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}

	return c.JSON(200, helpers.SuccessResponseHelper(msg))
}

func (client *Controller) Delete(c echo.Context) error {
	userID := middlewares.ExtractToken(c)

	msg, err := client.control.Delete(userID)
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}

	return c.JSON(200, helpers.SuccessResponseHelper(msg))
}

func (client *Controller) EditProfilePatch(c echo.Context) error {
	userID := middlewares.ExtractToken(c)

	var user UpdateRequest
	erb := c.Bind(&user)
	if erb != nil {
		return helpers.ErrorBind(c)
	}

	msg, err := client.control.PatchUpdate(userID, user.ReqToCoreUpdate())
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}

	return c.JSON(200, helpers.SuccessResponseHelper(msg))
}
