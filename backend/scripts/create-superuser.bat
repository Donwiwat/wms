@echo off
echo Creating WMS Superuser...
echo.

cd /d "%~dp0\.."
go run cmd/createuser/main.go

pause