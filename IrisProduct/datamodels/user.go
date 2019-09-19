package datamodels

import "github.com/jinzhu/gorm"

type User struct {
	ID       uint   `gorm:"primary_key"`
	NickName string `form:"nickName"`
	UserName string `form:"userName"`
	Password string `form:"password"`
}

func (u *User) TableName(db *gorm.DB) string {
	return "user"
}
