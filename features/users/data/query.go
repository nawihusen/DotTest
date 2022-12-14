package repository

import (
	"dot/config"
	user "dot/features/users"
	"dot/models"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &Storage{
		query: db,
	}
}

func (data *Storage) Create(core user.CoreUser) (string, error) {
	model := models.CoreToModel(core)

	tx := data.query.Create(&model)
	if tx.Error != nil {
		return config.DBerror, tx.Error
	}

	return "Created", nil
}

func (data *Storage) Login(email string) (user.CoreUser, error) {
	var model models.User

	txEmail := data.query.Where("email = ?", email).First(&model)
	if txEmail.Error != nil {
		return user.CoreUser{}, txEmail.Error
	}

	dataUser := model.ModelToCore()

	return dataUser, nil
}

func (data *Storage) GetProfile(userid uint) (user.CoreUser, string, error) {
	var model models.User

	tx := data.query.Where("id = ?", userid).First(&model)
	if tx.Error != nil {
		return user.CoreUser{}, config.DBerror, tx.Error
	}

	dataUser := model.ModelToCore()

	return dataUser, "Your Profile", nil
}

func (data *Storage) PutUpdate(userid uint, core user.CoreUser) (string, error) {
	update := models.CoreToModel(core)

	tx := data.query.Model(&models.User{}).Where("id = ?", userid).Updates(update)
	if tx.Error != nil {
		return config.DBerror, tx.Error
	}

	return "Success To Update Data", nil
}

func (data *Storage) PatchUpdate(userid uint, core user.CoreUser) (string, error) {
	update := models.CoreToModel(core)

	tx := data.query.Model(&models.User{}).Where("id = ?", userid).Updates(update)
	if tx.Error != nil {
		return config.DBerror, tx.Error
	}

	return "Success To Update Data", nil
}

func (data *Storage) Delete(userid uint) (string, error) {

	tx := data.query.Where("id = ?", userid).Delete(&models.User{})
	// tx := data.query.Where("id = ?", userid).Unscoped().Delete(&models.User{})
	if tx.Error != nil {
		return config.DBerror, tx.Error
	}

	return "Success Delete Account", nil
}
