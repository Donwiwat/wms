package repositories

import (
	"database/sql"
	"wms-backend/internal/models"
)

// WarehouseRepository interface defines warehouse repository methods
type WarehouseRepository interface {
	Create(warehouse *models.Warehouse) error
	GetByID(id int) (*models.Warehouse, error)
	Update(warehouse *models.Warehouse) error
	Delete(id int) error
	List() ([]*models.Warehouse, error)
}

// warehouseRepository implements WarehouseRepository
type warehouseRepository struct {
	db *sql.DB
}

// NewWarehouseRepository creates a new warehouse repository
func NewWarehouseRepository(db *sql.DB) WarehouseRepository {
	return &warehouseRepository{db: db}
}

// Create creates a new warehouse
func (r *warehouseRepository) Create(warehouse *models.Warehouse) error {
	query := `
		INSERT INTO warehouses (name, location, note)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(query, warehouse.Name, warehouse.Location, warehouse.Note).
		Scan(&warehouse.ID, &warehouse.CreatedAt, &warehouse.UpdatedAt)
}

// GetByID gets a warehouse by ID
func (r *warehouseRepository) GetByID(id int) (*models.Warehouse, error) {
	warehouse := &models.Warehouse{}
	query := `
		SELECT id, name, location, note, created_at, updated_at
		FROM warehouses WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&warehouse.ID, &warehouse.Name, &warehouse.Location,
		&warehouse.Note, &warehouse.CreatedAt, &warehouse.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return warehouse, nil
}

// Update updates a warehouse
func (r *warehouseRepository) Update(warehouse *models.Warehouse) error {
	query := `
		UPDATE warehouses 
		SET name = $2, location = $3, note = $4
		WHERE id = $1
		RETURNING updated_at`

	return r.db.QueryRow(query, warehouse.ID, warehouse.Name, warehouse.Location, warehouse.Note).
		Scan(&warehouse.UpdatedAt)
}

// Delete deletes a warehouse
func (r *warehouseRepository) Delete(id int) error {
	query := `DELETE FROM warehouses WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// List gets all warehouses
func (r *warehouseRepository) List() ([]*models.Warehouse, error) {
	query := `
		SELECT id, name, location, note, created_at, updated_at
		FROM warehouses ORDER BY name`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []*models.Warehouse
	for rows.Next() {
		warehouse := &models.Warehouse{}
		err := rows.Scan(
			&warehouse.ID, &warehouse.Name, &warehouse.Location,
			&warehouse.Note, &warehouse.CreatedAt, &warehouse.UpdatedAt)
		if err != nil {
			return nil, err
		}
		warehouses = append(warehouses, warehouse)
	}

	return warehouses, nil
}
