.PHONY: docs serve-docs clean-docs validate-docs install-tools check

# Generate Swagger documentation
docs:
	@echo "Generating Swagger documentation..."
	swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
	@echo "Documentation generated successfully!"

# Serve documentation in development
serve-docs: docs
	@echo "Starting server with Swagger UI..."
	@echo "Swagger UI will be available at: http://localhost:8080/swagger/index.html"
	set GIN_MODE=debug && go run cmd/server/main.go

# Clean generated documentation
clean-docs:
	@echo "Cleaning generated documentation..."
	@if exist docs rmdir /s /q docs
	@echo "Documentation cleaned!"

# Validate generated documentation (requires swagger CLI)
validate-docs: docs
	@echo "Validating Swagger documentation..."
	@echo "Note: Install swagger CLI with 'go install github.com/go-swagger/go-swagger/cmd/swagger@latest'"
	swagger validate docs/swagger.yaml
	@echo "Documentation is valid!"

# Install required tools
install-tools:
	@echo "Installing Swagger tools..."
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	@echo "Tools installed successfully!"

# Run all checks
check: clean-docs docs
	@echo "All documentation checks passed!"

# Build the application
build:
	@echo "Building WMS API..."
	go build -o bin/wms-api cmd/server/main.go
	@echo "Build completed!"

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Development server with live reload
dev: docs
	@echo "Starting development server..."
	@echo "API: http://localhost:8080"
	@echo "Swagger UI: http://localhost:8080/swagger/index.html"
	@echo "Health Check: http://localhost:8080/health"
	set GIN_MODE=debug && go run cmd/server/main.go

# Help
help:
	@echo "Available commands:"
	@echo "  docs         - Generate Swagger documentation"
	@echo "  serve-docs   - Start server with Swagger UI"
	@echo "  clean-docs   - Clean generated documentation"
	@echo "  validate-docs- Validate Swagger documentation"
	@echo "  install-tools- Install required tools"
	@echo "  check        - Run all documentation checks"
	@echo "  build        - Build the application"
	@echo "  test         - Run tests"
	@echo "  dev          - Start development server"
	@echo "  help         - Show this help message"