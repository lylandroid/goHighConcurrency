package datamodels

import "github.com/jinzhu/gorm"

type Order struct {
	ID          int64 `gorm:"column:ID" sql:Id`
	UserId      int64 `gorm:"column:userId" sql:userID`
	ProductId   int64 `gorm:"column:productId" sql:productID`
	OrderStatus int64 `gorm:"column:orderStatus" sql:orderStatus`
}

const (
	OrderWait    = iota
	OrderSuccess  //1
	OrderFailed   //2
)

func (u *Order) TableName(db *gorm.DB) string {
	return "order"
}
