package datamodels

type Order struct {
	ID          int64 `sq;:ID`
	UserId      int64 `sql:userID`
	ProductId   int64 `sql:productID`
	OrderStatus int64 `sql:orderStatus`
}

const (
	OrderWait    = iota
	OrderSuccess  //1
	OrderFailed   //2
)
