package models

import (
	"dot/features/products"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID        uint
	Name          string
	Price         uint `gorm:"default:0"`
	Stock         uint `gorm:"default:0"`
	OrderProducts []OrderProduct
}

func CoreToModelProduct(core products.CoreProduct) Product {
	return Product{
		UserID: core.UserID,
		Name:   core.Name,
		Price:  core.Price,
		Stock:  core.Stock,
	}
}

func ToCore(prod Product) products.CoreProduct {
	return products.CoreProduct{
		ID:     prod.ID,
		UserID: prod.UserID,
		Name:   prod.Name,
		Price:  prod.Price,
		Stock:  prod.Stock,
	}
}

func ToCoreList(prod []Product) []products.CoreProduct {
	result := []products.CoreProduct{}
	for _, v := range prod {
		result = append(result, ToCore(v))
	}

	return result
}
