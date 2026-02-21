# WMS API Swagger Documentation Architecture

## System Architecture Overview

```mermaid
graph TB
    subgraph "Development Environment"
        DEV[Developer]
        CODE[Go Source Code]
        SWAG[Swag CLI Tool]
    end
    
    subgraph "Swagger Generation Pipeline"
        ANNOTATIONS[Swagger Annotations]
        GENERATOR[Documentation Generator]
        DOCS[Generated Docs]
    end
    
    subgraph "API Server Components"
        MAIN[main.go]
        HANDLERS[Handler Files]
        MODELS[Model Definitions]
        MIDDLEWARE[Auth Middleware]
    end
    
    subgraph "Generated Documentation"
        JSON[swagger.json]
        YAML[swagger.yaml]
        GO[docs.go]
    end
    
    subgraph "Swagger UI"
        UI[Interactive UI]
        ASSETS[Static Assets]
        CONFIG[UI Configuration]
    end
    
    subgraph "Client Applications"
        FRONTEND[Frontend App]
        MOBILE[Mobile App]
        EXTERNAL[External APIs]
    end
    
    DEV --> CODE
    CODE --> ANNOTATIONS
    ANNOTATIONS --> GENERATOR
    GENERATOR --> DOCS
    
    MAIN --> HANDLERS
    HANDLERS --> MODELS
    HANDLERS --> MIDDLEWARE
    
    DOCS --> JSON
    DOCS --> YAML
    DOCS --> GO
    
    JSON --> UI
    YAML --> UI
    GO --> UI
    
    UI --> CONFIG
    UI --> ASSETS
    
    UI --> FRONTEND
    UI --> MOBILE
    UI --> EXTERNAL
    
    SWAG --> GENERATOR
```

## API Documentation Structure

```mermaid
graph LR
    subgraph "WMS API v1"
        AUTH[Authentication]
        PROD[Products]
        CUST[Customers]
        ORDER[Orders]
        WAREHOUSE[Warehouses]
        STOCK[Stock Operations]
        DOCS[Documents]
    end
    
    subgraph "Authentication Endpoints"
        LOGIN[POST /auth/login]
        REGISTER[POST /auth/register]
    end
    
    subgraph "Product Endpoints"
        PROD_LIST[GET /products]
        PROD_GET[GET /products/:id]
        PROD_CREATE[POST /products]
        PROD_UPDATE[PUT /products/:id]
        PROD_DELETE[DELETE /products/:id]
        PROD_PRICES[GET /products/:id/prices]
    end
    
    subgraph "Stock Operations"
        STOCK_IN[POST /stock/in]
        STOCK_OUT[POST /stock/out]
        STOCK_BREAK[POST /stock/break]
        STOCK_PACK[POST /stock/pack]
        STOCK_TRANSFER[POST /stock/transfer]
        STOCK_ADJUST[POST /stock/adjust]
    end
    
    AUTH --> LOGIN
    AUTH --> REGISTER
    
    PROD --> PROD_LIST
    PROD --> PROD_GET
    PROD --> PROD_CREATE
    PROD --> PROD_UPDATE
    PROD --> PROD_DELETE
    PROD --> PROD_PRICES
    
    STOCK --> STOCK_IN
    STOCK --> STOCK_OUT
    STOCK --> STOCK_BREAK
    STOCK --> STOCK_PACK
    STOCK --> STOCK_TRANSFER
    STOCK --> STOCK_ADJUST
```

## Data Model Relationships

```mermaid
erDiagram
    Product ||--o{ ProductPrice : has
    Product ||--o{ Stock : stored_in
    Product ||--o{ StockMovement : tracks
    Product ||--o{ OrderItem : contains
    
    Customer ||--o{ Order : places
    CustomerGroup ||--o{ Customer : belongs_to
    CustomerGroup ||--o{ ProductPrice : applies_to
    
    Warehouse ||--o{ Stock : contains
    Warehouse ||--o{ StockMovement : location
    
    Order ||--o{ OrderItem : contains
    Order }o--|| Customer : placed_by
    
    StockMovement }o--|| Product : tracks
    StockMovement }o--|| Warehouse : location
    
    User ||--o{ Order : created_by
    User ||--o{ StockMovement : performed_by
```

## Swagger Implementation Flow

```mermaid
sequenceDiagram
    participant Dev as Developer
    participant Code as Source Code
    participant Swag as Swag CLI
    participant Server as API Server
    participant UI as Swagger UI
    participant Client as API Client
    
    Dev->>Code: Add Swagger annotations
    Dev->>Swag: Run swag init
    Swag->>Code: Parse annotations
    Swag->>Server: Generate docs.go
    Swag->>Server: Generate swagger.json
    
    Note over Server: Server starts with Swagger route
    
    Client->>Server: GET /swagger/index.html
    Server->>UI: Serve Swagger UI
    UI->>Server: Load swagger.json
    Server->>UI: Return API specification
    UI->>Client: Display interactive docs
    
    Client->>UI: Test API endpoint
    UI->>Server: Make API request
    Server->>UI: Return response
    UI->>Client: Display result
```

## Security Architecture

```mermaid
graph TB
    subgraph "Authentication Flow"
        LOGIN_REQ[Login Request]
        JWT_TOKEN[JWT Token]
        PROTECTED[Protected Endpoints]
    end
    
    subgraph "Swagger Security"
        SEC_DEF[Security Definition]
        BEARER[Bearer Auth]
        AUTH_HEADER[Authorization Header]
    end
    
    subgraph "API Security"
        MIDDLEWARE[Auth Middleware]
        VALIDATION[Token Validation]
        USER_CONTEXT[User Context]
    end
    
    LOGIN_REQ --> JWT_TOKEN
    JWT_TOKEN --> BEARER
    BEARER --> AUTH_HEADER
    AUTH_HEADER --> MIDDLEWARE
    MIDDLEWARE --> VALIDATION
    VALIDATION --> USER_CONTEXT
    USER_CONTEXT --> PROTECTED
    
    SEC_DEF --> BEARER
```

## File Structure and Organization

```
backend/
├── cmd/server/
│   └── main.go                 # Main Swagger config
├── docs/                       # Generated documentation
│   ├── docs.go                # Generated Go code
│   ├── swagger.json           # OpenAPI JSON spec
│   └── swagger.yaml           # OpenAPI YAML spec
├── internal/
│   ├── handlers/              # API handlers with annotations
│   │   ├── auth_handler.go    # @Tags auth
│   │   ├── product_handler.go # @Tags products
│   │   ├── customer_handler.go# @Tags customers
│   │   ├── order_handler.go   # @Tags orders
│   │   ├── stock_handler.go   # @Tags stock
│   │   └── ...
│   ├── models/                # Data models with examples
│   │   └── models.go          # @Description annotations
│   └── middleware/
│       └── auth.go            # Security middleware
├── scripts/
│   ├── generate-docs.sh       # Documentation generation
│   └── serve-docs.sh          # Development server
└── swagger/                   # Custom Swagger assets
    ├── custom.css             # UI customization
    └── logo.png               # Branding assets
```

## Implementation Phases Visualization

```mermaid
gantt
    title Swagger Documentation Implementation Timeline
    dateFormat  YYYY-MM-DD
    section Phase 1: Foundation
    Install Dependencies    :p1-1, 2024-01-01, 1d
    Setup Infrastructure    :p1-2, after p1-1, 1d
    Configure Generation    :p1-3, after p1-2, 1d
    
    section Phase 2: Models
    Core Models            :p2-1, after p1-3, 2d
    Request/Response DTOs  :p2-2, after p2-1, 2d
    Validation Rules       :p2-3, after p2-2, 1d
    
    section Phase 3: Endpoints
    Authentication APIs    :p3-1, after p2-3, 2d
    Product Management     :p3-2, after p3-1, 2d
    Customer Management    :p3-3, after p3-2, 2d
    Order Management       :p3-4, after p3-3, 2d
    Stock Operations       :p3-5, after p3-4, 2d
    Document Management    :p3-6, after p3-5, 2d
    
    section Phase 4: Advanced
    Error Documentation    :p4-1, after p3-6, 1d
    Security Integration   :p4-2, after p4-1, 1d
    UI Customization       :p4-3, after p4-2, 1d
    Examples & Guidelines  :p4-4, after p4-3, 2d
    
    section Phase 5: Automation
    CI/CD Integration      :p5-1, after p4-4, 1d
    Deployment Config      :p5-2, after p5-1, 1d
    Testing & Validation   :p5-3, after p5-2, 2d
```

## Quality Assurance Checklist

### Documentation Completeness
- [ ] All endpoints documented with proper HTTP methods
- [ ] Request/response schemas defined for all endpoints
- [ ] Authentication requirements specified
- [ ] Error responses documented with status codes
- [ ] Examples provided for complex request/response structures

### Technical Implementation
- [ ] Swagger annotations follow OpenAPI 3.0 specification
- [ ] Generated documentation validates without errors
- [ ] Interactive UI functions correctly with live API
- [ ] Security schemes work with JWT authentication
- [ ] Custom styling and branding applied

### Developer Experience
- [ ] Clear, descriptive endpoint summaries
- [ ] Comprehensive parameter descriptions
- [ ] Realistic example values provided
- [ ] Error scenarios well documented
- [ ] API usage guidelines included

This architecture provides a comprehensive foundation for implementing robust Swagger documentation that will serve as both developer reference and interactive testing environment for the WMS API.