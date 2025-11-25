package handlers

import (
	"net/http"
	"strconv"

	"wms-backend/internal/models"
	"wms-backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UserHandler handles user requests
type UserHandler struct {
	userService services.UserService
	validator   *validator.Validate
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService, validator: validator.New()}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// ProductHandler handles product requests
type ProductHandler struct {
	productService services.ProductService
	validator      *validator.Validate
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService, validator: validator.New()}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.productService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.productService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.productService.Create(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = id
	if err := h.productService.Update(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.productService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func (h *ProductHandler) GetProductPrices(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"product_id": id, "prices": []interface{}{}})
}

// ProductPriceHandler handles product price requests
type ProductPriceHandler struct {
	service   services.ProductPriceService
	validator *validator.Validate
}

func NewProductPriceHandler(service services.ProductPriceService) *ProductPriceHandler {
	return &ProductPriceHandler{service: service, validator: validator.New()}
}

func (h *ProductPriceHandler) CreateProductPrice(c *gin.Context) {
	var price models.ProductPrice
	if err := c.ShouldBindJSON(&price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&price); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, price)
}

func (h *ProductPriceHandler) UpdateProductPrice(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var price models.ProductPrice
	if err := c.ShouldBindJSON(&price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	price.ID = id
	if err := h.service.Update(&price); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, price)
}

func (h *ProductPriceHandler) DeleteProductPrice(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Price deleted"})
}

// CustomerGroupHandler handles customer group requests
type CustomerGroupHandler struct {
	service   services.CustomerGroupService
	validator *validator.Validate
}

func NewCustomerGroupHandler(service services.CustomerGroupService) *CustomerGroupHandler {
	return &CustomerGroupHandler{service: service, validator: validator.New()}
}

func (h *CustomerGroupHandler) GetCustomerGroups(c *gin.Context) {
	groups, err := h.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (h *CustomerGroupHandler) CreateCustomerGroup(c *gin.Context) {
	var group models.CustomerGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, group)
}

func (h *CustomerGroupHandler) UpdateCustomerGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var group models.CustomerGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	group.ID = id
	if err := h.service.Update(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *CustomerGroupHandler) DeleteCustomerGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer group deleted"})
}

// WarehouseHandler handles warehouse requests
type WarehouseHandler struct {
	service   services.WarehouseService
	validator *validator.Validate
}

func NewWarehouseHandler(service services.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{service: service, validator: validator.New()}
}

func (h *WarehouseHandler) GetWarehouses(c *gin.Context) {
	warehouses, err := h.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, warehouses)
}

func (h *WarehouseHandler) GetWarehouse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	warehouse, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warehouse not found"})
		return
	}
	c.JSON(http.StatusOK, warehouse)
}

func (h *WarehouseHandler) CreateWarehouse(c *gin.Context) {
	var warehouse models.Warehouse
	if err := c.ShouldBindJSON(&warehouse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&warehouse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, warehouse)
}

func (h *WarehouseHandler) UpdateWarehouse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var warehouse models.Warehouse
	if err := c.ShouldBindJSON(&warehouse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	warehouse.ID = id
	if err := h.service.Update(&warehouse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, warehouse)
}

func (h *WarehouseHandler) DeleteWarehouse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Warehouse deleted"})
}

// StockMovementHandler handles stock movement requests
type StockMovementHandler struct {
	service   services.StockMovementService
	validator *validator.Validate
}

func NewStockMovementHandler(service services.StockMovementService) *StockMovementHandler {
	return &StockMovementHandler{service: service, validator: validator.New()}
}

func (h *StockMovementHandler) GetStockMovements(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Query("product_id"))
	warehouseID, _ := strconv.Atoi(c.Query("warehouse_id"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	movements, err := h.service.List(productID, warehouseID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movements)
}

// Stub handlers for documents
type SalesOrderHandler struct{ service services.SalesOrderService }

func NewSalesOrderHandler(service services.SalesOrderService) *SalesOrderHandler {
	return &SalesOrderHandler{service: service}
}
func (h *SalesOrderHandler) GetSalesOrders(c *gin.Context) { c.JSON(http.StatusOK, []interface{}{}) }
func (h *SalesOrderHandler) CreateSalesOrder(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Created"})
}
func (h *SalesOrderHandler) UpdateSalesOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}
func (h *SalesOrderHandler) DeleteSalesOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

type DeliveryOrderHandler struct{ service services.DeliveryOrderService }

func NewDeliveryOrderHandler(service services.DeliveryOrderService) *DeliveryOrderHandler {
	return &DeliveryOrderHandler{service: service}
}
func (h *DeliveryOrderHandler) GetDeliveryOrders(c *gin.Context) {
	c.JSON(http.StatusOK, []interface{}{})
}
func (h *DeliveryOrderHandler) CreateDeliveryOrder(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Created"})
}
func (h *DeliveryOrderHandler) UpdateDeliveryOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}
func (h *DeliveryOrderHandler) DeleteDeliveryOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

type PurchaseOrderHandler struct{ service services.PurchaseOrderService }

func NewPurchaseOrderHandler(service services.PurchaseOrderService) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{service: service}
}
func (h *PurchaseOrderHandler) GetPurchaseOrders(c *gin.Context) {
	c.JSON(http.StatusOK, []interface{}{})
}
func (h *PurchaseOrderHandler) CreatePurchaseOrder(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Created"})
}
func (h *PurchaseOrderHandler) UpdatePurchaseOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}
func (h *PurchaseOrderHandler) DeletePurchaseOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

type GoodsReceiptHandler struct{ service services.GoodsReceiptService }

func NewGoodsReceiptHandler(service services.GoodsReceiptService) *GoodsReceiptHandler {
	return &GoodsReceiptHandler{service: service}
}
func (h *GoodsReceiptHandler) GetGoodsReceipts(c *gin.Context) {
	c.JSON(http.StatusOK, []interface{}{})
}
func (h *GoodsReceiptHandler) CreateGoodsReceipt(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Created"})
}
func (h *GoodsReceiptHandler) UpdateGoodsReceipt(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}
func (h *GoodsReceiptHandler) DeleteGoodsReceipt(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

type TransferHandler struct{ service services.TransferService }

func NewTransferHandler(service services.TransferService) *TransferHandler {
	return &TransferHandler{service: service}
}
func (h *TransferHandler) GetTransfers(c *gin.Context) { c.JSON(http.StatusOK, []interface{}{}) }
func (h *TransferHandler) CreateTransfer(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Created"})
}
func (h *TransferHandler) UpdateTransfer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}
func (h *TransferHandler) DeleteTransfer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

type StockAdjustmentHandler struct {
	service services.StockAdjustmentService
}

func NewStockAdjustmentHandler(service services.StockAdjustmentService) *StockAdjustmentHandler {
	return &StockAdjustmentHandler{service: service}
}
func (h *StockAdjustmentHandler) GetStockAdjustments(c *gin.Context) {
	c.JSON(http.StatusOK, []interface{}{})
}
func (h *StockAdjustmentHandler) CreateStockAdjustment(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Created"})
}
func (h *StockAdjustmentHandler) UpdateStockAdjustment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}
func (h *StockAdjustmentHandler) DeleteStockAdjustment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
