package services

import (
	"wms-backend/internal/models"
	"wms-backend/internal/repositories"
)

// UserService interface
type UserService interface {
	GetByID(id int) (*models.User, error)
	List() ([]*models.User, error)
	Update(user *models.User) error
	Delete(id int) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetByID(id int) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) List() ([]*models.User, error) {
	return s.userRepo.List()
}

func (s *userService) Update(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) Delete(id int) error {
	return s.userRepo.Delete(id)
}

// ProductService interface
type ProductService interface {
	Create(product *models.Product) error
	GetByID(id int) (*models.Product, error)
	List() ([]*models.Product, error)
	Update(product *models.Product) error
	Delete(id int) error
	Search(query string) ([]*models.Product, error)
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) Create(product *models.Product) error {
	return s.productRepo.Create(product)
}

func (s *productService) GetByID(id int) (*models.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *productService) List() ([]*models.Product, error) {
	return s.productRepo.List()
}

func (s *productService) Update(product *models.Product) error {
	return s.productRepo.Update(product)
}

func (s *productService) Delete(id int) error {
	return s.productRepo.Delete(id)
}

func (s *productService) Search(query string) ([]*models.Product, error) {
	return s.productRepo.Search(query)
}

// ProductPriceService interface
type ProductPriceService interface {
	Create(price *models.ProductPrice) error
	GetByProductID(productID int) ([]*models.ProductPrice, error)
	Update(price *models.ProductPrice) error
	Delete(id int) error
}

type productPriceService struct {
	priceRepo repositories.ProductPriceRepository
}

func NewProductPriceService(priceRepo repositories.ProductPriceRepository) ProductPriceService {
	return &productPriceService{priceRepo: priceRepo}
}

func (s *productPriceService) Create(price *models.ProductPrice) error {
	return s.priceRepo.Create(price)
}

func (s *productPriceService) GetByProductID(productID int) ([]*models.ProductPrice, error) {
	return s.priceRepo.GetByProductID(productID)
}

func (s *productPriceService) Update(price *models.ProductPrice) error {
	return s.priceRepo.Update(price)
}

func (s *productPriceService) Delete(id int) error {
	return s.priceRepo.Delete(id)
}

// CustomerGroupService interface
type CustomerGroupService interface {
	Create(group *models.CustomerGroup) error
	GetByID(id int) (*models.CustomerGroup, error)
	List() ([]*models.CustomerGroup, error)
	Update(group *models.CustomerGroup) error
	Delete(id int) error
}

type customerGroupService struct {
	groupRepo repositories.CustomerGroupRepository
}

func NewCustomerGroupService(groupRepo repositories.CustomerGroupRepository) CustomerGroupService {
	return &customerGroupService{groupRepo: groupRepo}
}

func (s *customerGroupService) Create(group *models.CustomerGroup) error {
	return s.groupRepo.Create(group)
}

func (s *customerGroupService) GetByID(id int) (*models.CustomerGroup, error) {
	return s.groupRepo.GetByID(id)
}

func (s *customerGroupService) List() ([]*models.CustomerGroup, error) {
	return s.groupRepo.List()
}

func (s *customerGroupService) Update(group *models.CustomerGroup) error {
	return s.groupRepo.Update(group)
}

func (s *customerGroupService) Delete(id int) error {
	return s.groupRepo.Delete(id)
}

// WarehouseService interface
type WarehouseService interface {
	Create(warehouse *models.Warehouse) error
	GetByID(id int) (*models.Warehouse, error)
	List() ([]*models.Warehouse, error)
	Update(warehouse *models.Warehouse) error
	Delete(id int) error
}

type warehouseService struct {
	warehouseRepo repositories.WarehouseRepository
}

func NewWarehouseService(warehouseRepo repositories.WarehouseRepository) WarehouseService {
	return &warehouseService{warehouseRepo: warehouseRepo}
}

func (s *warehouseService) Create(warehouse *models.Warehouse) error {
	return s.warehouseRepo.Create(warehouse)
}

func (s *warehouseService) GetByID(id int) (*models.Warehouse, error) {
	return s.warehouseRepo.GetByID(id)
}

func (s *warehouseService) List() ([]*models.Warehouse, error) {
	return s.warehouseRepo.List()
}

func (s *warehouseService) Update(warehouse *models.Warehouse) error {
	return s.warehouseRepo.Update(warehouse)
}

func (s *warehouseService) Delete(id int) error {
	return s.warehouseRepo.Delete(id)
}

// StockMovementService interface
type StockMovementService interface {
	GetByID(id int) (*models.StockMovement, error)
	List(productID, warehouseID int, limit, offset int) ([]*models.StockMovement, error)
}

type stockMovementService struct {
	movementRepo repositories.StockMovementRepository
}

func NewStockMovementService(movementRepo repositories.StockMovementRepository) StockMovementService {
	return &stockMovementService{movementRepo: movementRepo}
}

func (s *stockMovementService) GetByID(id int) (*models.StockMovement, error) {
	return s.movementRepo.GetByID(id)
}

func (s *stockMovementService) List(productID, warehouseID int, limit, offset int) ([]*models.StockMovement, error) {
	return s.movementRepo.List(productID, warehouseID, limit, offset)
}

// Stub services for documents
type SalesOrderService interface {
	Create(order *models.SalesOrder) error
	GetByID(id int) (*models.SalesOrder, error)
	List() ([]*models.SalesOrder, error)
	Update(order *models.SalesOrder) error
	Delete(id int) error
}

type salesOrderService struct {
	repo repositories.SalesOrderRepository
}

func NewSalesOrderService(repo repositories.SalesOrderRepository) SalesOrderService {
	return &salesOrderService{repo: repo}
}
func (s *salesOrderService) Create(order *models.SalesOrder) error      { return s.repo.Create(order) }
func (s *salesOrderService) GetByID(id int) (*models.SalesOrder, error) { return s.repo.GetByID(id) }
func (s *salesOrderService) List() ([]*models.SalesOrder, error)        { return s.repo.List() }
func (s *salesOrderService) Update(order *models.SalesOrder) error      { return s.repo.Update(order) }
func (s *salesOrderService) Delete(id int) error                        { return s.repo.Delete(id) }

type DeliveryOrderService interface {
	Create(order *models.DeliveryOrder) error
	GetByID(id int) (*models.DeliveryOrder, error)
	List() ([]*models.DeliveryOrder, error)
	Update(order *models.DeliveryOrder) error
	Delete(id int) error
}

type deliveryOrderService struct {
	repo repositories.DeliveryOrderRepository
}

func NewDeliveryOrderService(repo repositories.DeliveryOrderRepository) DeliveryOrderService {
	return &deliveryOrderService{repo: repo}
}
func (s *deliveryOrderService) Create(order *models.DeliveryOrder) error { return s.repo.Create(order) }
func (s *deliveryOrderService) GetByID(id int) (*models.DeliveryOrder, error) {
	return s.repo.GetByID(id)
}
func (s *deliveryOrderService) List() ([]*models.DeliveryOrder, error)   { return s.repo.List() }
func (s *deliveryOrderService) Update(order *models.DeliveryOrder) error { return s.repo.Update(order) }
func (s *deliveryOrderService) Delete(id int) error                      { return s.repo.Delete(id) }

type PurchaseOrderService interface {
	Create(order *models.PurchaseOrder) error
	GetByID(id int) (*models.PurchaseOrder, error)
	List() ([]*models.PurchaseOrder, error)
	Update(order *models.PurchaseOrder) error
	Delete(id int) error
}

type purchaseOrderService struct {
	repo repositories.PurchaseOrderRepository
}

func NewPurchaseOrderService(repo repositories.PurchaseOrderRepository) PurchaseOrderService {
	return &purchaseOrderService{repo: repo}
}
func (s *purchaseOrderService) Create(order *models.PurchaseOrder) error { return s.repo.Create(order) }
func (s *purchaseOrderService) GetByID(id int) (*models.PurchaseOrder, error) {
	return s.repo.GetByID(id)
}
func (s *purchaseOrderService) List() ([]*models.PurchaseOrder, error)   { return s.repo.List() }
func (s *purchaseOrderService) Update(order *models.PurchaseOrder) error { return s.repo.Update(order) }
func (s *purchaseOrderService) Delete(id int) error                      { return s.repo.Delete(id) }

type GoodsReceiptService interface {
	Create(receipt *models.GoodsReceipt) error
	GetByID(id int) (*models.GoodsReceipt, error)
	List() ([]*models.GoodsReceipt, error)
	Update(receipt *models.GoodsReceipt) error
	Delete(id int) error
}

type goodsReceiptService struct {
	repo repositories.GoodsReceiptRepository
}

func NewGoodsReceiptService(repo repositories.GoodsReceiptRepository) GoodsReceiptService {
	return &goodsReceiptService{repo: repo}
}
func (s *goodsReceiptService) Create(receipt *models.GoodsReceipt) error {
	return s.repo.Create(receipt)
}
func (s *goodsReceiptService) GetByID(id int) (*models.GoodsReceipt, error) {
	return s.repo.GetByID(id)
}
func (s *goodsReceiptService) List() ([]*models.GoodsReceipt, error) { return s.repo.List() }
func (s *goodsReceiptService) Update(receipt *models.GoodsReceipt) error {
	return s.repo.Update(receipt)
}
func (s *goodsReceiptService) Delete(id int) error { return s.repo.Delete(id) }

type TransferService interface {
	Create(transfer *models.Transfer) error
	GetByID(id int) (*models.Transfer, error)
	List() ([]*models.Transfer, error)
	Update(transfer *models.Transfer) error
	Delete(id int) error
}

type transferService struct {
	repo repositories.TransferRepository
}

func NewTransferService(repo repositories.TransferRepository) TransferService {
	return &transferService{repo: repo}
}
func (s *transferService) Create(transfer *models.Transfer) error   { return s.repo.Create(transfer) }
func (s *transferService) GetByID(id int) (*models.Transfer, error) { return s.repo.GetByID(id) }
func (s *transferService) List() ([]*models.Transfer, error)        { return s.repo.List() }
func (s *transferService) Update(transfer *models.Transfer) error   { return s.repo.Update(transfer) }
func (s *transferService) Delete(id int) error                      { return s.repo.Delete(id) }

type StockAdjustmentService interface {
	Create(adjustment *models.StockAdjustment) error
	GetByID(id int) (*models.StockAdjustment, error)
	List() ([]*models.StockAdjustment, error)
	Update(adjustment *models.StockAdjustment) error
	Delete(id int) error
}

type stockAdjustmentService struct {
	repo repositories.StockAdjustmentRepository
}

func NewStockAdjustmentService(repo repositories.StockAdjustmentRepository) StockAdjustmentService {
	return &stockAdjustmentService{repo: repo}
}
func (s *stockAdjustmentService) Create(adjustment *models.StockAdjustment) error {
	return s.repo.Create(adjustment)
}
func (s *stockAdjustmentService) GetByID(id int) (*models.StockAdjustment, error) {
	return s.repo.GetByID(id)
}
func (s *stockAdjustmentService) List() ([]*models.StockAdjustment, error) { return s.repo.List() }
func (s *stockAdjustmentService) Update(adjustment *models.StockAdjustment) error {
	return s.repo.Update(adjustment)
}
func (s *stockAdjustmentService) Delete(id int) error { return s.repo.Delete(id) }
