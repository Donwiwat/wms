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
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.Initialize(cfg.Database)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(cfg.Database); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize repositories
	repos := repositories.NewRepositories(db)

	// Initialize services
	services := services.NewServices(repos)

	// Initialize handlers
	handlers := handlers.NewHandlers(services)

	// Setup router
	router := setupRouter(cfg, handlers)

	// Start server
	log.Printf("Server starting on %s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Swagger UI available at: http://%s:%s/swagger/index.html", cfg.Server.Host, cfg.Server.Port)
	if err := router.Run(cfg.Server.Host + ":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRouter(cfg *config.Config, h *handlers.Handlers) *gin.Engine {
	// Set gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   cfg.CORS.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	router.Use(func(ctx *gin.Context) {
		c.HandlerFunc(ctx.Writer, ctx.Request)
		ctx.Next()
	})

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Swagger endpoint (only in development)
	if os.Getenv("GIN_MODE") != "release" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// API routes
	api := router.Group("/api/v1")
	{
		// Public routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.Auth.Login)
			auth.POST("/register", h.Auth.Register)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
		{
			// Products
			products := protected.Group("/products")
			{
				products.GET("", h.Product.GetProducts)
				products.GET("/:id", h.Product.GetProduct)
				products.POST("", h.Product.CreateProduct)
				products.PUT("/:id", h.Product.UpdateProduct)
				products.DELETE("/:id", h.Product.DeleteProduct)
				products.GET("/:id/prices", h.Product.GetProductPrices)
			}

			// Product Prices
			protected.POST("/product-prices", h.ProductPrice.CreateProductPrice)
			protected.PUT("/product-prices/:id", h.ProductPrice.UpdateProductPrice)
			protected.DELETE("/product-prices/:id", h.ProductPrice.DeleteProductPrice)

			// Customer Groups
			customerGroups := protected.Group("/customer-groups")
			{
				customerGroups.GET("", h.CustomerGroup.GetCustomerGroups)
				customerGroups.POST("", h.CustomerGroup.CreateCustomerGroup)
				customerGroups.PUT("/:id", h.CustomerGroup.UpdateCustomerGroup)
				customerGroups.DELETE("/:id", h.CustomerGroup.DeleteCustomerGroup)
			}

			// Customers
			customers := protected.Group("/customers")
			{
				customers.GET("", h.Customer.GetCustomers)
				customers.GET("/:id", h.Customer.GetCustomer)
				customers.POST("", h.Customer.CreateCustomer)
				customers.PUT("/:id", h.Customer.UpdateCustomer)
				customers.DELETE("/:id", h.Customer.DeleteCustomer)
				customers.GET("/search", h.Customer.SearchCustomers)
			}

			// Orders
			orders := protected.Group("/orders")
			{
				orders.GET("", h.Order.GetOrders)
				orders.GET("/:id", h.Order.GetOrder)
				orders.POST("", h.Order.CreateOrder)
				orders.PUT("/:id", h.Order.UpdateOrder)
				orders.DELETE("/:id", h.Order.DeleteOrder)
				orders.GET("/customer/:customerId", h.Order.GetOrdersByCustomer)
				orders.PATCH("/:id/status", h.Order.UpdateOrderStatus)
			}

			// Warehouses
			warehouses := protected.Group("/warehouses")
			{
				warehouses.GET("", h.Warehouse.GetWarehouses)
				warehouses.GET("/:id", h.Warehouse.GetWarehouse)
				warehouses.POST("", h.Warehouse.CreateWarehouse)
				warehouses.PUT("/:id", h.Warehouse.UpdateWarehouse)
				warehouses.DELETE("/:id", h.Warehouse.DeleteWarehouse)
			}

			// Stock
			stock := protected.Group("/stock")
			{
				stock.GET("", h.Stock.GetStock)
				stock.GET("/card", h.Stock.GetStockCard)
				stock.POST("/in", h.Stock.StockIn)
				stock.POST("/out", h.Stock.StockOut)
				stock.POST("/break", h.Stock.BreakDown)
				stock.POST("/pack", h.Stock.PackUp)
				stock.POST("/transfer", h.Stock.Transfer)
				stock.POST("/adjust", h.Stock.Adjust)
			}

			// Stock Movements
			protected.GET("/stock-movements", h.StockMovement.GetStockMovements)

			// Documents
			salesOrders := protected.Group("/sales-orders")
			{
				salesOrders.GET("", h.SalesOrder.GetSalesOrders)
				salesOrders.POST("", h.SalesOrder.CreateSalesOrder)
				salesOrders.PUT("/:id", h.SalesOrder.UpdateSalesOrder)
				salesOrders.DELETE("/:id", h.SalesOrder.DeleteSalesOrder)
			}

			deliveryOrders := protected.Group("/delivery-orders")
			{
				deliveryOrders.GET("", h.DeliveryOrder.GetDeliveryOrders)
				deliveryOrders.POST("", h.DeliveryOrder.CreateDeliveryOrder)
				deliveryOrders.PUT("/:id", h.DeliveryOrder.UpdateDeliveryOrder)
				deliveryOrders.DELETE("/:id", h.DeliveryOrder.DeleteDeliveryOrder)
			}

			purchaseOrders := protected.Group("/purchase-orders")
			{
				purchaseOrders.GET("", h.PurchaseOrder.GetPurchaseOrders)
				purchaseOrders.POST("", h.PurchaseOrder.CreatePurchaseOrder)
				purchaseOrders.PUT("/:id", h.PurchaseOrder.UpdatePurchaseOrder)
				purchaseOrders.DELETE("/:id", h.PurchaseOrder.DeletePurchaseOrder)
			}

			goodsReceipts := protected.Group("/goods-receipts")
			{
				goodsReceipts.GET("", h.GoodsReceipt.GetGoodsReceipts)
				goodsReceipts.POST("", h.GoodsReceipt.CreateGoodsReceipt)
				goodsReceipts.PUT("/:id", h.GoodsReceipt.UpdateGoodsReceipt)
				goodsReceipts.DELETE("/:id", h.GoodsReceipt.DeleteGoodsReceipt)
			}

			transfers := protected.Group("/transfers")
			{
				transfers.GET("", h.Transfer.GetTransfers)
				transfers.POST("", h.Transfer.CreateTransfer)
				transfers.PUT("/:id", h.Transfer.UpdateTransfer)
				transfers.DELETE("/:id", h.Transfer.DeleteTransfer)
			}

			stockAdjustments := protected.Group("/stock-adjustments")
			{
				stockAdjustments.GET("", h.StockAdjustment.GetStockAdjustments)
				stockAdjustments.POST("", h.StockAdjustment.CreateStockAdjustment)
				stockAdjustments.PUT("/:id", h.StockAdjustment.UpdateStockAdjustment)
				stockAdjustments.DELETE("/:id", h.StockAdjustment.DeleteStockAdjustment)
			}
		}
	}

	return router
}
