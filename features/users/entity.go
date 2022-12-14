package user

type CoreUser struct {
	ID       uint
	Username string
	Email    string
	Password string
}

type ServiceInterface interface {
	Register(CoreUser) (string, error)
	Login(CoreUser) (string, error)
	GetProfile(uint) (CoreUser, string, error)
	PutUpdate(uint, CoreUser) (string, error)
	PatchUpdate(uint, CoreUser) (string, error)
	Delete(uint) (string, error)
}

type DataInterface interface {
	Create(CoreUser) (string, error)
	Login(string) (CoreUser, error)
	GetProfile(uint) (CoreUser, string, error)
	PutUpdate(uint, CoreUser) (string, error)
	PatchUpdate(uint, CoreUser) (string, error)
	Delete(uint) (string, error)
}
