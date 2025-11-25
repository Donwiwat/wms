package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

// Product represents a product in the warehouse
type Product struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	ShortName string    `json:"short_name" db:"short_name"`
	Brand     string    `json:"brand" db:"brand"`
	Model     string    `json:"model" db:"model"`
	Size      string    `json:"size" db:"size"`
	Group     string    `json:"group" db:"group"`
	Unit1     string    `json:"unit1" db:"unit1" validate:"required"`
	Unit2     string    `json:"unit2" db:"unit2"`
	Ratio     float64   `json:"ratio" db:"ratio"`
	Cost      float64   `json:"cost" db:"cost"`
	Message   string    `json:"message" db:"message"`
	Note      string    `json:"note" db:"note"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
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

// StockInRequest represents stock in request
type StockInRequest struct {
	ProductID   int     `json:"product_id" validate:"required"`
	WarehouseID int     `json:"warehouse_id" validate:"required"`
	Qty         float64 `json:"qty" validate:"required,min=0"`
	Unit        string  `json:"unit" validate:"required"`
	RefType     string  `json:"ref_type"`
	RefID       *int    `json:"ref_id"`
	Note        string  `json:"note"`
}

// StockOutRequest represents stock out request
type StockOutRequest struct {
	ProductID   int     `json:"product_id" validate:"required"`
	WarehouseID int     `json:"warehouse_id" validate:"required"`
	Qty         float64 `json:"qty" validate:"required,min=0"`
	Unit        string  `json:"unit" validate:"required"`
	RefType     string  `json:"ref_type"`
	RefID       *int    `json:"ref_id"`
	Note        string  `json:"note"`
}

// BreakDownRequest represents break down request
type BreakDownRequest struct {
	ProductID   int    `json:"product_id" validate:"required"`
	WarehouseID int    `json:"warehouse_id" validate:"required"`
	QtyUnit2    int    `json:"qty_unit2" validate:"required,min=1"`
	Note        string `json:"note"`
}

// PackUpRequest represents pack up request
type PackUpRequest struct {
	ProductID   int    `json:"product_id" validate:"required"`
	WarehouseID int    `json:"warehouse_id" validate:"required"`
	QtyUnit2    int    `json:"qty_unit2" validate:"required,min=1"`
	Note        string `json:"note"`
}

// TransferRequest represents transfer request
type TransferRequest struct {
	ProductID       int     `json:"product_id" validate:"required"`
	FromWarehouseID int     `json:"from_warehouse_id" validate:"required"`
	ToWarehouseID   int     `json:"to_warehouse_id" validate:"required"`
	Qty             float64 `json:"qty" validate:"required,min=0"`
	Unit            string  `json:"unit" validate:"required"`
	RefType         string  `json:"ref_type"`
	RefID           *int    `json:"ref_id"`
	Note            string  `json:"note"`
}

// StockAdjustRequest represents stock adjustment request
type StockAdjustRequest struct {
	ProductID   int    `json:"product_id" validate:"required"`
	WarehouseID int    `json:"warehouse_id" validate:"required"`
	NewRemain1  int    `json:"new_remain1" validate:"min=0"`
	NewRemain2  int    `json:"new_remain2" validate:"min=0"`
	Reason      string `json:"reason" validate:"required"`
	Note        string `json:"note"`
}

// LoginRequest represents login request
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// RegisterRequest represents register request
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role"`
}

// AuthResponse represents authentication response
type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
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
