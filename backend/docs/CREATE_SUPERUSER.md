# Creating a Superuser

This document explains how to create a superuser (admin user) for the WMS (Warehouse Management System).

## Prerequisites

1. Ensure PostgreSQL is running and accessible
2. Database should be created and migrated
3. Environment variables should be properly configured in `.env` file

## Methods to Create Superuser

### Method 1: Using Batch Scripts (Windows)

#### Secure Mode (Recommended)
```bash
# Navigate to backend directory
cd backend

# Run the secure script (password will be hidden)
scripts\create-superuser.bat
```

#### Simple Mode
```bash
# Navigate to backend directory
cd backend

# Run the simple script (password will be visible)
scripts\create-superuser-simple.bat
```

#### Demo Mode (No Database Required)
```bash
# Navigate to backend directory
cd backend

# Run the demo script (shows what would be created without database)
scripts\create-superuser-demo.bat
```

### Method 2: Using PowerShell Scripts

```powershell
# Navigate to backend directory
cd backend

# Run the PowerShell script
.\scripts\create-superuser.ps1
```

### Method 3: Direct Go Command

#### Secure Mode (Hidden Password)
```bash
cd backend
go run cmd/createuser/main.go
```

#### Simple Mode (Visible Password)
```bash
cd backend
go run cmd/createuser/main.go -simple
```

#### Demo Mode (No Database Required)
```bash
cd backend
go run cmd/createuser/main.go -demo
```

### Method 4: Build and Run Executable

```bash
cd backend

# Build the executable
go build -o createuser.exe cmd/createuser/main.go

# Run the executable
./createuser.exe

# Or with simple mode
./createuser.exe -simple
```

## User Input Required

When you run the superuser creation tool, you'll be prompted for:

1. **Username**: Must be unique and at least 3 characters
2. **Email**: Must be unique and valid email format
3. **Password**: Must be at least 6 characters long
4. **Password Confirmation**: (Only in secure mode)

## User Roles

The superuser will be created with the following properties:
- **Role**: `admin` (gives full access to the system)
- **Status**: `active` (user can log in immediately)
- **Permissions**: Full access to all WMS features

## Example Usage

```bash
$ go run cmd/createuser/main.go

WMS Superuser Creation Tool (Secure Mode)
=========================================
Enter username: admin
Enter email: admin@company.com
Enter password: [hidden]
Confirm password: [hidden]

Superuser 'admin' created successfully!
User ID: 1
Email: admin@company.com
Role: admin
Created at: 2024-01-15 10:30:45
```

## Troubleshooting

### Database Connection Issues
- Verify PostgreSQL is running
- Check database credentials in `.env` file
- Ensure database exists and is accessible

### Username/Email Already Exists
- Choose a different username or email
- Check existing users in the database

### Password Requirements
- Password must be at least 6 characters long
- Use a strong password for security

### Permission Issues
- Ensure you have write access to the database
- Check database user permissions

## Security Notes

1. **Use Secure Mode**: Always use secure mode in production to hide password input
2. **Strong Passwords**: Use strong, unique passwords for admin accounts
3. **Limit Admin Users**: Only create admin users when necessary
4. **Regular Updates**: Regularly update admin passwords

## Database Schema

The user will be created in the `users` table with the following structure:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,  -- bcrypt hashed
    role VARCHAR(20) DEFAULT 'user',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Next Steps

After creating a superuser:
1. Test login functionality through the API
2. Use the admin account to create additional users if needed
3. Configure proper role-based access controls
4. Set up regular backup procedures for user data