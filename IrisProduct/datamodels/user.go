package datamodels

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID           int64  `json:"id" form:"ID" sql:"ID" gorm:"primary_key"`
	NickName     string `json:"nickName" form:"nickName" sql:"nickName"`
	UserName     string `json:"userName" form:"userName" sql:"userName"`
	HashPassword string /*`json:"-" form:"passWord" sql:"passWord"`*/`json:"-" form:"password" sql:"password"`
}
