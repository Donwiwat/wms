package handlers

import (
	"net/http"
	"strconv"

	"wms-backend/internal/models"
	"wms-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// GetOrders godoc
// @Summary List all orders
// @Description Get a list of all orders in the system with optional filtering and pagination
// @Tags orders
// @Accept json
// @Produce json
// @Param page query int false "Page number" minimum(1) default(1)
// @Param limit query int false "Items per page" minimum(1) maximum(100) default(10)
// @Param status query string false "Filter by order status"
// @Param customer_id query int false "Filter by customer ID"
// @Success 200 {object} models.SuccessResponse "List of orders"
// @Failure 400 {object} models.ErrorResponse "Bad request"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /orders [get]
func (h *OrderHandler) GetOrders(c *gin.Context) {
	orders, err := h.orderService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// GetOrder godoc
// @Summary Get order by ID
// @Description Retrieve detailed information for a specific order including customer and items
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID" minimum(1)
// @Success 200 {object} models.SuccessResponse "Order details"
// @Failure 400 {object} models.ErrorResponse "Invalid order ID"
// @Failure 404 {object} models.ErrorResponse "Order not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.orderService.GetByID(id)
	if err != nil {
		if err.Error() == "order not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// CreateOrder godoc
// @Summary Create new order
// @Description Create a new order with items for a customer
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.OrderRequest true "Order information with items"
// @Success 201 {object} models.SuccessResponse "Created order"
// @Failure 400 {object} models.ErrorResponse "Invalid request data"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req models.OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderService.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": order})
}

// UpdateOrder godoc
// @Summary Update order
// @Description Update an existing order's information and items
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID" minimum(1)
// @Param order body models.OrderRequest true "Updated order information"
// @Success 200 {object} models.SuccessResponse "Updated order"
// @Failure 400 {object} models.ErrorResponse "Invalid request data"
// @Failure 404 {object} models.ErrorResponse "Order not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var req models.OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderService.Update(id, &req)
	if err != nil {
		if err.Error() == "order not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// DeleteOrder godoc
// @Summary Delete order
// @Description Remove an order from the system
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID" minimum(1)
// @Success 200 {object} models.SuccessResponse "Order deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid order ID"
// @Failure 404 {object} models.ErrorResponse "Order not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	err = h.orderService.Delete(id)
	if err != nil {
		if err.Error() == "order not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// GetOrdersByCustomer godoc
// @Summary Get orders by customer
// @Description Retrieve all orders for a specific customer
// @Tags orders
// @Accept json
// @Produce json
// @Param customerId path int true "Customer ID" minimum(1)
// @Success 200 {object} models.SuccessResponse "Customer orders"
// @Failure 400 {object} models.ErrorResponse "Invalid customer ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /orders/customer/{customerId} [get]
func (h *OrderHandler) GetOrdersByCustomer(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("customerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	orders, err := h.orderService.GetByCustomerID(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// UpdateOrderStatus godoc
// @Summary Update order status
// @Description Update the status of an existing order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID" minimum(1)
// @Param status body object{status=string} true "New order status"
// @Success 200 {object} models.SuccessResponse "Order status updated successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request data"
// @Failure 404 {object} models.ErrorResponse "Order not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /orders/{id}/status [patch]
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.orderService.UpdateStatus(id, req.Status)
	if err != nil {
		if err.Error() == "order not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}
