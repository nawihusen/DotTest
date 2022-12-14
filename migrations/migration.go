package migrations

import (
	"dot/models"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Product{})
	db.AutoMigrate(models.Order{})
	db.AutoMigrate(models.OrderProduct{})
}
