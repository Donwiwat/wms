Write-Host "Creating WMS Superuser..." -ForegroundColor Green
Write-Host ""

# Change to backend directory
Set-Location -Path (Split-Path -Parent $PSScriptRoot)

# Run the superuser creation tool
go run cmd/createuser/main.go

Write-Host ""
Write-Host "Press any key to continue..." -ForegroundColor Yellow
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")