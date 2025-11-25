package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"wms-backend/internal/models"
)

type OrderRepository interface {
	GetAll() ([]models.OrderWithDetails, error)
	GetByID(id int) (*models.OrderWithDetails, error)
	Create(order *models.OrderRequest) (*models.Order, error)
	Update(id int, order *models.OrderRequest) (*models.Order, error)
	Delete(id int) error
	GetByCustomerID(customerID int) ([]models.OrderWithDetails, error)
	UpdateStatus(id int, status string) error
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAll() ([]models.OrderWithDetails, error) {
	query := `
		SELECT o.id, o.order_number, o.customer_id, o.order_date, o.delivery_date,
		       o.status, o.total_amount, o.discount, o.tax_amount, o.final_amount,
		       o.payment_terms, o.delivery_address, o.note, o.created_by,
		       o.created_at, o.updated_at,
		       c.id, c.prefix, c.name, c.address, c.phone, c.contact_person,
		       c.level, c.delivery_place, c.transport, c.credit_limit,
		       c.credit_term, c.outstanding, c.last_contact, c.note,
		       c.created_at, c.updated_at
		FROM orders o
		JOIN customers c ON o.customer_id = c.id
		ORDER BY o.created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}
	defer rows.Close()

	var orders []models.OrderWithDetails
	for rows.Next() {
		var orderDetail models.OrderWithDetails
		err := rows.Scan(
			&orderDetail.Order.ID, &orderDetail.Order.OrderNumber, &orderDetail.Order.CustomerID,
			&orderDetail.Order.OrderDate, &orderDetail.Order.DeliveryDate, &orderDetail.Order.Status,
			&orderDetail.Order.TotalAmount, &orderDetail.Order.Discount, &orderDetail.Order.TaxAmount,
			&orderDetail.Order.FinalAmount, &orderDetail.Order.PaymentTerms, &orderDetail.Order.DeliveryAddress,
			&orderDetail.Order.Note, &orderDetail.Order.CreatedBy, &orderDetail.Order.CreatedAt,
			&orderDetail.Order.UpdatedAt,
			&orderDetail.Customer.ID, &orderDetail.Customer.Prefix, &orderDetail.Customer.Name,
			&orderDetail.Customer.Address, &orderDetail.Customer.Phone, &orderDetail.Customer.ContactPerson,
			&orderDetail.Customer.Level, &orderDetail.Customer.DeliveryPlace, &orderDetail.Customer.Transport,
			&orderDetail.Customer.CreditLimit, &orderDetail.Customer.CreditTerm, &orderDetail.Customer.Outstanding,
			&orderDetail.Customer.LastContact, &orderDetail.Customer.Note, &orderDetail.Customer.CreatedAt,
			&orderDetail.Customer.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}

		// Get order items
		items, err := r.getOrderItems(orderDetail.Order.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get order items: %w", err)
		}
		orderDetail.Items = items

		orders = append(orders, orderDetail)
	}

	return orders, nil
}

func (r *orderRepository) GetByID(id int) (*models.OrderWithDetails, error) {
	query := `
		SELECT o.id, o.order_number, o.customer_id, o.order_date, o.delivery_date,
		       o.status, o.total_amount, o.discount, o.tax_amount, o.final_amount,
		       o.payment_terms, o.delivery_address, o.note, o.created_by,
		       o.created_at, o.updated_at,
		       c.id, c.prefix, c.name, c.address, c.phone, c.contact_person,
		       c.level, c.delivery_place, c.transport, c.credit_limit,
		       c.credit_term, c.outstanding, c.last_contact, c.note,
		       c.created_at, c.updated_at
		FROM orders o
		JOIN customers c ON o.customer_id = c.id
		WHERE o.id = $1
	`

	var orderDetail models.OrderWithDetails
	err := r.db.QueryRow(query, id).Scan(
		&orderDetail.Order.ID, &orderDetail.Order.OrderNumber, &orderDetail.Order.CustomerID,
		&orderDetail.Order.OrderDate, &orderDetail.Order.DeliveryDate, &orderDetail.Order.Status,
		&orderDetail.Order.TotalAmount, &orderDetail.Order.Discount, &orderDetail.Order.TaxAmount,
		&orderDetail.Order.FinalAmount, &orderDetail.Order.PaymentTerms, &orderDetail.Order.DeliveryAddress,
		&orderDetail.Order.Note, &orderDetail.Order.CreatedBy, &orderDetail.Order.CreatedAt,
		&orderDetail.Order.UpdatedAt,
		&orderDetail.Customer.ID, &orderDetail.Customer.Prefix, &orderDetail.Customer.Name,
		&orderDetail.Customer.Address, &orderDetail.Customer.Phone, &orderDetail.Customer.ContactPerson,
		&orderDetail.Customer.Level, &orderDetail.Customer.DeliveryPlace, &orderDetail.Customer.Transport,
		&orderDetail.Customer.CreditLimit, &orderDetail.Customer.CreditTerm, &orderDetail.Customer.Outstanding,
		&orderDetail.Customer.LastContact, &orderDetail.Customer.Note, &orderDetail.Customer.CreatedAt,
		&orderDetail.Customer.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("order not found")
		}
		return nil, fmt.Errorf("failed to get order: %w", err)
	}

	// Get order items
	items, err := r.getOrderItems(orderDetail.Order.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order items: %w", err)
	}
	orderDetail.Items = items

	return &orderDetail, nil
}

func (r *orderRepository) Create(orderReq *models.OrderRequest) (*models.Order, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Generate order number
	orderNumber, err := r.generateOrderNumber(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate order number: %w", err)
	}

	// Calculate totals
	var totalAmount float64
	for _, item := range orderReq.Items {
		totalAmount += item.Quantity * item.UnitPrice
	}

	finalAmount := totalAmount - 0 + 0 // totalAmount - discount + tax

	// Create order
	orderQuery := `
		INSERT INTO orders (order_number, customer_id, order_date, delivery_date,
		                   status, total_amount, discount, tax_amount, final_amount,
		                   payment_terms, delivery_address, note, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id, created_at, updated_at
	`

	var order models.Order
	err = tx.QueryRow(
		orderQuery,
		orderNumber, orderReq.CustomerID, orderReq.OrderDate, orderReq.DeliveryDate,
		"pending", totalAmount, 0, 0, finalAmount,
		orderReq.PaymentTerms, orderReq.DeliveryAddress, orderReq.Note, "system",
	).Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	// Set order fields
	order.OrderNumber = orderNumber
	order.CustomerID = orderReq.CustomerID
	order.OrderDate = orderReq.OrderDate
	if orderReq.DeliveryDate != nil {
		order.DeliveryDate = sql.NullTime{Time: *orderReq.DeliveryDate, Valid: true}
	} else {
		order.DeliveryDate = sql.NullTime{Valid: false}
	}
	order.Status = "pending"
	order.TotalAmount = totalAmount
	order.Discount = 0
	order.TaxAmount = 0
	order.FinalAmount = finalAmount
	order.PaymentTerms = sql.NullString{String: orderReq.PaymentTerms, Valid: orderReq.PaymentTerms != ""}
	order.DeliveryAddress = sql.NullString{String: orderReq.DeliveryAddress, Valid: orderReq.DeliveryAddress != ""}
	order.Note = sql.NullString{String: orderReq.Note, Valid: orderReq.Note != ""}
	order.CreatedBy = sql.NullString{String: "system", Valid: true}

	// Create order items
	for _, item := range orderReq.Items {
		itemQuery := `
			INSERT INTO order_items (order_id, product_id, quantity, unit, unit_price, total_price, note)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`
		totalPrice := item.Quantity * item.UnitPrice
		_, err = tx.Exec(itemQuery, order.ID, item.ProductID, item.Quantity, item.Unit, item.UnitPrice, totalPrice, item.Note)
		if err != nil {
			return nil, fmt.Errorf("failed to create order item: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &order, nil
}

func (r *orderRepository) Update(id int, orderReq *models.OrderRequest) (*models.Order, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Calculate totals
	var totalAmount float64
	for _, item := range orderReq.Items {
		totalAmount += item.Quantity * item.UnitPrice
	}

	finalAmount := totalAmount - 0 + 0 // totalAmount - discount + tax

	// Update order
	orderQuery := `
		UPDATE orders 
		SET customer_id = $2, order_date = $3, delivery_date = $4,
		    total_amount = $5, final_amount = $6, payment_terms = $7,
		    delivery_address = $8, note = $9
		WHERE id = $1
		RETURNING order_number, status, discount, tax_amount, created_by, created_at, updated_at
	`

	var order models.Order
	err = tx.QueryRow(
		orderQuery, id,
		orderReq.CustomerID, orderReq.OrderDate, orderReq.DeliveryDate,
		totalAmount, finalAmount, orderReq.PaymentTerms,
		orderReq.DeliveryAddress, orderReq.Note,
	).Scan(&order.OrderNumber, &order.Status, &order.Discount, &order.TaxAmount, &order.CreatedBy, &order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("order not found")
		}
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	// Set order fields
	order.ID = id
	order.CustomerID = orderReq.CustomerID
	order.OrderDate = orderReq.OrderDate
	if orderReq.DeliveryDate != nil {
		order.DeliveryDate = sql.NullTime{Time: *orderReq.DeliveryDate, Valid: true}
	} else {
		order.DeliveryDate = sql.NullTime{Valid: false}
	}
	order.TotalAmount = totalAmount
	order.FinalAmount = finalAmount
	order.PaymentTerms = sql.NullString{String: orderReq.PaymentTerms, Valid: orderReq.PaymentTerms != ""}
	order.DeliveryAddress = sql.NullString{String: orderReq.DeliveryAddress, Valid: orderReq.DeliveryAddress != ""}
	order.Note = sql.NullString{String: orderReq.Note, Valid: orderReq.Note != ""}

	// Delete existing order items
	_, err = tx.Exec("DELETE FROM order_items WHERE order_id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete existing order items: %w", err)
	}

	// Create new order items
	for _, item := range orderReq.Items {
		itemQuery := `
			INSERT INTO order_items (order_id, product_id, quantity, unit, unit_price, total_price, note)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`
		totalPrice := item.Quantity * item.UnitPrice
		_, err = tx.Exec(itemQuery, id, item.ProductID, item.Quantity, item.Unit, item.UnitPrice, totalPrice, item.Note)
		if err != nil {
			return nil, fmt.Errorf("failed to create order item: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &order, nil
}

func (r *orderRepository) Delete(id int) error {
	query := `DELETE FROM orders WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("order not found")
	}

	return nil
}

func (r *orderRepository) GetByCustomerID(customerID int) ([]models.OrderWithDetails, error) {
	query := `
		SELECT o.id, o.order_number, o.customer_id, o.order_date, o.delivery_date,
		       o.status, o.total_amount, o.discount, o.tax_amount, o.final_amount,
		       o.payment_terms, o.delivery_address, o.note, o.created_by,
		       o.created_at, o.updated_at,
		       c.id, c.prefix, c.name, c.address, c.phone, c.contact_person,
		       c.level, c.delivery_place, c.transport, c.credit_limit,
		       c.credit_term, c.outstanding, c.last_contact, c.note,
		       c.created_at, c.updated_at
		FROM orders o
		JOIN customers c ON o.customer_id = c.id
		WHERE o.customer_id = $1
		ORDER BY o.created_at DESC
	`

	rows, err := r.db.Query(query, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders by customer: %w", err)
	}
	defer rows.Close()

	var orders []models.OrderWithDetails
	for rows.Next() {
		var orderDetail models.OrderWithDetails
		err := rows.Scan(
			&orderDetail.Order.ID, &orderDetail.Order.OrderNumber, &orderDetail.Order.CustomerID,
			&orderDetail.Order.OrderDate, &orderDetail.Order.DeliveryDate, &orderDetail.Order.Status,
			&orderDetail.Order.TotalAmount, &orderDetail.Order.Discount, &orderDetail.Order.TaxAmount,
			&orderDetail.Order.FinalAmount, &orderDetail.Order.PaymentTerms, &orderDetail.Order.DeliveryAddress,
			&orderDetail.Order.Note, &orderDetail.Order.CreatedBy, &orderDetail.Order.CreatedAt,
			&orderDetail.Order.UpdatedAt,
			&orderDetail.Customer.ID, &orderDetail.Customer.Prefix, &orderDetail.Customer.Name,
			&orderDetail.Customer.Address, &orderDetail.Customer.Phone, &orderDetail.Customer.ContactPerson,
			&orderDetail.Customer.Level, &orderDetail.Customer.DeliveryPlace, &orderDetail.Customer.Transport,
			&orderDetail.Customer.CreditLimit, &orderDetail.Customer.CreditTerm, &orderDetail.Customer.Outstanding,
			&orderDetail.Customer.LastContact, &orderDetail.Customer.Note, &orderDetail.Customer.CreatedAt,
			&orderDetail.Customer.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}

		// Get order items
		items, err := r.getOrderItems(orderDetail.Order.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get order items: %w", err)
		}
		orderDetail.Items = items

		orders = append(orders, orderDetail)
	}

	return orders, nil
}

func (r *orderRepository) UpdateStatus(id int, status string) error {
	query := `UPDATE orders SET status = $2 WHERE id = $1`

	result, err := r.db.Exec(query, id, status)
	if err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("order not found")
	}

	return nil
}

func (r *orderRepository) getOrderItems(orderID int) ([]models.OrderItemWithProduct, error) {
	query := `
		SELECT oi.id, oi.order_id, oi.product_id, oi.quantity, oi.unit,
		       oi.unit_price, oi.total_price, oi.note, oi.created_at,
		       p.name
		FROM order_items oi
		JOIN products p ON oi.product_id = p.id
		WHERE oi.order_id = $1
		ORDER BY oi.id
	`

	rows, err := r.db.Query(query, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order items: %w", err)
	}
	defer rows.Close()

	var items []models.OrderItemWithProduct
	for rows.Next() {
		var item models.OrderItemWithProduct
		err := rows.Scan(
			&item.OrderItem.ID, &item.OrderItem.OrderID, &item.OrderItem.ProductID,
			&item.OrderItem.Quantity, &item.OrderItem.Unit, &item.OrderItem.UnitPrice,
			&item.OrderItem.TotalPrice, &item.OrderItem.Note, &item.OrderItem.CreatedAt,
			&item.ProductName,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order item: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *orderRepository) generateOrderNumber(tx *sql.Tx) (string, error) {
	var count int
	err := tx.QueryRow("SELECT COUNT(*) FROM orders WHERE DATE(created_at) = CURRENT_DATE").Scan(&count)
	if err != nil {
		return "", fmt.Errorf("failed to get order count: %w", err)
	}

	return fmt.Sprintf("ORD-%s-%04d", strings.ReplaceAll(strings.Split(fmt.Sprintf("%v", sql.NullTime{}), " ")[0], "-", ""), count+1), nil
}
