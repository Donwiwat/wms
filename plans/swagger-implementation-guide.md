# WMS API Swagger Implementation Guide

## Quick Start Implementation

### 1. Dependencies Configuration

**Update [`go.mod`](backend/go.mod) to include Swagger dependencies:**

```go
require (
    // ... existing dependencies
    github.com/swaggo/swag v1.16.2
    github.com/swaggo/gin-swagger v1.6.0
    github.com/swaggo/files v1.0.1
)
```

**Install Swag CLI tool:**
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 2. Main Server Configuration

**Enhanced [`main.go`](backend/cmd/server/main.go) with complete Swagger setup:**

```go
package main

import (
    "log"
    "os"

    "wms-backend/internal/config"
    "wms-backend/internal/database"
    "wms-backend/internal/handlers"
    "wms-backend/internal/middleware"
    "wms-backend/internal/repositories"
    "wms-backend/internal/services"
    
    // Import generated docs
    _ "wms-backend/docs"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/rs/cors"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title WMS API
// @version 1.0
// @description Warehouse Management System API with comprehensive inventory, order, and customer management capabilities
// @termsOfService http://swagger.io/terms/

// @contact.name WMS API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @tag.name auth
// @tag.description Authentication operations

// @tag.name products
// @tag.description Product management operations

// @tag.name customers
// @tag.description Customer management operations

// @tag.name orders
// @tag.description Order management operations

// @tag.name warehouses
// @tag.description Warehouse management operations

// @tag.name stock
// @tag.description Stock operations and inventory management

// @tag.name documents
// @tag.description Document management (Sales Orders, Purchase Orders, etc.)

func main() {
    // ... existing main function code ...
    
    // Setup router with Swagger
    router := setupRouter(cfg, handlers)
    
    // Start server
    log.Printf("Server starting on %s:%s", cfg.Server.Host, cfg.Server.Port)
    log.Printf("Swagger UI available at: http://%s:%s/swagger/index.html", cfg.Server.Host, cfg.Server.Port)
    if err := router.Run(cfg.Server.Host + ":" + cfg.Server.Port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}

func setupRouter(cfg *config.Config, h *handlers.Handlers) *gin.Engine {
    // ... existing router setup ...
    
    // Swagger endpoint
    if os.Getenv("GIN_MODE") != "release" {
        router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    }
    
    // ... rest of router setup ...
    
    return router
}
```

### 3. Model Documentation Examples

**Enhanced [`models.go`](backend/internal/models/models.go) with comprehensive Swagger annotations:**

```go
// Product represents a product in the warehouse
// @Description Product information with pricing and inventory details
type Product struct {
    ID        int       `json:"id" example:"1" description:"Unique product identifier"`
    Name      string    `json:"name" example:"Premium Widget A" validate:"required" description:"Full product name"`
    ShortName string    `json:"short_name" example:"PWA-001" description:"Short product code for quick reference"`
    Brand     string    `json:"brand" example:"TechCorp" description:"Product brand or manufacturer"`
    Model     string    `json:"model" example:"TC-2024" description:"Product model number"`
    Size      string    `json:"size" example:"Large" description:"Product size specification"`
    Group     string    `json:"group" example:"Electronics" description:"Product category or group"`
    Unit1     string    `json:"unit1" example:"PCS" validate:"required" description:"Primary unit of measurement"`
    Unit2     string    `json:"unit2" example:"BOX" description:"Secondary unit of measurement"`
    Ratio     float64   `json:"ratio" example:"12" description:"Conversion ratio between Unit1 and Unit2"`
    Cost      float64   `json:"cost" example:"25.50" description:"Product cost price"`
    Message   string    `json:"message" example:"Handle with care" description:"Special handling instructions"`
    Note      string    `json:"note" example:"Fragile item" description:"Additional notes"`
    CreatedAt time.Time `json:"created_at" example:"2024-01-15T10:30:00Z" description:"Record creation timestamp"`
    UpdatedAt time.Time `json:"updated_at" example:"2024-01-15T10:30:00Z" description:"Last update timestamp"`
}

// LoginRequest represents user login credentials
// @Description User authentication request payload
type LoginRequest struct {
    Username string `json:"username" example:"admin" validate:"required" description:"User login name"`
    Password string `json:"password" example:"password123" validate:"required" description:"User password"`
}

// AuthResponse represents authentication response
// @Description Successful authentication response with JWT token
type AuthResponse struct {
    Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." description:"JWT access token"`
    User  User   `json:"user" description:"Authenticated user information"`
}

// StockInRequest represents stock in operation request
// @Description Request payload for adding stock to warehouse
type StockInRequest struct {
    ProductID   int     `json:"product_id" example:"1" validate:"required" description:"Product identifier"`
    WarehouseID int     `json:"warehouse_id" example:"1" validate:"required" description:"Target warehouse identifier"`
    Qty         float64 `json:"qty" example:"100" validate:"required,min=0" description:"Quantity to add"`
    Unit        string  `json:"unit" example:"PCS" validate:"required" description:"Unit of measurement"`
    RefType     string  `json:"ref_type" example:"PO" description:"Reference document type (PO, GRN, etc.)"`
    RefID       *int    `json:"ref_id" example:"123" description:"Reference document ID"`
    Note        string  `json:"note" example:"Initial stock" description:"Operation notes"`
}

// ErrorResponse represents API error response
// @Description Standard error response format
type ErrorResponse struct {
    Error   string `json:"error" example:"Invalid request parameters" description:"Error message"`
    Code    int    `json:"code,omitempty" example:"400" description:"Error code"`
    Details string `json:"details,omitempty" example:"Field 'name' is required" description:"Detailed error information"`
}

// SuccessResponse represents successful operation response
// @Description Standard success response format
type SuccessResponse struct {
    Message string      `json:"message" example:"Operation completed successfully" description:"Success message"`
    Data    interface{} `json:"data,omitempty" description:"Response data"`
}
```

### 4. Handler Documentation Examples

**Enhanced [`customer_handler.go`](backend/internal/handlers/customer_handler.go) with complete Swagger annotations:**

```go
// GetCustomers retrieves all customers
// @Summary List all customers
// @Description Get a list of all customers in the system with optional filtering
// @Tags customers
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search term for customer name"
// @Success 200 {object} SuccessResponse{data=[]models.Customer} "List of customers"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /customers [get]
func (h *CustomerHandler) GetCustomers(c *gin.Context) {
    // ... existing implementation ...
}

// GetCustomer retrieves a specific customer
// @Summary Get customer by ID
// @Description Retrieve detailed information for a specific customer
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID" minimum(1)
// @Success 200 {object} SuccessResponse{data=models.Customer} "Customer details"
// @Failure 400 {object} ErrorResponse "Invalid customer ID"
// @Failure 404 {object} ErrorResponse "Customer not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /customers/{id} [get]
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
    // ... existing implementation ...
}

// CreateCustomer creates a new customer
// @Summary Create new customer
// @Description Add a new customer to the system
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body models.CustomerFormData true "Customer information"
// @Success 201 {object} SuccessResponse{data=models.Customer} "Created customer"
// @Failure 400 {object} ErrorResponse "Invalid request data"
// @Failure 409 {object} ErrorResponse "Customer already exists"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /customers [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
    // ... existing implementation ...
}

// SearchCustomers searches for customers
// @Summary Search customers
// @Description Search for customers by name or other criteria
// @Tags customers
// @Accept json
// @Produce json
// @Param q query string true "Search query" minlength(2)
// @Param limit query int false "Maximum results" default(20)
// @Success 200 {object} SuccessResponse{data=[]models.Customer} "Search results"
// @Failure 400 {object} ErrorResponse "Invalid search parameters"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /customers/search [get]
func (h *CustomerHandler) SearchCustomers(c *gin.Context) {
    // ... existing implementation ...
}
```

**Authentication Handler Example:**

```go
// Login authenticates user and returns JWT token
// @Summary User login
// @Description Authenticate user credentials and return JWT access token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "User login credentials"
// @Success 200 {object} models.AuthResponse "Successful authentication"
// @Failure 400 {object} ErrorResponse "Invalid request format"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
    // ... existing implementation ...
}

// Register creates a new user account
// @Summary User registration
// @Description Create a new user account in the system
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.RegisterRequest true "User registration data"
// @Success 201 {object} models.AuthResponse "User created successfully"
// @Failure 400 {object} ErrorResponse "Invalid request data"
// @Failure 409 {object} ErrorResponse "User already exists"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
    // ... existing implementation ...
}
```

**Stock Operations Handler Example:**

```go
// StockIn adds stock to warehouse
// @Summary Add stock to warehouse
// @Description Perform stock in operation to increase inventory levels
// @Tags stock
// @Accept json
// @Produce json
// @Param operation body models.StockInRequest true "Stock in operation details"
// @Success 200 {object} SuccessResponse{data=models.StockMovement} "Stock operation completed"
// @Failure 400 {object} ErrorResponse "Invalid operation parameters"
// @Failure 404 {object} ErrorResponse "Product or warehouse not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /stock/in [post]
func (h *StockHandler) StockIn(c *gin.Context) {
    // ... existing implementation ...
}

// Transfer moves stock between warehouses
// @Summary Transfer stock between warehouses
// @Description Move inventory from one warehouse to another
// @Tags stock
// @Accept json
// @Produce json
// @Param transfer body models.TransferRequest true "Transfer operation details"
// @Success 200 {object} SuccessResponse "Transfer completed successfully"
// @Failure 400 {object} ErrorResponse "Invalid transfer parameters"
// @Failure 404 {object} ErrorResponse "Product or warehouse not found"
// @Failure 409 {object} ErrorResponse "Insufficient stock"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /stock/transfer [post]
func (h *StockHandler) Transfer(c *gin.Context) {
    // ... existing implementation ...
}
```

### 5. Generation Scripts

**Create `scripts/generate-docs.sh`:**

```bash
#!/bin/bash

echo "Generating Swagger documentation..."

# Navigate to backend directory
cd backend

# Generate Swagger docs
swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal

echo "Swagger documentation generated successfully!"
echo "Files created:"
echo "  - docs/docs.go"
echo "  - docs/swagger.json"
echo "  - docs/swagger.yaml"

# Validate generated documentation
if [ -f "docs/swagger.json" ]; then
    echo "✅ swagger.json generated"
else
    echo "❌ swagger.json not found"
    exit 1
fi

if [ -f "docs/swagger.yaml" ]; then
    echo "✅ swagger.yaml generated"
else
    echo "❌ swagger.yaml not found"
    exit 1
fi

echo "🎉 Documentation generation completed!"
```

**Create `scripts/serve-docs.sh`:**

```bash
#!/bin/bash

echo "Starting WMS API server with Swagger documentation..."

cd backend

# Set development mode
export GIN_MODE=debug

# Start the server
go run cmd/server/main.go

echo "Server started!"
echo "API Documentation: http://localhost:8080/swagger/index.html"
echo "Health Check: http://localhost:8080/health"
```

### 6. Docker Integration

**Update `docker-compose.yml` to include Swagger documentation:**

```yaml
version: '3.8'

services:
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=debug
    volumes:
      - ./backend:/app
    labels:
      - "traefik.http.routers.api.rule=PathPrefix(`/api`) || PathPrefix(`/swagger`)"
      - "traefik.http.services.api.loadbalancer.server.port=8080"

  swagger-ui:
    image: swaggerapi/swagger-ui
    ports:
      - "8081:8080"
    environment:
      - SWAGGER_JSON=/swagger/swagger.json
    volumes:
      - ./backend/docs:/swagger
    depends_on:
      - backend
```

### 7. Makefile for Automation

**Create `Makefile` in backend directory:**

```makefile
.PHONY: docs serve-docs clean-docs validate-docs

# Generate Swagger documentation
docs:
	@echo "Generating Swagger documentation..."
	swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
	@echo "Documentation generated successfully!"

# Serve documentation in development
serve-docs: docs
	@echo "Starting server with Swagger UI..."
	GIN_MODE=debug go run cmd/server/main.go

# Clean generated documentation
clean-docs:
	@echo "Cleaning generated documentation..."
	rm -rf docs/
	@echo "Documentation cleaned!"

# Validate generated documentation
validate-docs: docs
	@echo "Validating Swagger documentation..."
	swagger validate docs/swagger.yaml
	@echo "Documentation is valid!"

# Install required tools
install-tools:
	@echo "Installing Swagger tools..."
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	@echo "Tools installed successfully!"

# Run all checks
check: clean-docs docs validate-docs
	@echo "All documentation checks passed!"
```

### 8. CI/CD Integration

**GitHub Actions workflow (`.github/workflows/swagger.yml`):**

```yaml
name: Swagger Documentation

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  generate-docs:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21
    
    - name: Install Swagger tools
      run: |
        go install github.com/swaggo/swag/cmd/swag@latest
        go install github.com/go-swagger/go-swagger/cmd/swagger@latest
    
    - name: Generate Swagger docs
      run: |
        cd backend
        swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
    
    - name: Validate documentation
      run: |
        cd backend
        swagger validate docs/swagger.yaml
    
    - name: Upload documentation artifacts
      uses: actions/upload-artifact@v3
      with:
        name: swagger-docs
        path: backend/docs/
```

### 9. Development Workflow

**Daily development process:**

1. **Add Swagger annotations** to new endpoints
2. **Run documentation generation**: `make docs`
3. **Test in Swagger UI**: `make serve-docs`
4. **Validate documentation**: `make validate-docs`
5. **Commit changes** including generated docs

**Example development commands:**

```bash
# Start development with live documentation
make serve-docs

# Generate docs only
make docs

# Clean and regenerate
make clean-docs docs

# Full validation check
make check
```

This implementation guide provides everything needed to get comprehensive Swagger documentation up and running for the WMS API, with automated generation, validation, and deployment capabilities.