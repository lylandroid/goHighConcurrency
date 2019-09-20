package datamodels

import "github.com/jinzhu/gorm"

type User struct {
	ID int64 `gorm:"primary_key:true" form:"id"`
	//ID       int64  `form:"id"` //`gorm:"primary_key"`
	NickName string `form:"nickName"`
	UserName string `form:"userName"`
	Password string `form:"password"`
}

func (u *User) TableName(db *gorm.DB) string {
	return "user"
}
