package services

import (
	"../datamodels"
	"../repositories"
	"github.com/kataras/iris/core/errors"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	GetIdUser(userId int64) (*datamodels.User, error)
	GetNameUser(userName string) (*datamodels.User, error)
	AddUser(user *datamodels.User) (int64, error)
	IsLoginSuccess(userName string, pwd string) (*datamodels.User, bool)
}

func NewUserService(table string) IUserService {
	return &UserManagerService{
		UserRepository: repositories.NewUserRepository(table),
	}

}

type UserManagerService struct {
	UserRepository repositories.IUserRepository
}

func (s *UserManagerService) GetIdUser(userId int64) (*datamodels.User, error) {
	panic("implement me")
}

func (s *UserManagerService) GetNameUser(userName string) (*datamodels.User, error) {
	panic("implement me")
}

func (s *UserManagerService) AddUser(user *datamodels.User) (int64, error) {
	if user == nil || user.UserName == "" || user.HashPassword == "" {
		return 0, errors.New("用户信息不完整！")
	}
	enPwd, err := GeneratePwd(user.HashPassword)
	if err != nil {
		return 0, errors.New("密码异常！")
	}
	user.HashPassword = string(enPwd)
	return s.UserRepository.Insert(user)
}

func (s *UserManagerService) IsLoginSuccess(userName string, uiPwd string) (user *datamodels.User, isOk bool) {
	if userName == "" || uiPwd == "" {
		return nil, false
	}
	user, err := s.UserRepository.Select(userName)
	if err != nil {
		return nil, false
	}
	if isOk, _ := ValidatePwd(user.HashPassword, uiPwd); isOk {
		return nil, false
	}
	return user, true
}

func ValidatePwd(dbPwd string, uiPwd string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(uiPwd)); err != nil {
		return false, err
	}
	return true, nil
}

func GeneratePwd(uiPwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(uiPwd), bcrypt.DefaultCost)
}
