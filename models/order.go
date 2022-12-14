package models

import (
	"dot/features/products"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID        uint
	Name          string
	Address       string
	Nomor         string
	Total         uint   `gorm:"default:0"`
	Status        string `gorm:"default: failed"`
	OrderProducts []OrderProduct
}

func CoreToOrder(core products.CoreOrder, status string) Order {
	return Order{
		UserID:  core.UserID,
		Name:    core.Name,
		Address: core.Address,
		Nomor:   core.Nomor,
		Status:  status,
		Total:   core.Total,
	}
}
