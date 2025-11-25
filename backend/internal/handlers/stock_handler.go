package handlers

import (
	"net/http"
	"strconv"

	"wms-backend/internal/models"
	"wms-backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// StockHandler handles stock-related requests
type StockHandler struct {
	stockService services.StockService
	validator    *validator.Validate
}

// NewStockHandler creates a new stock handler
func NewStockHandler(stockService services.StockService) *StockHandler {
	return &StockHandler{
		stockService: stockService,
		validator:    validator.New(),
	}
}

// GetStock gets stock summary
// @Summary Get stock summary
// @Description Get stock summary with optional filters
// @Tags stock
// @Accept json
// @Produce json
// @Param product_id query int false "Product ID filter"
// @Param warehouse_id query int false "Warehouse ID filter"
// @Success 200 {array} models.StockSummary
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /stock [get]
func (h *StockHandler) GetStock(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Query("product_id"))
	warehouseID, _ := strconv.Atoi(c.Query("warehouse_id"))

	stocks, err := h.stockService.GetStockSummary(productID, warehouseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)
}

// GetStockCard gets stock movement history
// @Summary Get stock card
// @Description Get stock movement history for a product and warehouse
// @Tags stock
// @Accept json
// @Produce json
// @Param product_id query int true "Product ID"
// @Param warehouse_id query int true "Warehouse ID"
// @Success 200 {array} models.StockCardEntry
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /stock/card [get]
func (h *StockHandler) GetStockCard(c *gin.Context) {
	productID, err := strconv.Atoi(c.Query("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product_id"})
		return
	}

	warehouseID, err := strconv.Atoi(c.Query("warehouse_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warehouse_id"})
		return
	}

	entries, err := h.stockService.GetStockCard(productID, warehouseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entries)
}

// StockIn handles stock in operation
// @Summary Stock in
// @Description Add stock to warehouse
// @Tags stock
// @Accept json
// @Produce json
// @Param request body models.StockInRequest true "Stock in details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /stock/in [post]
func (h *StockHandler) StockIn(c *gin.Context) {
	var req models.StockInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, _ := c.Get("username")
	usernameStr, _ := username.(string)

	if err := h.stockService.StockIn(&req, usernameStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock in successful"})
}

// StockOut handles stock out operation
// @Summary Stock out
// @Description Remove stock from warehouse
// @Tags stock
// @Accept json
// @Produce json
// @Param request body models.StockOutRequest true "Stock out details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /stock/out [post]
func (h *StockHandler) StockOut(c *gin.Context) {
	var req models.StockOutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, _ := c.Get("username")
	usernameStr, _ := username.(string)

	if err := h.stockService.StockOut(&req, usernameStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock out successful"})
}

// BreakDown handles break down operation
// @Summary Break down
// @Description Break down unit2 to unit1
// @Tags stock
// @Accept json
// @Produce json
// @Param request body models.BreakDownRequest true "Break down details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /stock/break [post]
func (h *StockHandler) BreakDown(c *gin.Context) {
	var req models.BreakDownRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, _ := c.Get("username")
	usernameStr, _ := username.(string)

	if err := h.stockService.BreakDown(&req, usernameStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Break down successful"})
}

// PackUp handles pack up operation
// @Summary Pack up
// @Description Pack up unit1 to unit2
// @Tags stock
// @Accept json
// @Produce json
// @Param request body models.PackUpRequest true "Pack up details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /stock/pack [post]
func (h *StockHandler) PackUp(c *gin.Context) {
	var req models.PackUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, _ := c.Get("username")
	usernameStr, _ := username.(string)

	if err := h.stockService.PackUp(&req, usernameStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pack up successful"})
}

// Transfer handles transfer operation
// @Summary Transfer stock
// @Description Transfer stock between warehouses
// @Tags stock
// @Accept json
// @Produce json
// @Param request body models.TransferRequest true "Transfer details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /stock/transfer [post]
func (h *StockHandler) Transfer(c *gin.Context) {
	var req models.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, _ := c.Get("username")
	usernameStr, _ := username.(string)

	if err := h.stockService.Transfer(&req, usernameStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}

// Adjust handles stock adjustment
// @Summary Adjust stock
// @Description Adjust stock levels
// @Tags stock
// @Accept json
// @Produce json
// @Param request body models.StockAdjustRequest true "Adjustment details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /stock/adjust [post]
func (h *StockHandler) Adjust(c *gin.Context) {
	var req models.StockAdjustRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, _ := c.Get("username")
	usernameStr, _ := username.(string)

	if err := h.stockService.Adjust(&req, usernameStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock adjustment successful"})
}
