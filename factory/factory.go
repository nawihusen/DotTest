package factory

import (
	userController "dot/features/users/controller"
	userData "dot/features/users/data"
	userService "dot/features/users/service"

	productController "dot/features/products/controller"
	productData "dot/features/products/data"
	productService "dot/features/products/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	userDataFactory := userData.New(db)
	userUsecaseFactory := userService.New(userDataFactory)
	userController.New(e, userUsecaseFactory)

	productDataFactory := productData.New(db)
	productUsecaseFactory := productService.New(productDataFactory)
	productController.New(e, productUsecaseFactory)

}
