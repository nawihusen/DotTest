package controller

import "dot/features/products"

type ResponseProduct struct {
	ID      uint   `json:"productid" form:"productid"`
	OwnerID uint   `json:"ownerid" form:"ownerid"`
	Name    string `json:"name" form:"name"`
	Price   uint   `json:"price" form:"price"`
	Stock   uint   `json:"stock" form:"stock"`
}

func CoreToResProduct(core products.CoreProduct) ResponseProduct {
	return ResponseProduct{
		ID:      core.ID,
		OwnerID: core.UserID,
		Name:    core.Name,
		Price:   core.Price,
		Stock:   core.Stock,
	}
}

func CoreToResProductList(core []products.CoreProduct) []ResponseProduct {
	var result []ResponseProduct
	for _, v := range core {
		result = append(result, CoreToResProduct(v))
	}
	return result
}
