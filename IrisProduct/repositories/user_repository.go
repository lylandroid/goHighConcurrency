package repositories

import (
	"../common"
	"../datamodels"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/core/errors"
)

type IUserRepository interface {
	Conn() error
	Insert(user *datamodels.User) (int64, error)
	Delete(int64) bool
	Update(user *datamodels.User) error
	SelectById(id int64) (user *datamodels.User, err error)
	Select(userName string) (user *datamodels.User, err error)
}

func NewUserRepository(table string) IUserRepository {
	return &UserManagerRepository{
		table: table,
	}
}

type UserManagerRepository struct {
	table     string
	mySqlConn *gorm.DB
}

func (r *UserManagerRepository) Conn() error {
	if r.mySqlConn == nil {
		db, err := common.NewMySqlGormConn()
		if err != nil {
			return err
		}
		r.mySqlConn = db
		if r.table == "" {
			r.table = "user"
		}
	}
	return nil
}

func (r *UserManagerRepository) Insert(user *datamodels.User) (id int64, err error) {
	if user == nil || user.UserName == "" || user.Password == "" {
		return 0, errors.New("用户信息不完整！")
	}
	if err = r.Conn(); err != nil {
		return 0, err
	}
	if !r.mySqlConn.NewRecord(user) {
		return 0, errors.New("主键不为空！")
	}
	db := r.mySqlConn.Create(user)
	if db.Error != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *UserManagerRepository) Delete(int64) bool {
	panic("implement me")
}

func (r *UserManagerRepository) Update(user *datamodels.User) error {
	panic("implement me")
}

func (r *UserManagerRepository) Select(userName string) (user *datamodels.User, err error) {
	if userName == "" {
		return &datamodels.User{}, errors.New("条件不能为空！")
	}
	if err = r.Conn(); err != nil {
		return nil, err
	}
	user = &datamodels.User{}
	db := r.mySqlConn.First(user, "user_name=?", userName)
	//defer db.Close()
	if db.Error != nil {
		return user, db.Error
	}
	return user, nil
}

func (r *UserManagerRepository) SelectById(userId int64) (user *datamodels.User, err error) {
	if userId < 1 {
		return &datamodels.User{}, errors.New("id 未知错误")
	}
	if err = r.Conn(); err != nil {
		return nil, err
	}
	user = &datamodels.User{}
	db := r.mySqlConn.First(user, userId)
	if db.Error != nil {
		return nil, db.Error
	}
	return user, nil
}
