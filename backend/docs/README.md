# WMS API Documentation

## Overview

This directory contains the auto-generated Swagger/OpenAPI documentation for the Warehouse Management System (WMS) API.

## Generated Files

- **`docs.go`** - Go package containing embedded Swagger specification
- **`swagger.json`** - OpenAPI 3.0 specification in JSON format
- **`swagger.yaml`** - OpenAPI 3.0 specification in YAML format

## Accessing the Documentation

### Interactive Swagger UI

When running the development server, the interactive Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

### API Endpoints

The API is organized into the following main sections:

#### Authentication (`/auth`)
- `POST /auth/login` - User authentication
- `POST /auth/register` - User registration

#### Customers (`/customers`)
- `GET /customers` - List all customers with pagination
- `GET /customers/{id}` - Get customer by ID
- `POST /customers` - Create new customer
- `PUT /customers/{id}` - Update customer
- `DELETE /customers/{id}` - Delete customer
- `GET /customers/search` - Search customers

#### Products (`/products`)
- `GET /products` - List all products
- `GET /products/{id}` - Get product by ID
- `POST /products` - Create new product
- `PUT /products/{id}` - Update product
- `DELETE /products/{id}` - Delete product
- `GET /products/{id}/prices` - Get product pricing

#### Stock Operations (`/stock`)
- `GET /stock` - Get stock summary
- `GET /stock/card` - Get stock card history
- `POST /stock/in` - Stock in operation
- `POST /stock/out` - Stock out operation
- `POST /stock/break` - Break down operation
- `POST /stock/pack` - Pack up operation
- `POST /stock/transfer` - Transfer between warehouses
- `POST /stock/adjust` - Stock adjustment

#### Orders (`/orders`)
- `GET /orders` - List orders
- `GET /orders/{id}` - Get order details
- `POST /orders` - Create new order
- `PUT /orders/{id}` - Update order
- `DELETE /orders/{id}` - Delete order
- `GET /orders/customer/{customerId}` - Get customer orders
- `PATCH /orders/{id}/status` - Update order status

#### Warehouses (`/warehouses`)
- `GET /warehouses` - List warehouses
- `GET /warehouses/{id}` - Get warehouse by ID
- `POST /warehouses` - Create warehouse
- `PUT /warehouses/{id}` - Update warehouse
- `DELETE /warehouses/{id}` - Delete warehouse

#### Document Management
- Sales Orders (`/sales-orders`)
- Purchase Orders (`/purchase-orders`)
- Delivery Orders (`/delivery-orders`)
- Goods Receipts (`/goods-receipts`)
- Transfers (`/transfers`)
- Stock Adjustments (`/stock-adjustments`)

## Authentication

The API uses JWT (JSON Web Token) authentication. To access protected endpoints:

1. **Login** via `POST /auth/login` to get a JWT token
2. **Include the token** in the `Authorization` header: `Bearer <your-jwt-token>`

### Example Authentication Flow

```bash
# 1. Login to get token
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "password123"}'

# Response: {"token": "eyJhbGciOiJIUzI1NiIs...", "user": {...}}

# 2. Use token for protected endpoints
curl -X GET http://localhost:8080/api/v1/customers \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIs..."
```

## Data Models

### Core Business Models

- **Product** - Product information with pricing and inventory details
- **Customer** - Customer information with credit terms and contact details
- **Order** - Order information with items and delivery details
- **Stock** - Current stock levels by product and warehouse
- **StockMovement** - Historical stock movement records
- **Warehouse** - Warehouse location information

### Request/Response Models

- **LoginRequest** - User authentication credentials
- **AuthResponse** - Authentication response with JWT token
- **StockInRequest** - Stock in operation parameters
- **StockOutRequest** - Stock out operation parameters
- **TransferRequest** - Stock transfer operation parameters
- **ErrorResponse** - Standard error response format
- **SuccessResponse** - Standard success response format

## Response Formats

### Success Response
```json
{
  "message": "Operation completed successfully",
  "data": { ... }
}
```

### Error Response
```json
{
  "error": "Error message",
  "code": 400,
  "details": "Detailed error information"
}
```

### Paginated Response
```json
{
  "data": [...],
  "page": 1,
  "limit": 10,
  "total": 100,
  "total_pages": 10
}
```

## Development

### Regenerating Documentation

To regenerate the Swagger documentation after making changes to the API:

```bash
# Using the provided script
./scripts/generate-docs.sh

# Or using make
make docs

# Or directly with swag
swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
```

### Validation

To validate the generated documentation:

```bash
# Install swagger CLI tool
go install github.com/go-swagger/go-swagger/cmd/swagger@latest

# Validate the specification
swagger validate docs/swagger.yaml
```

### Development Server

To start the development server with Swagger UI:

```bash
# Using make
make dev

# Or directly
GIN_MODE=debug go run cmd/server/main.go
```

## API Guidelines

### Request Headers
- `Content-Type: application/json` for POST/PUT requests
- `Authorization: Bearer <token>` for protected endpoints

### HTTP Status Codes
- `200` - Success
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `404` - Not Found
- `409` - Conflict
- `500` - Internal Server Error

### Query Parameters
- `page` - Page number for pagination (default: 1)
- `limit` - Items per page (default: 10, max: 100)
- `search` - Search term for filtering
- `q` - Query parameter for search endpoints

## Support

For API support and questions:
- Email: support@example.com
- Documentation: http://localhost:8080/swagger/index.html
- Health Check: http://localhost:8080/health