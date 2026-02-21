package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

// Product represents a product in the warehouse
// @Description Product information with pricing and inventory details
type Product struct {
	ID        int       `json:"id" db:"id" example:"1" description:"Unique product identifier"`
	Name      string    `json:"name" db:"name" validate:"required" example:"Premium Widget A" description:"Full product name"`
	ShortName string    `json:"short_name" db:"short_name" example:"PWA-001" description:"Short product code for quick reference"`
	Brand     string    `json:"brand" db:"brand" example:"TechCorp" description:"Product brand or manufacturer"`
	Model     string    `json:"model" db:"model" example:"TC-2024" description:"Product model number"`
	Size      string    `json:"size" db:"size" example:"Large" description:"Product size specification"`
	Group     string    `json:"group" db:"group" example:"Electronics" description:"Product category or group"`
	Unit1     string    `json:"unit1" db:"unit1" validate:"required" example:"PCS" description:"Primary unit of measurement"`
	Unit2     string    `json:"unit2" db:"unit2" example:"BOX" description:"Secondary unit of measurement"`
	Ratio     float64   `json:"ratio" db:"ratio" example:"12" description:"Conversion ratio between Unit1 and Unit2"`
	Cost      float64   `json:"cost" db:"cost" example:"25.50" description:"Product cost price"`
	Message   string    `json:"message" db:"message" example:"Handle with care" description:"Special handling instructions"`
	Note      string    `json:"note" db:"note" example:"Fragile item" description:"Additional notes"`
	CreatedAt time.Time `json:"created_at" db:"created_at" example:"2024-01-15T10:30:00Z" description:"Record creation timestamp"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" example:"2024-01-15T10:30:00Z" description:"Last update timestamp"`
}

// CustomerGroup represents customer groups for pricing
type CustomerGroup struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" validate:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// ProductPrice represents product pricing for different customer groups
type ProductPrice struct {
	ID              int       `json:"id" db:"id"`
	ProductID       int       `json:"product_id" db:"product_id" validate:"required"`
	CustomerGroupID int       `json:"customer_group_id" db:"customer_group_id" validate:"required"`
	Unit            string    `json:"unit" db:"unit" validate:"required"`
	Price           float64   `json:"price" db:"price" validate:"required,min=0"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Warehouse represents a warehouse location
type Warehouse struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Location  string    `json:"location" db:"location"`
	Note      string    `json:"note" db:"note"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Stock represents current stock levels
type Stock struct {
	ID          int       `json:"id" db:"id"`
	ProductID   int       `json:"product_id" db:"product_id"`
	WarehouseID int       `json:"warehouse_id" db:"warehouse_id"`
	Remain1     int       `json:"remain1" db:"remain1"`
	Remain2     int       `json:"remain2" db:"remain2"`
	TotalRemain int       `json:"total_remain" db:"total_remain"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// StockMovement represents all stock movements
type StockMovement struct {
	ID          int       `json:"id" db:"id"`
	ProductID   int       `json:"product_id" db:"product_id"`
	WarehouseID int       `json:"warehouse_id" db:"warehouse_id"`
	Type        string    `json:"type" db:"type"` // IN, OUT, BREAK, PACK, TF-IN, TF-OUT, ADJUST
	Qty         float64   `json:"qty" db:"qty"`
	Unit        string    `json:"unit" db:"unit"`
	RefID       *int      `json:"ref_id" db:"ref_id"`
	RefType     string    `json:"ref_type" db:"ref_type"` // PO, DO, GRN, SO, ADJ, TF
	Date        time.Time `json:"date" db:"date"`
	Note        string    `json:"note" db:"note"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// SalesOrder represents sales orders
type SalesOrder struct {
	ID        int       `json:"id" db:"id"`
	SONumber  string    `json:"so_number" db:"so_number" validate:"required"`
	Date      time.Time `json:"date" db:"date" validate:"required"`
	Customer  string    `json:"customer" db:"customer" validate:"required"`
	Note      string    `json:"note" db:"note"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// DeliveryOrder represents delivery orders
type DeliveryOrder struct {
	ID          int       `json:"id" db:"id"`
	DONumber    string    `json:"do_number" db:"do_number" validate:"required"`
	Date        time.Time `json:"date" db:"date" validate:"required"`
	Customer    string    `json:"customer" db:"customer" validate:"required"`
	WarehouseID int       `json:"warehouse_id" db:"warehouse_id" validate:"required"`
	Note        string    `json:"note" db:"note"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// PurchaseOrder represents purchase orders
type PurchaseOrder struct {
	ID        int       `json:"id" db:"id"`
	PONumber  string    `json:"po_number" db:"po_number" validate:"required"`
	Date      time.Time `json:"date" db:"date" validate:"required"`
	Supplier  string    `json:"supplier" db:"supplier" validate:"required"`
	Note      string    `json:"note" db:"note"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// GoodsReceipt represents goods receipts
type GoodsReceipt struct {
	ID          int       `json:"id" db:"id"`
	GRNNumber   string    `json:"grn_number" db:"grn_number" validate:"required"`
	Date        time.Time `json:"date" db:"date" validate:"required"`
	Supplier    string    `json:"supplier" db:"supplier" validate:"required"`
	WarehouseID int       `json:"warehouse_id" db:"warehouse_id" validate:"required"`
	Note        string    `json:"note" db:"note"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Transfer represents warehouse transfers
type Transfer struct {
	ID              int       `json:"id" db:"id"`
	TFNumber        string    `json:"tf_number" db:"tf_number" validate:"required"`
	FromWarehouseID int       `json:"from_warehouse_id" db:"from_warehouse_id" validate:"required"`
	ToWarehouseID   int       `json:"to_warehouse_id" db:"to_warehouse_id" validate:"required"`
	Date            time.Time `json:"date" db:"date" validate:"required"`
	Note            string    `json:"note" db:"note"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// StockAdjustment represents stock adjustments
type StockAdjustment struct {
	ID          int       `json:"id" db:"id"`
	AdjNumber   string    `json:"adj_number" db:"adj_number" validate:"required"`
	Date        time.Time `json:"date" db:"date" validate:"required"`
	WarehouseID int       `json:"warehouse_id" db:"warehouse_id" validate:"required"`
	Reason      string    `json:"reason" db:"reason" validate:"required"`
	Note        string    `json:"note" db:"note"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// User represents system users
type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username" validate:"required,min=3"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	Password  string    `json:"-" db:"password"`
	Role      string    `json:"role" db:"role"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Customer represents a customer
type Customer struct {
	ID            int            `json:"id" db:"id"`
	Prefix        sql.NullString `json:"-" db:"prefix"`
	Name          string         `json:"name" db:"name" validate:"required"`
	Address       sql.NullString `json:"-" db:"address"`
	Phone         sql.NullString `json:"-" db:"phone"`
	ContactPerson sql.NullString `json:"-" db:"contact_person"`
	Level         sql.NullString `json:"-" db:"level"`
	DeliveryPlace sql.NullString `json:"-" db:"delivery_place"`
	Transport     sql.NullString `json:"-" db:"transport"`
	CreditLimit   float64        `json:"credit_limit" db:"credit_limit"`
	CreditTerm    int            `json:"credit_term" db:"credit_term"`
	Outstanding   float64        `json:"outstanding" db:"outstanding"`
	LastContact   sql.NullTime   `json:"-" db:"last_contact"`
	Note          sql.NullString `json:"-" db:"note"`
	CreatedAt     time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" db:"updated_at"`
}

// MarshalJSON implements custom JSON marshaling for Customer
func (c Customer) MarshalJSON() ([]byte, error) {
	type Alias Customer
	return json.Marshal(&struct {
		Prefix        string  `json:"prefix"`
		Address       string  `json:"address"`
		Phone         string  `json:"phone"`
		ContactPerson string  `json:"contact_person"`
		Level         string  `json:"level"`
		DeliveryPlace string  `json:"delivery_place"`
		Transport     string  `json:"transport"`
		LastContact   *string `json:"last_contact"`
		Note          string  `json:"note"`
		*Alias
	}{
		Prefix:        c.Prefix.String,
		Address:       c.Address.String,
		Phone:         c.Phone.String,
		ContactPerson: c.ContactPerson.String,
		Level:         c.Level.String,
		DeliveryPlace: c.DeliveryPlace.String,
		Transport:     c.Transport.String,
		LastContact: func() *string {
			if c.LastContact.Valid {
				formatted := c.LastContact.Time.Format("2006-01-02")
				return &formatted
			}
			return nil
		}(),
		Note:  c.Note.String,
		Alias: (*Alias)(&c),
	})
}

// Order represents an order
type Order struct {
	ID              int            `json:"id" db:"id"`
	OrderNumber     string         `json:"order_number" db:"order_number"`
	CustomerID      int            `json:"customer_id" db:"customer_id"`
	OrderDate       time.Time      `json:"order_date" db:"order_date"`
	DeliveryDate    sql.NullTime   `json:"delivery_date" db:"delivery_date"`
	Status          string         `json:"status" db:"status"`
	TotalAmount     float64        `json:"total_amount" db:"total_amount"`
	Discount        float64        `json:"discount" db:"discount"`
	TaxAmount       float64        `json:"tax_amount" db:"tax_amount"`
	FinalAmount     float64        `json:"final_amount" db:"final_amount"`
	PaymentTerms    sql.NullString `json:"payment_terms" db:"payment_terms"`
	DeliveryAddress sql.NullString `json:"delivery_address" db:"delivery_address"`
	Note            sql.NullString `json:"note" db:"note"`
	CreatedBy       sql.NullString `json:"created_by" db:"created_by"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
}

// OrderItem represents an order item
type OrderItem struct {
	ID         int            `json:"id" db:"id"`
	OrderID    int            `json:"order_id" db:"order_id"`
	ProductID  int            `json:"product_id" db:"product_id"`
	Quantity   float64        `json:"quantity" db:"quantity"`
	Unit       string         `json:"unit" db:"unit"`
	UnitPrice  float64        `json:"unit_price" db:"unit_price"`
	TotalPrice float64        `json:"total_price" db:"total_price"`
	Note       sql.NullString `json:"note" db:"note"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
}

// DTOs for API requests/responses

// StockInRequest represents stock in operation request
// @Description Request payload for adding stock to warehouse
type StockInRequest struct {
	ProductID   int     `json:"product_id" validate:"required" example:"1" description:"Product identifier"`
	WarehouseID int     `json:"warehouse_id" validate:"required" example:"1" description:"Target warehouse identifier"`
	Qty         float64 `json:"qty" validate:"required,min=0" example:"100" description:"Quantity to add"`
	Unit        string  `json:"unit" validate:"required" example:"PCS" description:"Unit of measurement"`
	RefType     string  `json:"ref_type" example:"PO" description:"Reference document type (PO, GRN, etc.)"`
	RefID       *int    `json:"ref_id" example:"123" description:"Reference document ID"`
	Note        string  `json:"note" example:"Initial stock" description:"Operation notes"`
}

// StockOutRequest represents stock out operation request
// @Description Request payload for removing stock from warehouse
type StockOutRequest struct {
	ProductID   int     `json:"product_id" validate:"required" example:"1" description:"Product identifier"`
	WarehouseID int     `json:"warehouse_id" validate:"required" example:"1" description:"Source warehouse identifier"`
	Qty         float64 `json:"qty" validate:"required,min=0" example:"50" description:"Quantity to remove"`
	Unit        string  `json:"unit" validate:"required" example:"PCS" description:"Unit of measurement"`
	RefType     string  `json:"ref_type" example:"SO" description:"Reference document type (SO, DO, etc.)"`
	RefID       *int    `json:"ref_id" example:"456" description:"Reference document ID"`
	Note        string  `json:"note" example:"Sales order fulfillment" description:"Operation notes"`
}

// BreakDownRequest represents break down operation request
// @Description Request payload for breaking down products from larger to smaller units
type BreakDownRequest struct {
	ProductID   int    `json:"product_id" validate:"required" example:"1" description:"Product identifier"`
	WarehouseID int    `json:"warehouse_id" validate:"required" example:"1" description:"Warehouse identifier"`
	QtyUnit2    int    `json:"qty_unit2" validate:"required,min=1" example:"5" description:"Quantity in Unit2 to break down"`
	Note        string `json:"note" example:"Break down for retail" description:"Operation notes"`
}

// PackUpRequest represents pack up operation request
// @Description Request payload for packing products from smaller to larger units
type PackUpRequest struct {
	ProductID   int    `json:"product_id" validate:"required" example:"1" description:"Product identifier"`
	WarehouseID int    `json:"warehouse_id" validate:"required" example:"1" description:"Warehouse identifier"`
	QtyUnit2    int    `json:"qty_unit2" validate:"required,min=1" example:"3" description:"Quantity in Unit2 to pack up"`
	Note        string `json:"note" example:"Pack for wholesale" description:"Operation notes"`
}

// TransferRequest represents transfer operation request
// @Description Request payload for transferring stock between warehouses
type TransferRequest struct {
	ProductID       int     `json:"product_id" validate:"required" example:"1" description:"Product identifier"`
	FromWarehouseID int     `json:"from_warehouse_id" validate:"required" example:"1" description:"Source warehouse identifier"`
	ToWarehouseID   int     `json:"to_warehouse_id" validate:"required" example:"2" description:"Destination warehouse identifier"`
	Qty             float64 `json:"qty" validate:"required,min=0" example:"25" description:"Quantity to transfer"`
	Unit            string  `json:"unit" validate:"required" example:"PCS" description:"Unit of measurement"`
	RefType         string  `json:"ref_type" example:"TF" description:"Reference document type"`
	RefID           *int    `json:"ref_id" example:"789" description:"Reference document ID"`
	Note            string  `json:"note" example:"Rebalancing inventory" description:"Transfer notes"`
}

// StockAdjustRequest represents stock adjustment request
// @Description Request payload for adjusting stock levels (corrections)
type StockAdjustRequest struct {
	ProductID   int    `json:"product_id" validate:"required" example:"1" description:"Product identifier"`
	WarehouseID int    `json:"warehouse_id" validate:"required" example:"1" description:"Warehouse identifier"`
	NewRemain1  int    `json:"new_remain1" validate:"min=0" example:"150" description:"New quantity in Unit1"`
	NewRemain2  int    `json:"new_remain2" validate:"min=0" example:"12" description:"New quantity in Unit2"`
	Reason      string `json:"reason" validate:"required" example:"Physical count correction" description:"Reason for adjustment"`
	Note        string `json:"note" example:"Annual inventory count" description:"Additional notes"`
}

// LoginRequest represents user login credentials
// @Description User authentication request payload
type LoginRequest struct {
	Username string `json:"username" validate:"required" example:"admin" description:"User login name"`
	Password string `json:"password" validate:"required" example:"password123" description:"User password"`
}

// RegisterRequest represents user registration data
// @Description User registration request payload
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3" example:"newuser" description:"Unique username (minimum 3 characters)"`
	Email    string `json:"email" validate:"required,email" example:"user@example.com" description:"Valid email address"`
	Password string `json:"password" validate:"required,min=6" example:"securepass123" description:"Password (minimum 6 characters)"`
	Role     string `json:"role" example:"user" description:"User role (admin, user, etc.)"`
}

// AuthResponse represents authentication response
// @Description Successful authentication response with JWT token
type AuthResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." description:"JWT access token"`
	User  User   `json:"user" description:"Authenticated user information"`
}

// StockSummary represents stock summary with product and warehouse details
type StockSummary struct {
	ProductID     int       `json:"product_id" db:"product_id"`
	ProductName   string    `json:"product_name" db:"product_name"`
	WarehouseID   int       `json:"warehouse_id" db:"warehouse_id"`
	WarehouseName string    `json:"warehouse_name" db:"warehouse_name"`
	Remain1       int       `json:"remain1" db:"remain1"`
	Remain2       int       `json:"remain2" db:"remain2"`
	TotalRemain   int       `json:"total_remain" db:"total_remain"`
	Unit1         string    `json:"unit1" db:"unit1"`
	Unit2         string    `json:"unit2" db:"unit2"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// StockCardEntry represents stock card entry with details
type StockCardEntry struct {
	Date      time.Time `json:"date" db:"date"`
	Type      string    `json:"type" db:"type"`
	Qty       float64   `json:"qty" db:"qty"`
	Unit      string    `json:"unit" db:"unit"`
	RefType   string    `json:"ref_type" db:"ref_type"`
	RefID     *int      `json:"ref_id" db:"ref_id"`
	Note      string    `json:"note" db:"note"`
	CreatedBy string    `json:"created_by" db:"created_by"`
}

// CustomerFormData represents customer form data
type CustomerFormData struct {
	Prefix        string     `json:"prefix"`
	Name          string     `json:"name" validate:"required"`
	Address       string     `json:"address"`
	Phone         string     `json:"phone"`
	ContactPerson string     `json:"contact_person"`
	Level         string     `json:"level"`
	DeliveryPlace string     `json:"delivery_place"`
	Transport     string     `json:"transport"`
	CreditLimit   float64    `json:"credit_limit"`
	CreditTerm    int        `json:"credit_term"`
	Outstanding   float64    `json:"outstanding"`
	LastContact   *time.Time `json:"last_contact"`
	Note          string     `json:"note"`
}

// OrderRequest represents order creation request
type OrderRequest struct {
	CustomerID      int                `json:"customer_id" validate:"required"`
	OrderDate       time.Time          `json:"order_date" validate:"required"`
	DeliveryDate    *time.Time         `json:"delivery_date"`
	PaymentTerms    string             `json:"payment_terms"`
	DeliveryAddress string             `json:"delivery_address"`
	Note            string             `json:"note"`
	Items           []OrderItemRequest `json:"items" validate:"required,min=1"`
}

// OrderItemRequest represents order item request
type OrderItemRequest struct {
	ProductID int     `json:"product_id" validate:"required"`
	Quantity  float64 `json:"quantity" validate:"required,min=0"`
	Unit      string  `json:"unit" validate:"required"`
	UnitPrice float64 `json:"unit_price" validate:"required,min=0"`
	Note      string  `json:"note"`
}

// OrderWithDetails represents order with customer and items details
type OrderWithDetails struct {
	Order    Order                  `json:"order"`
	Customer Customer               `json:"customer"`
	Items    []OrderItemWithProduct `json:"items"`
}

// OrderItemWithProduct represents order item with product details
type OrderItemWithProduct struct {
	OrderItem   OrderItem `json:"order_item"`
	ProductName string    `json:"product_name"`
}

func (o Order) MarshalJSON() ([]byte, error) {
	type Alias Order
	return json.Marshal(&struct {
		PaymentTerms    string  `json:"payment_terms"`
		DeliveryAddress string  `json:"delivery_address"`
		Note            string  `json:"note"`
		DeliveryDate    *string `json:"delivery_date"`
		*Alias
	}{
		PaymentTerms:    o.PaymentTerms.String,
		DeliveryAddress: o.DeliveryAddress.String,
		Note:            o.Note.String,
		DeliveryDate: func() *string {
			if o.DeliveryDate.Valid {
				s := o.DeliveryDate.Time.Format("2006-01-02T15:04:05Z")
				return &s
			}
			return nil
		}(),
		Alias: (*Alias)(&o),
	})
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

// PaginationResponse represents paginated response
// @Description Standard pagination response format
type PaginationResponse struct {
	Data       interface{} `json:"data" description:"Response data array"`
	Page       int         `json:"page" example:"1" description:"Current page number"`
	Limit      int         `json:"limit" example:"10" description:"Items per page"`
	Total      int64       `json:"total" example:"100" description:"Total number of items"`
	TotalPages int         `json:"total_pages" example:"10" description:"Total number of pages"`
}
