package datamodels

import "github.com/jinzhu/gorm"

type Order struct {
	ID          int64 `sql:Id gorm:'column:ID'`
	UserId      int64 `sql:userID gorm:'column:userId'`
	ProductId   int64 `sql:productID gorm:'column:productId'`
	OrderStatus int64 `sql:orderStatus gorm:'column:orderStatus'`
}

const (
	OrderWait    = iota
	OrderSuccess  //1
	OrderFailed   //2
)

func (u *Order) TableName(db *gorm.DB) string {
	return "order"
}
