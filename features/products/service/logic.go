package service

import (
	"dot/features/products"
	"errors"
)

type Service struct {
	do products.DataInterface
}

func New(data products.DataInterface) products.ServiceInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) GetProducts() ([]products.CoreProduct, string, error) {
	data, msg, err := service.do.GetProducts()
	return data, msg, err
}

func (service *Service) PostProduct(core products.CoreProduct) (string, error) {
	msg, err := service.do.PostProduct(core)
	return msg, err
}

func (service *Service) UpdatePut(core products.CoreProduct, updateid uint) (string, error) {
	own, msg, er := service.do.CheckOwner(core, updateid)
	if er != nil {
		return msg, er
	} else if own == false {
		return "Cant Update Others Product", errors.New("Error")
	}

	msg, er = service.do.UpdatePut(core, updateid)
	return msg, er
}

func (service *Service) UpdatePatch(core products.CoreProduct, updateid uint) (string, error) {
	own, msg, er := service.do.CheckOwner(core, updateid)
	if er != nil {
		return msg, er
	} else if own == false {
		return "Cant Update Others Product", errors.New("Error")
	}
	msg, er = service.do.UpdatePatch(core, updateid)
	return msg, er
}

func (service *Service) Delete(userid, deleteid uint) (string, error) {
	own, msg, er := service.do.CheckOwnerDel(userid, deleteid)
	if er != nil {
		return msg, er
	} else if own == false {
		return "Cant Delete Others Product", errors.New("Error")
	}

	msg, er = service.do.Delete(userid, deleteid)
	return msg, er
}

func (service *Service) Order(order products.CoreOrder) (string, error) {
	msg, err := service.do.CheckQuantity(order)
	if err != nil {
		return msg, err
	}

	orderid, msg, err := service.do.InsertOrder(order)
	if err != nil {
		return msg, err
	}

	msg, err = service.do.InsertOrderProduct(orderid, order)
	if err != nil {
		return msg, err
	}

	total, msg, err := service.do.GetTotal(orderid)
	if err != nil {
		return msg, err
	}
	order.Total = total
	msg, err = service.do.UpdateData(orderid, order)
	if err != nil {
		return msg, err
	}

	return msg, nil
}
