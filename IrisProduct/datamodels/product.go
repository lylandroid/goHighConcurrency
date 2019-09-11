package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID" imooc:"id"`
	ProductName  string `json:"ProductName" sql:"productName" imooc:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" imooc:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" imooc:"productImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" imooc:"ProductUrl"`
}
