package services

import (
	"../datamodels"
	"../repositories"
	"database/sql"
)

type IProductService interface {
	GetProduct(productId int64) (*datamodels.Product, error)
	GetProductAll() ([]*datamodels.Product, error)
	InsertProduct(product *datamodels.Product) (int64, error)
	UpdateProduct(product *datamodels.Product) error
	DeleteProduct(productId int64) bool
}

func NewProductServiceImp(tableName string, db *sql.DB) IProductService {
	return &ProductServiceImp{
		IProductRepository: repositories.NewProductRepositoryImp(tableName, db),
	}
}

type ProductServiceImp struct {
	IProductRepository repositories.IProductRepository
}

func (p *ProductServiceImp) GetProduct(productId int64) (*datamodels.Product, error) {
	return p.IProductRepository.SelectByKey(productId)
}

func (p *ProductServiceImp) GetProductAll() (products []*datamodels.Product, err error) {
	return p.IProductRepository.SelectAll()
}

func (p *ProductServiceImp) InsertProduct(product *datamodels.Product) (int64, error) {
	return p.IProductRepository.Insert(product)
}

func (p *ProductServiceImp) UpdateProduct(product *datamodels.Product) error {
	return p.IProductRepository.Update(product)
}

func (p *ProductServiceImp) DeleteProduct(productId int64) bool {
	return p.IProductRepository.Delete(productId)
}
