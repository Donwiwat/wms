package repositories

import (
	"database/sql"
	"wms-backend/internal/models"
)

// ProductRepository interface defines product repository methods
type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id int) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id int) error
	List() ([]*models.Product, error)
	Search(query string) ([]*models.Product, error)
}

// productRepository implements ProductRepository
type productRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

// Create creates a new product
func (r *productRepository) Create(product *models.Product) error {
	query := `
		INSERT INTO products (name, short_name, brand, model, size, "group", unit1, unit2, ratio, cost, message, note)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(query,
		product.Name, product.ShortName, product.Brand, product.Model,
		product.Size, product.Group, product.Unit1, product.Unit2,
		product.Ratio, product.Cost, product.Message, product.Note).
		Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)
}

// GetByID gets a product by ID
func (r *productRepository) GetByID(id int) (*models.Product, error) {
	product := &models.Product{}
	query := `
		SELECT id, name, short_name, brand, model, size, "group", unit1, unit2, ratio, cost, message, note, created_at, updated_at
		FROM products WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&product.ID, &product.Name, &product.ShortName, &product.Brand,
		&product.Model, &product.Size, &product.Group, &product.Unit1,
		&product.Unit2, &product.Ratio, &product.Cost, &product.Message,
		&product.Note, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return product, nil
}

// Update updates a product
func (r *productRepository) Update(product *models.Product) error {
	query := `
		UPDATE products 
		SET name = $2, short_name = $3, brand = $4, model = $5, size = $6, "group" = $7, 
		    unit1 = $8, unit2 = $9, ratio = $10, cost = $11, message = $12, note = $13
		WHERE id = $1
		RETURNING updated_at`

	return r.db.QueryRow(query,
		product.ID, product.Name, product.ShortName, product.Brand,
		product.Model, product.Size, product.Group, product.Unit1,
		product.Unit2, product.Ratio, product.Cost, product.Message, product.Note).
		Scan(&product.UpdatedAt)
}

// Delete deletes a product
func (r *productRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// List gets all products
func (r *productRepository) List() ([]*models.Product, error) {
	query := `
		SELECT id, name, short_name, brand, model, size, "group", unit1, unit2, ratio, cost, message, note, created_at, updated_at
		FROM products ORDER BY name`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(
			&product.ID, &product.Name, &product.ShortName, &product.Brand,
			&product.Model, &product.Size, &product.Group, &product.Unit1,
			&product.Unit2, &product.Ratio, &product.Cost, &product.Message,
			&product.Note, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// Search searches products by name, short_name, brand, or model
func (r *productRepository) Search(query string) ([]*models.Product, error) {
	searchQuery := `
		SELECT id, name, short_name, brand, model, size, "group", unit1, unit2, ratio, cost, message, note, created_at, updated_at
		FROM products 
		WHERE name ILIKE $1 OR short_name ILIKE $1 OR brand ILIKE $1 OR model ILIKE $1
		ORDER BY name`

	searchTerm := "%" + query + "%"
	rows, err := r.db.Query(searchQuery, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(
			&product.ID, &product.Name, &product.ShortName, &product.Brand,
			&product.Model, &product.Size, &product.Group, &product.Unit1,
			&product.Unit2, &product.Ratio, &product.Cost, &product.Message,
			&product.Note, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
