package services

import (
	"database/sql"
	"wms-backend/internal/models"
	"wms-backend/internal/repositories"
)

type CustomerService interface {
	GetAll() ([]models.Customer, error)
	GetByID(id int) (*models.Customer, error)
	Create(data *models.CustomerFormData) (*models.Customer, error)
	Update(id int, data *models.CustomerFormData) (*models.Customer, error)
	Delete(id int) error
	Search(query string) ([]models.Customer, error)
}

type customerService struct {
	customerRepo repositories.CustomerRepository
}

func NewCustomerService(customerRepo repositories.CustomerRepository) CustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}

func (s *customerService) GetAll() ([]models.Customer, error) {
	return s.customerRepo.GetAll()
}

func (s *customerService) GetByID(id int) (*models.Customer, error) {
	return s.customerRepo.GetByID(id)
}

func (s *customerService) Create(data *models.CustomerFormData) (*models.Customer, error) {
	customer := &models.Customer{
		Prefix:        sql.NullString{String: data.Prefix, Valid: data.Prefix != ""},
		Name:          data.Name,
		Address:       sql.NullString{String: data.Address, Valid: data.Address != ""},
		Phone:         sql.NullString{String: data.Phone, Valid: data.Phone != ""},
		ContactPerson: sql.NullString{String: data.ContactPerson, Valid: data.ContactPerson != ""},
		Level:         sql.NullString{String: data.Level, Valid: data.Level != ""},
		DeliveryPlace: sql.NullString{String: data.DeliveryPlace, Valid: data.DeliveryPlace != ""},
		Transport:     sql.NullString{String: data.Transport, Valid: data.Transport != ""},
		CreditLimit:   data.CreditLimit,
		CreditTerm:    data.CreditTerm,
		Outstanding:   data.Outstanding,
		LastContact: func() sql.NullTime {
			if data.LastContact != nil {
				return sql.NullTime{Time: *data.LastContact, Valid: true}
			}
			return sql.NullTime{Valid: false}
		}(),
		Note: sql.NullString{String: data.Note, Valid: data.Note != ""},
	}

	err := s.customerRepo.Create(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *customerService) Update(id int, data *models.CustomerFormData) (*models.Customer, error) {
	customer := &models.Customer{
		Prefix:        sql.NullString{String: data.Prefix, Valid: data.Prefix != ""},
		Name:          data.Name,
		Address:       sql.NullString{String: data.Address, Valid: data.Address != ""},
		Phone:         sql.NullString{String: data.Phone, Valid: data.Phone != ""},
		ContactPerson: sql.NullString{String: data.ContactPerson, Valid: data.ContactPerson != ""},
		Level:         sql.NullString{String: data.Level, Valid: data.Level != ""},
		DeliveryPlace: sql.NullString{String: data.DeliveryPlace, Valid: data.DeliveryPlace != ""},
		Transport:     sql.NullString{String: data.Transport, Valid: data.Transport != ""},
		CreditLimit:   data.CreditLimit,
		CreditTerm:    data.CreditTerm,
		Outstanding:   data.Outstanding,
		LastContact: func() sql.NullTime {
			if data.LastContact != nil {
				return sql.NullTime{Time: *data.LastContact, Valid: true}
			}
			return sql.NullTime{Valid: false}
		}(),
		Note: sql.NullString{String: data.Note, Valid: data.Note != ""},
	}

	err := s.customerRepo.Update(id, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *customerService) Delete(id int) error {
	return s.customerRepo.Delete(id)
}

func (s *customerService) Search(query string) ([]models.Customer, error) {
	return s.customerRepo.Search(query)
}
