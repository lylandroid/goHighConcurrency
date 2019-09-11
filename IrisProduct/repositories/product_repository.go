package repositories

import (
	"../datamodels"
)

type IProduct interface {
	//连接数据库
	Conn() error
	Insert(product *datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(int64) *datamodels.Product
	SelectAll() []**datamodels.Product
}
