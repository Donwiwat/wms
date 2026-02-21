package handlers

import (
	"net/http"
	"strconv"

	"wms-backend/internal/models"
	"wms-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

// GetCustomers retrieves all customers
// @Summary List all customers
// @Description Get a list of all customers in the system with optional filtering and pagination
// @Tags customers
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1) minimum(1)
// @Param limit query int false "Items per page" default(10) minimum(1) maximum(100)
// @Param search query string false "Search term for customer name"
// @Success 200 {object} models.SuccessResponse{data=[]models.Customer} "List of customers"
// @Failure 400 {object} models.ErrorResponse "Bad request"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /customers [get]
func (h *CustomerHandler) GetCustomers(c *gin.Context) {
	customers, err := h.customerService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// GetCustomer retrieves a specific customer
// @Summary Get customer by ID
// @Description Retrieve detailed information for a specific customer
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID" minimum(1)
// @Success 200 {object} models.SuccessResponse{data=models.Customer} "Customer details"
// @Failure 400 {object} models.ErrorResponse "Invalid customer ID"
// @Failure 404 {object} models.ErrorResponse "Customer not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /customers/{id} [get]
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	customer, err := h.customerService.GetByID(id)
	if err != nil {
		if err.Error() == "customer not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// CreateCustomer creates a new customer
// @Summary Create new customer
// @Description Add a new customer to the system
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body models.CustomerFormData true "Customer information"
// @Success 201 {object} models.SuccessResponse{data=models.Customer} "Created customer"
// @Failure 400 {object} models.ErrorResponse "Invalid request data"
// @Failure 409 {object} models.ErrorResponse "Customer already exists"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /customers [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var req models.CustomerFormData
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.customerService.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": customer})
}

// UpdateCustomer handles PUT /api/customers/:id
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	var req models.CustomerFormData
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.customerService.Update(id, &req)
	if err != nil {
		if err.Error() == "customer not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// DeleteCustomer handles DELETE /api/customers/:id
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	err = h.customerService.Delete(id)
	if err != nil {
		if err.Error() == "customer not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

// SearchCustomers searches for customers
// @Summary Search customers
// @Description Search for customers by name or other criteria
// @Tags customers
// @Accept json
// @Produce json
// @Param q query string true "Search query" minlength(2)
// @Param limit query int false "Maximum results" default(20) minimum(1) maximum(100)
// @Success 200 {object} models.SuccessResponse{data=[]models.Customer} "Search results"
// @Failure 400 {object} models.ErrorResponse "Invalid search parameters"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /customers/search [get]
func (h *CustomerHandler) SearchCustomers(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	customers, err := h.customerService.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customers})
}
