package services

import (
	"../datamodels"
	"../repositories"
	"database/sql"
)

type IOrderService interface {
	GetOrderById(int64) (*datamodels.Order, error)
	DeleteOrderById(int64) bool
	UpdateOrder(order *datamodels.Order) error
	InsertOrder(order *datamodels.Order) (int64, error)
	GetAllOrder() ([]*datamodels.Order, error)
	GetAllOrderInfo() (map[int]map[string]string, error)
}

func NewOrderService(table string, db *sql.DB) IOrderService {
	return &OrderService{
		OrderRepository: repositories.NewOrderManagerRepository(table, db),
	}
}

type OrderService struct {
	OrderRepository repositories.IOrderRepository
}

func (s *OrderService) GetOrderById(orderId int64) (*datamodels.Order, error) {
	return s.OrderRepository.SelectByKey(orderId)
}

func (s *OrderService) DeleteOrderById(orderId int64) bool {
	return s.OrderRepository.Delete(orderId)
}

func (s *OrderService) UpdateOrder(order *datamodels.Order) error {
	return s.OrderRepository.Update(order)
}

func (s *OrderService) InsertOrder(order *datamodels.Order) (int64, error) {
	return s.OrderRepository.Insert(order)
}

func (s *OrderService) GetAllOrder() ([]*datamodels.Order, error) {
	return s.OrderRepository.SelectAll()
}

func (s *OrderService) GetAllOrderInfo() (map[int]map[string]string, error) {
	return s.OrderRepository.SelectAllWithInfo()
}
