package repositories

import (
	"database/sql"
	"time"
	"wms-backend/internal/models"
)

// StockMovementRepository interface defines stock movement repository methods
type StockMovementRepository interface {
	Create(movement *models.StockMovement) error
	GetByID(id int) (*models.StockMovement, error)
	List(productID, warehouseID int, limit, offset int) ([]*models.StockMovement, error)
	GetByDateRange(productID, warehouseID int, startDate, endDate time.Time) ([]*models.StockMovement, error)
	GetByReference(refType string, refID int) ([]*models.StockMovement, error)
}

// stockMovementRepository implements StockMovementRepository
type stockMovementRepository struct {
	db *sql.DB
}

// NewStockMovementRepository creates a new stock movement repository
func NewStockMovementRepository(db *sql.DB) StockMovementRepository {
	return &stockMovementRepository{db: db}
}

// Create creates a new stock movement
func (r *stockMovementRepository) Create(movement *models.StockMovement) error {
	query := `
		INSERT INTO stock_movements (product_id, warehouse_id, type, qty, unit, ref_id, ref_type, date, note, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, created_at`

	return r.db.QueryRow(query,
		movement.ProductID, movement.WarehouseID, movement.Type, movement.Qty,
		movement.Unit, movement.RefID, movement.RefType, movement.Date,
		movement.Note, movement.CreatedBy).
		Scan(&movement.ID, &movement.CreatedAt)
}

// GetByID gets a stock movement by ID
func (r *stockMovementRepository) GetByID(id int) (*models.StockMovement, error) {
	movement := &models.StockMovement{}
	query := `
		SELECT id, product_id, warehouse_id, type, qty, unit, ref_id, ref_type, date, note, created_by, created_at
		FROM stock_movements WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&movement.ID, &movement.ProductID, &movement.WarehouseID, &movement.Type,
		&movement.Qty, &movement.Unit, &movement.RefID, &movement.RefType,
		&movement.Date, &movement.Note, &movement.CreatedBy, &movement.CreatedAt)

	if err != nil {
		return nil, err
	}
	return movement, nil
}

// List gets stock movements with optional filters
func (r *stockMovementRepository) List(productID, warehouseID int, limit, offset int) ([]*models.StockMovement, error) {
	query := `
		SELECT id, product_id, warehouse_id, type, qty, unit, ref_id, ref_type, date, note, created_by, created_at
		FROM stock_movements
		WHERE ($1 = 0 OR product_id = $1) 
		  AND ($2 = 0 OR warehouse_id = $2)
		ORDER BY date DESC, created_at DESC
		LIMIT $3 OFFSET $4`

	rows, err := r.db.Query(query, productID, warehouseID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movements []*models.StockMovement
	for rows.Next() {
		movement := &models.StockMovement{}
		err := rows.Scan(
			&movement.ID, &movement.ProductID, &movement.WarehouseID, &movement.Type,
			&movement.Qty, &movement.Unit, &movement.RefID, &movement.RefType,
			&movement.Date, &movement.Note, &movement.CreatedBy, &movement.CreatedAt)
		if err != nil {
			return nil, err
		}
		movements = append(movements, movement)
	}

	return movements, nil
}

// GetByDateRange gets stock movements within a date range
func (r *stockMovementRepository) GetByDateRange(productID, warehouseID int, startDate, endDate time.Time) ([]*models.StockMovement, error) {
	query := `
		SELECT id, product_id, warehouse_id, type, qty, unit, ref_id, ref_type, date, note, created_by, created_at
		FROM stock_movements
		WHERE ($1 = 0 OR product_id = $1) 
		  AND ($2 = 0 OR warehouse_id = $2)
		  AND date >= $3 AND date <= $4
		ORDER BY date DESC, created_at DESC`

	rows, err := r.db.Query(query, productID, warehouseID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movements []*models.StockMovement
	for rows.Next() {
		movement := &models.StockMovement{}
		err := rows.Scan(
			&movement.ID, &movement.ProductID, &movement.WarehouseID, &movement.Type,
			&movement.Qty, &movement.Unit, &movement.RefID, &movement.RefType,
			&movement.Date, &movement.Note, &movement.CreatedBy, &movement.CreatedAt)
		if err != nil {
			return nil, err
		}
		movements = append(movements, movement)
	}

	return movements, nil
}

// GetByReference gets stock movements by reference document
func (r *stockMovementRepository) GetByReference(refType string, refID int) ([]*models.StockMovement, error) {
	query := `
		SELECT id, product_id, warehouse_id, type, qty, unit, ref_id, ref_type, date, note, created_by, created_at
		FROM stock_movements
		WHERE ref_type = $1 AND ref_id = $2
		ORDER BY date DESC, created_at DESC`

	rows, err := r.db.Query(query, refType, refID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movements []*models.StockMovement
	for rows.Next() {
		movement := &models.StockMovement{}
		err := rows.Scan(
			&movement.ID, &movement.ProductID, &movement.WarehouseID, &movement.Type,
			&movement.Qty, &movement.Unit, &movement.RefID, &movement.RefType,
			&movement.Date, &movement.Note, &movement.CreatedBy, &movement.CreatedAt)
		if err != nil {
			return nil, err
		}
		movements = append(movements, movement)
	}

	return movements, nil
}
