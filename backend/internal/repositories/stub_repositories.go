package repositories

import (
	"database/sql"
	"wms-backend/internal/models"
)

// ProductPriceRepository interface
type ProductPriceRepository interface {
	Create(price *models.ProductPrice) error
	GetByProductID(productID int) ([]*models.ProductPrice, error)
	Update(price *models.ProductPrice) error
	Delete(id int) error
}

type productPriceRepository struct {
	db *sql.DB
}

func NewProductPriceRepository(db *sql.DB) ProductPriceRepository {
	return &productPriceRepository{db: db}
}

func (r *productPriceRepository) Create(price *models.ProductPrice) error {
	query := `INSERT INTO product_prices (product_id, customer_group_id, unit, price) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, price.ProductID, price.CustomerGroupID, price.Unit, price.Price).
		Scan(&price.ID, &price.CreatedAt, &price.UpdatedAt)
}

func (r *productPriceRepository) GetByProductID(productID int) ([]*models.ProductPrice, error) {
	query := `SELECT id, product_id, customer_group_id, unit, price, created_at, updated_at FROM product_prices WHERE product_id = $1`
	rows, err := r.db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prices []*models.ProductPrice
	for rows.Next() {
		price := &models.ProductPrice{}
		err := rows.Scan(&price.ID, &price.ProductID, &price.CustomerGroupID, &price.Unit, &price.Price, &price.CreatedAt, &price.UpdatedAt)
		if err != nil {
			return nil, err
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (r *productPriceRepository) Update(price *models.ProductPrice) error {
	query := `UPDATE product_prices SET unit = $2, price = $3 WHERE id = $1 RETURNING updated_at`
	return r.db.QueryRow(query, price.ID, price.Unit, price.Price).Scan(&price.UpdatedAt)
}

func (r *productPriceRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM product_prices WHERE id = $1`, id)
	return err
}

// CustomerGroupRepository interface
type CustomerGroupRepository interface {
	Create(group *models.CustomerGroup) error
	GetByID(id int) (*models.CustomerGroup, error)
	List() ([]*models.CustomerGroup, error)
	Update(group *models.CustomerGroup) error
	Delete(id int) error
}

type customerGroupRepository struct {
	db *sql.DB
}

func NewCustomerGroupRepository(db *sql.DB) CustomerGroupRepository {
	return &customerGroupRepository{db: db}
}

func (r *customerGroupRepository) Create(group *models.CustomerGroup) error {
	query := `INSERT INTO customer_groups (name, description) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, group.Name, group.Description).
		Scan(&group.ID, &group.CreatedAt, &group.UpdatedAt)
}

func (r *customerGroupRepository) GetByID(id int) (*models.CustomerGroup, error) {
	group := &models.CustomerGroup{}
	query := `SELECT id, name, description, created_at, updated_at FROM customer_groups WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&group.ID, &group.Name, &group.Description, &group.CreatedAt, &group.UpdatedAt)
	return group, err
}

func (r *customerGroupRepository) List() ([]*models.CustomerGroup, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM customer_groups ORDER BY name`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []*models.CustomerGroup
	for rows.Next() {
		group := &models.CustomerGroup{}
		err := rows.Scan(&group.ID, &group.Name, &group.Description, &group.CreatedAt, &group.UpdatedAt)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (r *customerGroupRepository) Update(group *models.CustomerGroup) error {
	query := `UPDATE customer_groups SET name = $2, description = $3 WHERE id = $1 RETURNING updated_at`
	return r.db.QueryRow(query, group.ID, group.Name, group.Description).Scan(&group.UpdatedAt)
}

func (r *customerGroupRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM customer_groups WHERE id = $1`, id)
	return err
}

// Stub repositories for documents
type SalesOrderRepository interface {
	Create(order *models.SalesOrder) error
	GetByID(id int) (*models.SalesOrder, error)
	List() ([]*models.SalesOrder, error)
	Update(order *models.SalesOrder) error
	Delete(id int) error
}

type salesOrderRepository struct{ db *sql.DB }

func NewSalesOrderRepository(db *sql.DB) SalesOrderRepository              { return &salesOrderRepository{db: db} }
func (r *salesOrderRepository) Create(order *models.SalesOrder) error      { return nil }
func (r *salesOrderRepository) GetByID(id int) (*models.SalesOrder, error) { return nil, nil }
func (r *salesOrderRepository) List() ([]*models.SalesOrder, error)        { return nil, nil }
func (r *salesOrderRepository) Update(order *models.SalesOrder) error      { return nil }
func (r *salesOrderRepository) Delete(id int) error                        { return nil }

type DeliveryOrderRepository interface {
	Create(order *models.DeliveryOrder) error
	GetByID(id int) (*models.DeliveryOrder, error)
	List() ([]*models.DeliveryOrder, error)
	Update(order *models.DeliveryOrder) error
	Delete(id int) error
}

type deliveryOrderRepository struct{ db *sql.DB }

func NewDeliveryOrderRepository(db *sql.DB) DeliveryOrderRepository {
	return &deliveryOrderRepository{db: db}
}
func (r *deliveryOrderRepository) Create(order *models.DeliveryOrder) error      { return nil }
func (r *deliveryOrderRepository) GetByID(id int) (*models.DeliveryOrder, error) { return nil, nil }
func (r *deliveryOrderRepository) List() ([]*models.DeliveryOrder, error)        { return nil, nil }
func (r *deliveryOrderRepository) Update(order *models.DeliveryOrder) error      { return nil }
func (r *deliveryOrderRepository) Delete(id int) error                           { return nil }

type PurchaseOrderRepository interface {
	Create(order *models.PurchaseOrder) error
	GetByID(id int) (*models.PurchaseOrder, error)
	List() ([]*models.PurchaseOrder, error)
	Update(order *models.PurchaseOrder) error
	Delete(id int) error
}

type purchaseOrderRepository struct{ db *sql.DB }

func NewPurchaseOrderRepository(db *sql.DB) PurchaseOrderRepository {
	return &purchaseOrderRepository{db: db}
}
func (r *purchaseOrderRepository) Create(order *models.PurchaseOrder) error      { return nil }
func (r *purchaseOrderRepository) GetByID(id int) (*models.PurchaseOrder, error) { return nil, nil }
func (r *purchaseOrderRepository) List() ([]*models.PurchaseOrder, error)        { return nil, nil }
func (r *purchaseOrderRepository) Update(order *models.PurchaseOrder) error      { return nil }
func (r *purchaseOrderRepository) Delete(id int) error                           { return nil }

type GoodsReceiptRepository interface {
	Create(receipt *models.GoodsReceipt) error
	GetByID(id int) (*models.GoodsReceipt, error)
	List() ([]*models.GoodsReceipt, error)
	Update(receipt *models.GoodsReceipt) error
	Delete(id int) error
}

type goodsReceiptRepository struct{ db *sql.DB }

func NewGoodsReceiptRepository(db *sql.DB) GoodsReceiptRepository {
	return &goodsReceiptRepository{db: db}
}
func (r *goodsReceiptRepository) Create(receipt *models.GoodsReceipt) error    { return nil }
func (r *goodsReceiptRepository) GetByID(id int) (*models.GoodsReceipt, error) { return nil, nil }
func (r *goodsReceiptRepository) List() ([]*models.GoodsReceipt, error)        { return nil, nil }
func (r *goodsReceiptRepository) Update(receipt *models.GoodsReceipt) error    { return nil }
func (r *goodsReceiptRepository) Delete(id int) error                          { return nil }

type TransferRepository interface {
	Create(transfer *models.Transfer) error
	GetByID(id int) (*models.Transfer, error)
	List() ([]*models.Transfer, error)
	Update(transfer *models.Transfer) error
	Delete(id int) error
}

type transferRepository struct{ db *sql.DB }

func NewTransferRepository(db *sql.DB) TransferRepository              { return &transferRepository{db: db} }
func (r *transferRepository) Create(transfer *models.Transfer) error   { return nil }
func (r *transferRepository) GetByID(id int) (*models.Transfer, error) { return nil, nil }
func (r *transferRepository) List() ([]*models.Transfer, error)        { return nil, nil }
func (r *transferRepository) Update(transfer *models.Transfer) error   { return nil }
func (r *transferRepository) Delete(id int) error                      { return nil }

type StockAdjustmentRepository interface {
	Create(adjustment *models.StockAdjustment) error
	GetByID(id int) (*models.StockAdjustment, error)
	List() ([]*models.StockAdjustment, error)
	Update(adjustment *models.StockAdjustment) error
	Delete(id int) error
}

type stockAdjustmentRepository struct{ db *sql.DB }

func NewStockAdjustmentRepository(db *sql.DB) StockAdjustmentRepository {
	return &stockAdjustmentRepository{db: db}
}
func (r *stockAdjustmentRepository) Create(adjustment *models.StockAdjustment) error { return nil }
func (r *stockAdjustmentRepository) GetByID(id int) (*models.StockAdjustment, error) { return nil, nil }
func (r *stockAdjustmentRepository) List() ([]*models.StockAdjustment, error)        { return nil, nil }
func (r *stockAdjustmentRepository) Update(adjustment *models.StockAdjustment) error { return nil }
func (r *stockAdjustmentRepository) Delete(id int) error                             { return nil }
