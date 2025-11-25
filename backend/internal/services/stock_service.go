package services

import (
	"errors"
	"fmt"
	"time"

	"wms-backend/internal/models"
	"wms-backend/internal/repositories"
)

// StockService interface defines stock service methods
type StockService interface {
	GetStock(productID, warehouseID int) (*models.Stock, error)
	GetStockSummary(productID, warehouseID int) ([]*models.StockSummary, error)
	GetStockCard(productID, warehouseID int) ([]*models.StockCardEntry, error)
	StockIn(req *models.StockInRequest, username string) error
	StockOut(req *models.StockOutRequest, username string) error
	BreakDown(req *models.BreakDownRequest, username string) error
	PackUp(req *models.PackUpRequest, username string) error
	Transfer(req *models.TransferRequest, username string) error
	Adjust(req *models.StockAdjustRequest, username string) error
}

// stockService implements StockService
type stockService struct {
	stockRepo         repositories.StockRepository
	stockMovementRepo repositories.StockMovementRepository
	productRepo       repositories.ProductRepository
}

// NewStockService creates a new stock service
func NewStockService(
	stockRepo repositories.StockRepository,
	stockMovementRepo repositories.StockMovementRepository,
	productRepo repositories.ProductRepository,
) StockService {
	return &stockService{
		stockRepo:         stockRepo,
		stockMovementRepo: stockMovementRepo,
		productRepo:       productRepo,
	}
}

// GetStock gets stock for a specific product and warehouse
func (s *stockService) GetStock(productID, warehouseID int) (*models.Stock, error) {
	return s.stockRepo.GetStock(productID, warehouseID)
}

// GetStockSummary gets stock summary with product and warehouse details
func (s *stockService) GetStockSummary(productID, warehouseID int) ([]*models.StockSummary, error) {
	return s.stockRepo.GetStockSummary(productID, warehouseID)
}

// GetStockCard gets stock movement history
func (s *stockService) GetStockCard(productID, warehouseID int) ([]*models.StockCardEntry, error) {
	return s.stockRepo.GetStockCard(productID, warehouseID)
}

// StockIn processes stock in operation
func (s *stockService) StockIn(req *models.StockInRequest, username string) error {
	// Get product to check units and ratio
	product, err := s.productRepo.GetByID(req.ProductID)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	// Calculate quantities based on unit
	var remain1Delta, remain2Delta int
	if req.Unit == product.Unit1 {
		remain1Delta = int(req.Qty)
	} else if req.Unit == product.Unit2 {
		remain2Delta = int(req.Qty)
	} else {
		return errors.New("invalid unit for this product")
	}

	// Get current stock or create if doesn't exist
	currentStock, err := s.stockRepo.GetStock(req.ProductID, req.WarehouseID)
	if err != nil {
		// Create new stock record
		if err := s.stockRepo.CreateOrUpdateStock(req.ProductID, req.WarehouseID, remain1Delta, remain2Delta); err != nil {
			return err
		}
	} else {
		// Update existing stock
		newRemain1 := currentStock.Remain1 + remain1Delta
		newRemain2 := currentStock.Remain2 + remain2Delta
		if err := s.stockRepo.CreateOrUpdateStock(req.ProductID, req.WarehouseID, newRemain1, newRemain2); err != nil {
			return err
		}
	}

	// Create stock movement record
	movement := &models.StockMovement{
		ProductID:   req.ProductID,
		WarehouseID: req.WarehouseID,
		Type:        "IN",
		Qty:         req.Qty,
		Unit:        req.Unit,
		RefID:       req.RefID,
		RefType:     req.RefType,
		Date:        time.Now(),
		Note:        req.Note,
		CreatedBy:   username,
	}

	return s.stockMovementRepo.Create(movement)
}

// StockOut processes stock out operation
func (s *stockService) StockOut(req *models.StockOutRequest, username string) error {
	// Get product to check units and ratio
	product, err := s.productRepo.GetByID(req.ProductID)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	// Get current stock
	currentStock, err := s.stockRepo.GetStock(req.ProductID, req.WarehouseID)
	if err != nil {
		return errors.New("stock not found")
	}

	// Calculate quantities based on unit
	var remain1Delta, remain2Delta int
	if req.Unit == product.Unit1 {
		remain1Delta = int(req.Qty)
		if currentStock.Remain1 < remain1Delta {
			return errors.New("insufficient stock in unit1")
		}
	} else if req.Unit == product.Unit2 {
		remain2Delta = int(req.Qty)
		if currentStock.Remain2 < remain2Delta {
			return errors.New("insufficient stock in unit2")
		}
	} else {
		return errors.New("invalid unit for this product")
	}

	// Update stock
	newRemain1 := currentStock.Remain1 - remain1Delta
	newRemain2 := currentStock.Remain2 - remain2Delta
	if err := s.stockRepo.CreateOrUpdateStock(req.ProductID, req.WarehouseID, newRemain1, newRemain2); err != nil {
		return err
	}

	// Create stock movement record
	movement := &models.StockMovement{
		ProductID:   req.ProductID,
		WarehouseID: req.WarehouseID,
		Type:        "OUT",
		Qty:         req.Qty,
		Unit:        req.Unit,
		RefID:       req.RefID,
		RefType:     req.RefType,
		Date:        time.Now(),
		Note:        req.Note,
		CreatedBy:   username,
	}

	return s.stockMovementRepo.Create(movement)
}

// BreakDown processes break down operation (unit2 to unit1)
func (s *stockService) BreakDown(req *models.BreakDownRequest, username string) error {
	// Get product to check ratio
	product, err := s.productRepo.GetByID(req.ProductID)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	if product.Ratio <= 0 {
		return errors.New("product ratio not configured for break down")
	}

	// Get current stock
	currentStock, err := s.stockRepo.GetStock(req.ProductID, req.WarehouseID)
	if err != nil {
		return errors.New("stock not found")
	}

	// Check if enough unit2 stock
	if currentStock.Remain2 < req.QtyUnit2 {
		return errors.New("insufficient stock in unit2 for break down")
	}

	// Calculate new quantities
	newRemain1 := currentStock.Remain1 + int(product.Ratio*float64(req.QtyUnit2))
	newRemain2 := currentStock.Remain2 - req.QtyUnit2

	// Update stock
	if err := s.stockRepo.CreateOrUpdateStock(req.ProductID, req.WarehouseID, newRemain1, newRemain2); err != nil {
		return err
	}

	// Create stock movement record
	movement := &models.StockMovement{
		ProductID:   req.ProductID,
		WarehouseID: req.WarehouseID,
		Type:        "BREAK",
		Qty:         float64(req.QtyUnit2),
		Unit:        product.Unit2,
		Date:        time.Now(),
		Note:        req.Note,
		CreatedBy:   username,
	}

	return s.stockMovementRepo.Create(movement)
}

// PackUp processes pack up operation (unit1 to unit2)
func (s *stockService) PackUp(req *models.PackUpRequest, username string) error {
	// Get product to check ratio
	product, err := s.productRepo.GetByID(req.ProductID)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	if product.Ratio <= 0 {
		return errors.New("product ratio not configured for pack up")
	}

	// Get current stock
	currentStock, err := s.stockRepo.GetStock(req.ProductID, req.WarehouseID)
	if err != nil {
		return errors.New("stock not found")
	}

	// Calculate required unit1 quantity
	requiredUnit1 := int(product.Ratio * float64(req.QtyUnit2))

	// Check if enough unit1 stock
	if currentStock.Remain1 < requiredUnit1 {
		return errors.New("insufficient stock in unit1 for pack up")
	}

	// Calculate new quantities
	newRemain1 := currentStock.Remain1 - requiredUnit1
	newRemain2 := currentStock.Remain2 + req.QtyUnit2

	// Update stock
	if err := s.stockRepo.CreateOrUpdateStock(req.ProductID, req.WarehouseID, newRemain1, newRemain2); err != nil {
		return err
	}

	// Create stock movement record
	movement := &models.StockMovement{
		ProductID:   req.ProductID,
		WarehouseID: req.WarehouseID,
		Type:        "PACK",
		Qty:         float64(req.QtyUnit2),
		Unit:        product.Unit2,
		Date:        time.Now(),
		Note:        req.Note,
		CreatedBy:   username,
	}

	return s.stockMovementRepo.Create(movement)
}

// Transfer processes transfer operation between warehouses
func (s *stockService) Transfer(req *models.TransferRequest, username string) error {
	// Get product to check units
	product, err := s.productRepo.GetByID(req.ProductID)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	// Validate unit
	if req.Unit != product.Unit1 && req.Unit != product.Unit2 {
		return errors.New("invalid unit for this product")
	}

	// Process stock out from source warehouse
	stockOutReq := &models.StockOutRequest{
		ProductID:   req.ProductID,
		WarehouseID: req.FromWarehouseID,
		Qty:         req.Qty,
		Unit:        req.Unit,
		RefID:       req.RefID,
		RefType:     req.RefType,
		Note:        req.Note,
	}

	if err := s.StockOut(stockOutReq, username); err != nil {
		return fmt.Errorf("transfer out failed: %w", err)
	}

	// Process stock in to destination warehouse
	stockInReq := &models.StockInRequest{
		ProductID:   req.ProductID,
		WarehouseID: req.ToWarehouseID,
		Qty:         req.Qty,
		Unit:        req.Unit,
		RefID:       req.RefID,
		RefType:     req.RefType,
		Note:        req.Note,
	}

	if err := s.StockIn(stockInReq, username); err != nil {
		// TODO: Rollback the stock out operation
		return fmt.Errorf("transfer in failed: %w", err)
	}

	return nil
}

// Adjust processes stock adjustment
func (s *stockService) Adjust(req *models.StockAdjustRequest, username string) error {
	// Get current stock
	currentStock, err := s.stockRepo.GetStock(req.ProductID, req.WarehouseID)
	if err != nil {
		return errors.New("stock not found")
	}

	// Calculate differences
	diff1 := req.NewRemain1 - currentStock.Remain1
	diff2 := req.NewRemain2 - currentStock.Remain2

	// Update stock
	if err := s.stockRepo.CreateOrUpdateStock(req.ProductID, req.WarehouseID, req.NewRemain1, req.NewRemain2); err != nil {
		return err
	}

	// Create stock movement records for adjustments
	if diff1 != 0 {
		product, _ := s.productRepo.GetByID(req.ProductID)
		movement := &models.StockMovement{
			ProductID:   req.ProductID,
			WarehouseID: req.WarehouseID,
			Type:        "ADJUST",
			Qty:         float64(diff1),
			Unit:        product.Unit1,
			Date:        time.Now(),
			Note:        fmt.Sprintf("%s - %s", req.Reason, req.Note),
			CreatedBy:   username,
		}
		s.stockMovementRepo.Create(movement)
	}

	if diff2 != 0 {
		product, _ := s.productRepo.GetByID(req.ProductID)
		movement := &models.StockMovement{
			ProductID:   req.ProductID,
			WarehouseID: req.WarehouseID,
			Type:        "ADJUST",
			Qty:         float64(diff2),
			Unit:        product.Unit2,
			Date:        time.Now(),
			Note:        fmt.Sprintf("%s - %s", req.Reason, req.Note),
			CreatedBy:   username,
		}
		s.stockMovementRepo.Create(movement)
	}

	return nil
}
