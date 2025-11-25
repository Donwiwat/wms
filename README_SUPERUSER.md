# How to Create a Superuser for WMS

This guide explains how to create a superuser (admin user) for the Warehouse Management System.

## Quick Start

### Option 1: Demo Mode (No Database Required)
Perfect for testing or understanding how the system works:

```bash
cd backend
go run cmd/createuser/main.go -demo
```

Or use the batch script:
```bash
cd backend
scripts\create-superuser-demo.bat
```

### Option 2: Simple Mode (Database Required)
Creates an actual user in the database with visible password input:

```bash
cd backend
go run cmd/createuser/main.go -simple
```

Or use the batch script:
```bash
cd backend
scripts\create-superuser-simple.bat
```

### Option 3: Secure Mode (Database Required)
Creates an actual user in the database with hidden password input:

```bash
cd backend
go run cmd/createuser/main.go
```

Or use the batch script:
```bash
cd backend
scripts\create-superuser.bat
```

## Prerequisites for Database Modes

1. **PostgreSQL Running**: Ensure PostgreSQL is installed and running
2. **Database Created**: The `wms_db` database should exist
3. **Environment Variables**: Configure `.env` file with correct database credentials
4. **Migrations Applied**: Run database migrations first

## Environment Setup

1. Copy the environment file:
   ```bash
   cd backend
   cp .env.example .env
   ```

2. Update database credentials in `.env`:
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=wms_db
   DB_SSLMODE=disable
   ```

3. Run database migrations:
   ```bash
   cd backend
   go run cmd/migrate/main.go up
   ```

## Superuser Details

The created superuser will have:
- **Role**: `admin` (full system access)
- **Status**: `active` (can login immediately)
- **Password**: Securely hashed using bcrypt
- **Permissions**: Full access to all WMS features

## Example Usage

```bash
$ cd backend
$ go run cmd/createuser/main.go -demo

WMS Superuser Creation Tool (Demo Mode)
======================================
This demo shows how the superuser creation works without database connection

Enter username: admin
Enter email: admin@company.com
Enter password: securepassword123

=== DEMO: Superuser would be created with these details ===
Username: admin
Email: admin@company.com
Password Hash: $2a$10$...
Role: admin
Is Active: true

=== SQL that would be executed ===
INSERT INTO users (username, email, password, role, is_active)
VALUES ('admin', 'admin@company.com', '$2a$10$...', 'admin', true);

Demo completed successfully!
```

## Troubleshooting

### Database Connection Failed
- Check if PostgreSQL is running
- Verify database credentials in `.env`
- Ensure database `wms_db` exists

### Username/Email Already Exists
- Choose different username or email
- Check existing users in database

### Password Requirements
- Minimum 6 characters required
- Use strong passwords for security

## Security Best Practices

1. **Use Secure Mode** in production environments
2. **Strong Passwords** - use complex, unique passwords
3. **Limit Admin Users** - only create when necessary
4. **Regular Updates** - change passwords periodically
5. **Monitor Access** - track admin user activities

## Next Steps

After creating a superuser:

1. **Test Login**: Verify login through the API
2. **Create Additional Users**: Use admin account to create other users
3. **Configure Roles**: Set up proper role-based access controls
4. **Setup Backup**: Configure regular database backups

For detailed documentation, see: `backend/docs/CREATE_SUPERUSER.md`