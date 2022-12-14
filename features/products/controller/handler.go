package controller

import (
	"dot/features/products"
	"dot/middlewares"
	"dot/utils/database"
	"dot/utils/helpers"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	control products.ServiceInterface
}

func New(e *echo.Echo, data products.ServiceInterface) {
	handler := &Controller{
		control: data,
	}

	e.GET("/product", handler.GetProducts, middlewares.JWTMiddleware())
	e.POST("/profile/product", handler.PostProduct, middlewares.JWTMiddleware())
	e.PUT("/profile/product", handler.UpdatePut, middlewares.JWTMiddleware())
	e.PATCH("/profile/product", handler.UpdatePatch, middlewares.JWTMiddleware())
	e.DELETE("/profile/product", handler.Delete, middlewares.JWTMiddleware())
	e.POST("/orders", handler.Order, middlewares.JWTMiddleware())

}

func (client *Controller) GetProducts(c echo.Context) error {
	products, msg, err := client.control.GetProducts()
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}

	database.SetDataRedis(products)

	return c.JSON(200, helpers.SuccessDataResponseHelper(msg, CoreToResProductList(products)))
}

func (client *Controller) PostProduct(c echo.Context) error {
	userID := middlewares.ExtractToken(c)

	var user ProductRequest
	erb := c.Bind(&user)
	if erb != nil {
		return helpers.ErrorBind(c)
	}

	if user.Name == "" {
		return c.JSON(400, helpers.FailedResponseHelper("Product Name Can't Blank"))
	}
	msg, err := client.control.PostProduct(user.ProductRequestToCore(userID))
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}

	return c.JSON(201, helpers.SuccessResponseHelper(msg))
}

func (client *Controller) UpdatePut(c echo.Context) error {
	userID := middlewares.ExtractToken(c)

	var user UpdateRequest
	erb := c.Bind(&user)
	if erb != nil {
		return helpers.ErrorBind(c)
	}
	if user.UpdateID == 0 {
		return c.JSON(400, helpers.FailedResponseHelper("Please Add Update ID"))
	}

	msg, err := client.control.UpdatePut(user.ProductRequestToCore(userID), user.UpdateID)
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}

	return c.JSON(200, helpers.SuccessResponseHelper(msg))
}

func (client *Controller) UpdatePatch(c echo.Context) error {
	userID := middlewares.ExtractToken(c)

	var user UpdateRequest
	erb := c.Bind(&user)
	if erb != nil {
		return helpers.ErrorBind(c)
	}

	if user.UpdateID == 0 {
		return c.JSON(400, helpers.FailedResponseHelper("Please Add Update ID"))
	}

	msg, err := client.control.UpdatePatch(user.ProductRequestToCore(userID), user.UpdateID)
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}

	return c.JSON(200, helpers.SuccessResponseHelper(msg))
}

func (client *Controller) Delete(c echo.Context) error {
	userID := middlewares.ExtractToken(c)

	var deleteid Delete
	erb := c.Bind(&deleteid)
	if erb != nil {
		return helpers.ErrorBind(c)
	}
	if deleteid.DeleteID == 0 {
		return c.JSON(400, helpers.FailedResponseHelper("Please Add Product ID"))
	}

	msg, err := client.control.Delete(userID, deleteid.DeleteID)
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}

	return c.JSON(200, helpers.SuccessResponseHelper(msg))
}

func (client *Controller) Order(c echo.Context) error {
	userID := middlewares.ExtractToken(c)

	var user OrderRequest
	erb := c.Bind(&user)
	if erb != nil {
		return helpers.ErrorBind(c)
	}

	if len(user.Buy) != len(user.Quantity) {
		return c.JSON(400, helpers.FailedResponseHelper("Buy and Quantity mush have same amount of number"))
	}

	msg, err := client.control.Order(user.RequestToCore(userID))
	if err != nil {
		return helpers.ErrorInternal(c, msg)
	}

	return c.JSON(200, helpers.SuccessResponseHelper(msg))
}
