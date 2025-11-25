package handlers

import (
	"wms-backend/internal/services"
)

// Handlers holds all handler interfaces
type Handlers struct {
	Auth            *AuthHandler
	User            *UserHandler
	Product         *ProductHandler
	ProductPrice    *ProductPriceHandler
	CustomerGroup   *CustomerGroupHandler
	Customer        *CustomerHandler
	Warehouse       *WarehouseHandler
	Stock           *StockHandler
	StockMovement   *StockMovementHandler
	SalesOrder      *SalesOrderHandler
	DeliveryOrder   *DeliveryOrderHandler
	PurchaseOrder   *PurchaseOrderHandler
	GoodsReceipt    *GoodsReceiptHandler
	Transfer        *TransferHandler
	StockAdjustment *StockAdjustmentHandler
	Order           *OrderHandler
}

// NewHandlers creates a new handlers instance
func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		Auth:            NewAuthHandler(services.Auth),
		User:            NewUserHandler(services.User),
		Product:         NewProductHandler(services.Product),
		ProductPrice:    NewProductPriceHandler(services.ProductPrice),
		CustomerGroup:   NewCustomerGroupHandler(services.CustomerGroup),
		Customer:        NewCustomerHandler(services.Customer),
		Warehouse:       NewWarehouseHandler(services.Warehouse),
		Stock:           NewStockHandler(services.Stock),
		StockMovement:   NewStockMovementHandler(services.StockMovement),
		SalesOrder:      NewSalesOrderHandler(services.SalesOrder),
		DeliveryOrder:   NewDeliveryOrderHandler(services.DeliveryOrder),
		PurchaseOrder:   NewPurchaseOrderHandler(services.PurchaseOrder),
		GoodsReceipt:    NewGoodsReceiptHandler(services.GoodsReceipt),
		Transfer:        NewTransferHandler(services.Transfer),
		StockAdjustment: NewStockAdjustmentHandler(services.StockAdjustment),
		Order:           NewOrderHandler(services.Order),
	}
}
