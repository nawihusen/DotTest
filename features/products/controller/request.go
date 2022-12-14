package controller

import "dot/features/products"

type ProductRequest struct {
	UserID uint
	Name   string `json:"name" form:"name" validate:"required"`
	Price  uint   `json:"price" form:"price" validate:"required"`
	Stock  uint   `json:"stock" form:"stock" validate:"required"`
}

type UpdateRequest struct {
	UserID   uint
	UpdateID uint   `json:"updateid" form:"updateid" validate:"required"`
	Name     string `json:"name" form:"name"`
	Price    uint   `json:"price" form:"price"`
	Stock    uint   `json:"stock" form:"stock"`
}

type Delete struct {
	DeleteID uint `json:"productid" form:"productid"`
}

type OrderRequest struct {
	UserID   uint
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Nomor    string `json:"nomor" form:"nomor"`
	Buy      []uint `json:"buy" form:"buy"`
	Quantity []uint `json:"quantity" form:"quantity"`
}

func (client *ProductRequest) ProductRequestToCore(userid uint) products.CoreProduct {
	return products.CoreProduct{
		UserID: userid,
		Name:   client.Name,
		Price:  client.Price,
		Stock:  client.Stock,
	}
}

func (client *UpdateRequest) ProductRequestToCore(userid uint) products.CoreProduct {
	return products.CoreProduct{
		UserID: userid,
		Name:   client.Name,
		Price:  client.Price,
		Stock:  client.Stock,
	}
}

func (client *OrderRequest) RequestToCore(userid uint) products.CoreOrder {
	return products.CoreOrder{
		UserID:   userid,
		Name:     client.Name,
		Address:  client.Address,
		Nomor:    client.Nomor,
		Buy:      client.Buy,
		Quantity: client.Quantity,
	}
}
