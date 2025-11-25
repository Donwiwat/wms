package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"wms-backend/internal/models"
)

type CustomerRepository interface {
	GetAll() ([]models.Customer, error)
	GetByID(id int) (*models.Customer, error)
	Create(customer *models.Customer) error
	Update(id int, customer *models.Customer) error
	Delete(id int) error
	Search(query string) ([]models.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) GetAll() ([]models.Customer, error) {
	query := `
		SELECT id, prefix, name, address, phone, contact_person, level, 
		       delivery_place, transport, credit_limit, credit_term, 
		       outstanding, last_contact, note, created_at, updated_at
		FROM customers 
		ORDER BY name ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get customers: %w", err)
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(
			&customer.ID, &customer.Prefix, &customer.Name, &customer.Address,
			&customer.Phone, &customer.ContactPerson, &customer.Level,
			&customer.DeliveryPlace, &customer.Transport, &customer.CreditLimit,
			&customer.CreditTerm, &customer.Outstanding, &customer.LastContact,
			&customer.Note, &customer.CreatedAt, &customer.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan customer: %w", err)
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (r *customerRepository) GetByID(id int) (*models.Customer, error) {
	query := `
		SELECT id, prefix, name, address, phone, contact_person, level, 
		       delivery_place, transport, credit_limit, credit_term, 
		       outstanding, last_contact, note, created_at, updated_at
		FROM customers 
		WHERE id = $1
	`

	var customer models.Customer
	err := r.db.QueryRow(query, id).Scan(
		&customer.ID, &customer.Prefix, &customer.Name, &customer.Address,
		&customer.Phone, &customer.ContactPerson, &customer.Level,
		&customer.DeliveryPlace, &customer.Transport, &customer.CreditLimit,
		&customer.CreditTerm, &customer.Outstanding, &customer.LastContact,
		&customer.Note, &customer.CreatedAt, &customer.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("customer not found")
		}
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}

	return &customer, nil
}

func (r *customerRepository) Create(customer *models.Customer) error {
	query := `
		INSERT INTO customers (prefix, name, address, phone, contact_person, level, 
		                      delivery_place, transport, credit_limit, credit_term, 
		                      outstanding, last_contact, note)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(
		query,
		customer.Prefix, customer.Name, customer.Address, customer.Phone,
		customer.ContactPerson, customer.Level, customer.DeliveryPlace,
		customer.Transport, customer.CreditLimit, customer.CreditTerm,
		customer.Outstanding, customer.LastContact, customer.Note,
	).Scan(&customer.ID, &customer.CreatedAt, &customer.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create customer: %w", err)
	}

	return nil
}

func (r *customerRepository) Update(id int, customer *models.Customer) error {
	query := `
		UPDATE customers 
		SET prefix = $2, name = $3, address = $4, phone = $5, contact_person = $6, 
		    level = $7, delivery_place = $8, transport = $9, credit_limit = $10, 
		    credit_term = $11, outstanding = $12, last_contact = $13, note = $14
		WHERE id = $1
		RETURNING updated_at
	`

	err := r.db.QueryRow(
		query, id,
		customer.Prefix, customer.Name, customer.Address, customer.Phone,
		customer.ContactPerson, customer.Level, customer.DeliveryPlace,
		customer.Transport, customer.CreditLimit, customer.CreditTerm,
		customer.Outstanding, customer.LastContact, customer.Note,
	).Scan(&customer.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("customer not found")
		}
		return fmt.Errorf("failed to update customer: %w", err)
	}

	customer.ID = id
	return nil
}

func (r *customerRepository) Delete(id int) error {
	query := `DELETE FROM customers WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("customer not found")
	}

	return nil
}

func (r *customerRepository) Search(query string) ([]models.Customer, error) {
	searchQuery := `
		SELECT id, prefix, name, address, phone, contact_person, level, 
		       delivery_place, transport, credit_limit, credit_term, 
		       outstanding, last_contact, note, created_at, updated_at
		FROM customers 
		WHERE LOWER(name) LIKE LOWER($1) 
		   OR LOWER(contact_person) LIKE LOWER($1)
		   OR LOWER(phone) LIKE LOWER($1)
		ORDER BY name ASC
	`

	searchTerm := "%" + strings.ToLower(query) + "%"
	rows, err := r.db.Query(searchQuery, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to search customers: %w", err)
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(
			&customer.ID, &customer.Prefix, &customer.Name, &customer.Address,
			&customer.Phone, &customer.ContactPerson, &customer.Level,
			&customer.DeliveryPlace, &customer.Transport, &customer.CreditLimit,
			&customer.CreditTerm, &customer.Outstanding, &customer.LastContact,
			&customer.Note, &customer.CreatedAt, &customer.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan customer: %w", err)
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
