@echo off
echo Generating Swagger documentation...

cd /d "%~dp0\.."

swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal

echo Swagger documentation generated successfully!
echo Files created:
echo   - docs/docs.go
echo   - docs/swagger.json
echo   - docs/swagger.yaml

if exist "docs\swagger.json" (
    echo ✅ swagger.json generated
) else (
    echo ❌ swagger.json not found
    exit /b 1
)

if exist "docs\swagger.yaml" (
    echo ✅ swagger.yaml generated
) else (
    echo ❌ swagger.yaml not found
    exit /b 1
)

echo 🎉 Documentation generation completed!