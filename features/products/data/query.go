package repository

import (
	"dot/config"
	"dot/features/products"
	"dot/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) products.DataInterface {
	return &Storage{
		query: db,
	}
}

func (data *Storage) GetProducts() ([]products.CoreProduct, string, error) {
	var model []models.Product

	tx := data.query.Find(&model)
	if tx.Error != nil {
		return nil, config.DBerror, tx.Error
	}

	core := models.ToCoreList(model)
	return core, "Success Get All Products", nil
}

func (data *Storage) PostProduct(core products.CoreProduct) (string, error) {
	model := models.CoreToModelProduct(core)

	tx := data.query.Create(&model)
	if tx.Error != nil {
		return config.DBerror, tx.Error
	}

	return "Success Post Product", nil
}

func (data *Storage) CheckOwner(core products.CoreProduct, updateid uint) (bool, string, error) {
	var model models.Product
	tx := data.query.Where("id = ?", updateid).First(&model)
	if tx.Error != nil {
		return false, config.DBerror, tx.Error
	}

	if model.UserID == core.UserID {
		return true, "Owner", nil
	} else {
		return false, "Not Owner", nil
	}
}

func (data *Storage) UpdatePut(core products.CoreProduct, updateid uint) (string, error) {
	update := models.CoreToModelProduct(core)

	tx := data.query.Model(&models.Product{}).Where("id = ? and user_id = ?", updateid, core.UserID).Updates(update)
	if tx.Error != nil {
		return config.DBerror, tx.Error
	}

	return "Success Updated Data", nil
}

func (data *Storage) UpdatePatch(core products.CoreProduct, updateid uint) (string, error) {
	update := models.CoreToModelProduct(core)

	tx := data.query.Model(&models.Product{}).Where("id = ? and user_id = ?", updateid, core.UserID).Updates(update)
	if tx.Error != nil {
		return config.DBerror, tx.Error
	}

	return "Success Updated Data", nil
}

func (data *Storage) CheckOwnerDel(userid, deleteid uint) (bool, string, error) {
	var model models.Product
	tx := data.query.Where("id = ?", deleteid).First(&model)
	if tx.Error != nil {
		return false, config.DBerror, tx.Error
	}

	if model.UserID == userid {
		return true, "Owner", nil
	} else {
		return false, "Not Owner", nil
	}
}

func (data *Storage) Delete(userid, deleteid uint) (string, error) {
	tx := data.query.Where("user_id = ? and id = ?", userid, deleteid).Delete(&models.Product{})
	if tx.Error != nil {
		return config.DBerror, tx.Error
	}

	return "Success Delete Product", nil
}

func (data *Storage) CheckQuantity(core products.CoreOrder) (string, error) {
	for i, v := range core.Buy {
		var temp models.Product
		tx := data.query.Where("id = ?", v).First(&temp)
		if tx.Error != nil {
			return config.DBerror, tx.Error
		} else if core.Quantity[i] > temp.Stock {
			msg := fmt.Sprintf("Stock For Product %d Is Not Enough", v)
			return msg, errors.New("error")
		}
	}
	return "Semua Stock Tersedia", nil
}

func (data *Storage) InsertOrder(core products.CoreOrder) (uint, string, error) {
	model := models.CoreToOrder(core, "failed")

	tx := data.query.Create(&model)
	if tx.Error != nil {
		return 0, config.DBerror, tx.Error
	}

	return model.ID, "Success Insert", nil
}

func (data *Storage) InsertOrderProduct(orderid uint, core products.CoreOrder) (string, error) {
	for i, v := range core.Buy {
		var temp models.Product
		tx := data.query.Where("id = ?", v).First(&temp)
		if tx.Error != nil {
			return config.DBerror, tx.Error
		}

		total := temp.Price * core.Quantity[i]
		model := models.CoreToOrderProduct(orderid, v, total)
		tx = data.query.Create(&model)
		if tx.Error != nil {
			return config.DBerror, tx.Error
		}
	}

	return "Succecc Insert", nil
}
func (data *Storage) GetTotal(orderid uint) (uint, string, error) {
	var order []models.OrderProduct
	tx := data.query.Where("order_id = ?", orderid).First(&order)
	if tx.Error != nil {
		return 0, config.DBerror, tx.Error
	}
	var total uint
	for _, v := range order {
		total += v.Total
	}

	return total, "Success Get Total", nil
}

func (data *Storage) UpdateData(orderid uint, core products.CoreOrder) (string, error) {
	model := models.CoreToOrder(core, "Success")
	model.ID = orderid

	for i, v := range core.Buy {
		var temp models.Product
		tx := data.query.Where("id = ?", v).First(&temp)
		if tx.Error != nil {
			return config.DBerror, tx.Error
		}
		temp.Stock -= core.Quantity[i]

		tx2 := data.query.Model(&models.Product{}).Where("id = ?", v).Updates(temp)
		if tx2.Error != nil {
			return config.DBerror, tx.Error
		}

	}
	fmt.Println(model)

	tx := data.query.Model(&models.Order{}).Where("id = ?", orderid).Updates(model)
	if tx.Error != nil {
		return config.DBerror, tx.Error
	}

	return "Success Order Product", nil
}
