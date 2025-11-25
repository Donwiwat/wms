@echo off
echo Creating WMS Superuser (Demo Mode - No Database Required)...
echo.

cd /d "%~dp0\.."
go run cmd/createuser/main.go -demo

pause