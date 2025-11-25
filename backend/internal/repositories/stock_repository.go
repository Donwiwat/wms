package repositories

import (
	"database/sql"
	"wms-backend/internal/models"
)

// StockRepository interface defines stock repository methods
type StockRepository interface {
	GetStock(productID, warehouseID int) (*models.Stock, error)
	GetStockSummary(productID, warehouseID int) ([]*models.StockSummary, error)
	GetStockCard(productID, warehouseID int) ([]*models.StockCardEntry, error)
	UpdateStock(stock *models.Stock) error
	CreateOrUpdateStock(productID, warehouseID int, remain1, remain2 int) error
	GetStockByProduct(productID int) ([]*models.Stock, error)
	GetStockByWarehouse(warehouseID int) ([]*models.Stock, error)
}

// stockRepository implements StockRepository
type stockRepository struct {
	db *sql.DB
}

// NewStockRepository creates a new stock repository
func NewStockRepository(db *sql.DB) StockRepository {
	return &stockRepository{db: db}
}

// GetStock gets stock for a specific product and warehouse
func (r *stockRepository) GetStock(productID, warehouseID int) (*models.Stock, error) {
	stock := &models.Stock{}
	query := `
		SELECT id, product_id, warehouse_id, remain1, remain2, total_remain, updated_at
		FROM stock 
		WHERE product_id = $1 AND warehouse_id = $2`

	err := r.db.QueryRow(query, productID, warehouseID).Scan(
		&stock.ID, &stock.ProductID, &stock.WarehouseID,
		&stock.Remain1, &stock.Remain2, &stock.TotalRemain, &stock.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return stock, nil
}

// GetStockSummary gets stock summary with product and warehouse details
func (r *stockRepository) GetStockSummary(productID, warehouseID int) ([]*models.StockSummary, error) {
	query := `
		SELECT 
			s.product_id, p.name as product_name,
			s.warehouse_id, w.name as warehouse_name,
			s.remain1, s.remain2, s.total_remain,
			p.unit1, p.unit2, s.updated_at
		FROM stock s
		JOIN products p ON s.product_id = p.id
		JOIN warehouses w ON s.warehouse_id = w.id
		WHERE ($1 = 0 OR s.product_id = $1) 
		  AND ($2 = 0 OR s.warehouse_id = $2)
		ORDER BY p.name, w.name`

	rows, err := r.db.Query(query, productID, warehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*models.StockSummary
	for rows.Next() {
		summary := &models.StockSummary{}
		err := rows.Scan(
			&summary.ProductID, &summary.ProductName,
			&summary.WarehouseID, &summary.WarehouseName,
			&summary.Remain1, &summary.Remain2, &summary.TotalRemain,
			&summary.Unit1, &summary.Unit2, &summary.UpdatedAt)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, summary)
	}

	return summaries, nil
}

// GetStockCard gets stock movement history for a product and warehouse
func (r *stockRepository) GetStockCard(productID, warehouseID int) ([]*models.StockCardEntry, error) {
	query := `
		SELECT date, type, qty, unit, ref_type, ref_id, note, created_by
		FROM stock_movements
		WHERE product_id = $1 AND warehouse_id = $2
		ORDER BY date DESC, created_at DESC`

	rows, err := r.db.Query(query, productID, warehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []*models.StockCardEntry
	for rows.Next() {
		entry := &models.StockCardEntry{}
		err := rows.Scan(
			&entry.Date, &entry.Type, &entry.Qty, &entry.Unit,
			&entry.RefType, &entry.RefID, &entry.Note, &entry.CreatedBy)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// UpdateStock updates stock levels
func (r *stockRepository) UpdateStock(stock *models.Stock) error {
	query := `
		UPDATE stock 
		SET remain1 = $3, remain2 = $4
		WHERE product_id = $1 AND warehouse_id = $2
		RETURNING total_remain, updated_at`

	return r.db.QueryRow(query, stock.ProductID, stock.WarehouseID, stock.Remain1, stock.Remain2).
		Scan(&stock.TotalRemain, &stock.UpdatedAt)
}

// CreateOrUpdateStock creates or updates stock record
func (r *stockRepository) CreateOrUpdateStock(productID, warehouseID int, remain1, remain2 int) error {
	query := `
		INSERT INTO stock (product_id, warehouse_id, remain1, remain2)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (product_id, warehouse_id)
		DO UPDATE SET remain1 = $3, remain2 = $4`

	_, err := r.db.Exec(query, productID, warehouseID, remain1, remain2)
	return err
}

// GetStockByProduct gets all stock records for a product
func (r *stockRepository) GetStockByProduct(productID int) ([]*models.Stock, error) {
	query := `
		SELECT id, product_id, warehouse_id, remain1, remain2, total_remain, updated_at
		FROM stock 
		WHERE product_id = $1
		ORDER BY warehouse_id`

	rows, err := r.db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []*models.Stock
	for rows.Next() {
		stock := &models.Stock{}
		err := rows.Scan(
			&stock.ID, &stock.ProductID, &stock.WarehouseID,
			&stock.Remain1, &stock.Remain2, &stock.TotalRemain, &stock.UpdatedAt)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}

// GetStockByWarehouse gets all stock records for a warehouse
func (r *stockRepository) GetStockByWarehouse(warehouseID int) ([]*models.Stock, error) {
	query := `
		SELECT id, product_id, warehouse_id, remain1, remain2, total_remain, updated_at
		FROM stock 
		WHERE warehouse_id = $1
		ORDER BY product_id`

	rows, err := r.db.Query(query, warehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []*models.Stock
	for rows.Next() {
		stock := &models.Stock{}
		err := rows.Scan(
			&stock.ID, &stock.ProductID, &stock.WarehouseID,
			&stock.Remain1, &stock.Remain2, &stock.TotalRemain, &stock.UpdatedAt)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}
