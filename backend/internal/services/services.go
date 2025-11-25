package services

import (
	"wms-backend/internal/repositories"
)

// Services holds all service interfaces
type Services struct {
	Auth            AuthService
	User            UserService
	Product         ProductService
	ProductPrice    ProductPriceService
	CustomerGroup   CustomerGroupService
	Customer        CustomerService
	Warehouse       WarehouseService
	Stock           StockService
	StockMovement   StockMovementService
	SalesOrder      SalesOrderService
	DeliveryOrder   DeliveryOrderService
	PurchaseOrder   PurchaseOrderService
	GoodsReceipt    GoodsReceiptService
	Transfer        TransferService
	StockAdjustment StockAdjustmentService
	Order           OrderService
}

// NewServices creates a new services instance
func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		Auth:            NewAuthService(repos.User),
		User:            NewUserService(repos.User),
		Product:         NewProductService(repos.Product),
		ProductPrice:    NewProductPriceService(repos.ProductPrice),
		CustomerGroup:   NewCustomerGroupService(repos.CustomerGroup),
		Customer:        NewCustomerService(repos.Customer),
		Warehouse:       NewWarehouseService(repos.Warehouse),
		Stock:           NewStockService(repos.Stock, repos.StockMovement, repos.Product),
		StockMovement:   NewStockMovementService(repos.StockMovement),
		SalesOrder:      NewSalesOrderService(repos.SalesOrder),
		DeliveryOrder:   NewDeliveryOrderService(repos.DeliveryOrder),
		PurchaseOrder:   NewPurchaseOrderService(repos.PurchaseOrder),
		GoodsReceipt:    NewGoodsReceiptService(repos.GoodsReceipt),
		Transfer:        NewTransferService(repos.Transfer),
		StockAdjustment: NewStockAdjustmentService(repos.StockAdjustment),
		Order:           NewOrderService(repos.Order),
	}
}
