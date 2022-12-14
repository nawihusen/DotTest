package service

import (
	user "dot/features/users"
	"dot/middlewares"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	do user.DataInterface
}

func New(data user.DataInterface) user.ServiceInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) Register(core user.CoreUser) (string, error) {
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(core.Password), bcrypt.DefaultCost)
	core.Password = string(hashPass)

	msg, err := service.do.Create(core)
	return msg, err
}

func (service *Service) Login(core user.CoreUser) (string, error) {
	results, errEmail := service.do.Login(core.Email)

	if errEmail != nil {
		return "email not found", errors.New("Error")
	}

	errPw := bcrypt.CompareHashAndPassword([]byte(results.Password), []byte(core.Password))
	if errPw != nil {
		return "wrong password", errors.New("Error")
	}

	token, errToken := middlewares.CreateToken(int(results.ID))

	return token, errToken
}

func (service *Service) GetProfile(userid uint) (user.CoreUser, string, error) {
	data, msg, err := service.do.GetProfile(userid)

	return data, msg, err
}

func (service *Service) PutUpdate(userid uint, core user.CoreUser) (string, error) {
	msg, err := service.do.PutUpdate(userid, core)
	return msg, err
}

func (service *Service) PatchUpdate(userid uint, core user.CoreUser) (string, error) {
	msg, err := service.do.PatchUpdate(userid, core)
	return msg, err
}

func (service *Service) Delete(userid uint) (string, error) {
	msg, err := service.do.Delete(userid)

	return msg, err
}
