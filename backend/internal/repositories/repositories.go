package repositories

import (
	"database/sql"
)

// Repositories holds all repository interfaces
type Repositories struct {
	User            UserRepository
	Product         ProductRepository
	ProductPrice    ProductPriceRepository
	CustomerGroup   CustomerGroupRepository
	Customer        CustomerRepository
	Warehouse       WarehouseRepository
	Stock           StockRepository
	StockMovement   StockMovementRepository
	SalesOrder      SalesOrderRepository
	DeliveryOrder   DeliveryOrderRepository
	PurchaseOrder   PurchaseOrderRepository
	GoodsReceipt    GoodsReceiptRepository
	Transfer        TransferRepository
	StockAdjustment StockAdjustmentRepository
	Order           OrderRepository
}

// NewRepositories creates a new repositories instance
func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		User:            NewUserRepository(db),
		Product:         NewProductRepository(db),
		ProductPrice:    NewProductPriceRepository(db),
		CustomerGroup:   NewCustomerGroupRepository(db),
		Customer:        NewCustomerRepository(db),
		Warehouse:       NewWarehouseRepository(db),
		Stock:           NewStockRepository(db),
		StockMovement:   NewStockMovementRepository(db),
		SalesOrder:      NewSalesOrderRepository(db),
		DeliveryOrder:   NewDeliveryOrderRepository(db),
		PurchaseOrder:   NewPurchaseOrderRepository(db),
		GoodsReceipt:    NewGoodsReceiptRepository(db),
		Transfer:        NewTransferRepository(db),
		StockAdjustment: NewStockAdjustmentRepository(db),
		Order:           NewOrderRepository(db),
	}
}
