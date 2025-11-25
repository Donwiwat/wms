package services

import (
	"wms-backend/internal/models"
	"wms-backend/internal/repositories"
)

type OrderService interface {
	GetAll() ([]models.OrderWithDetails, error)
	GetByID(id int) (*models.OrderWithDetails, error)
	Create(data *models.OrderRequest) (*models.Order, error)
	Update(id int, data *models.OrderRequest) (*models.Order, error)
	Delete(id int) error
	GetByCustomerID(customerID int) ([]models.OrderWithDetails, error)
	UpdateStatus(id int, status string) error
}

type orderService struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (s *orderService) GetAll() ([]models.OrderWithDetails, error) {
	return s.orderRepo.GetAll()
}

func (s *orderService) GetByID(id int) (*models.OrderWithDetails, error) {
	return s.orderRepo.GetByID(id)
}

func (s *orderService) Create(data *models.OrderRequest) (*models.Order, error) {
	return s.orderRepo.Create(data)
}

func (s *orderService) Update(id int, data *models.OrderRequest) (*models.Order, error) {
	return s.orderRepo.Update(id, data)
}

func (s *orderService) Delete(id int) error {
	return s.orderRepo.Delete(id)
}

func (s *orderService) GetByCustomerID(customerID int) ([]models.OrderWithDetails, error) {
	return s.orderRepo.GetByCustomerID(customerID)
}

func (s *orderService) UpdateStatus(id int, status string) error {
	return s.orderRepo.UpdateStatus(id, status)
}
