package products

type CoreProduct struct {
	ID     uint
	UserID uint
	Name   string
	Price  uint
	Stock  uint
}

type CoreOrder struct {
	ID       uint
	UserID   uint
	Name     string
	Address  string
	Nomor    string
	Total    uint
	Buy      []uint
	Quantity []uint
}

type ServiceInterface interface {
	GetProducts() ([]CoreProduct, string, error)
	PostProduct(CoreProduct) (string, error)
	UpdatePut(CoreProduct, uint) (string, error)
	UpdatePatch(CoreProduct, uint) (string, error)
	Delete(uint, uint) (string, error)
	Order(CoreOrder) (string, error)
}

type DataInterface interface {
	GetProducts() ([]CoreProduct, string, error)
	PostProduct(CoreProduct) (string, error)
	UpdatePut(CoreProduct, uint) (string, error)
	UpdatePatch(CoreProduct, uint) (string, error)
	CheckOwner(CoreProduct, uint) (bool, string, error)
	CheckOwnerDel(uint, uint) (bool, string, error)
	Delete(uint, uint) (string, error)
	//
	CheckQuantity(CoreOrder) (string, error)
	InsertOrder(CoreOrder) (uint, string, error)
	InsertOrderProduct(uint, CoreOrder) (string, error)
	GetTotal(uint) (uint, string, error)
	UpdateData(uint, CoreOrder) (string, error)
}
