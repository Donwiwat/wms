#!/bin/bash

echo "Generating Swagger documentation..."

# Navigate to backend directory
cd "$(dirname "$0")/.."

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