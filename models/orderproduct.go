package models

type OrderProduct struct {
	OrderID   uint
	ProductID uint
	Total     uint
}

func CoreToOrderProduct(orderId, produkid, total uint) OrderProduct {
	return OrderProduct{
		OrderID:   orderId,
		ProductID: produkid,
		Total:     total,
	}
}
