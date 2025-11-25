# Warehouse Management System (WMS)

A production-grade warehouse management system built with Go backend and React frontend.

## Features

### Core Functionality
- **Product Management**: Add, edit, and manage products with multiple units and pricing
- **Warehouse Management**: Manage multiple warehouse locations
- **Stock Management**: Real-time stock tracking with dual-unit support
- **Stock Operations**: 
  - Stock In/Out operations
  - Inter-warehouse transfers
  - Break down (unit2 to unit1) and Pack up (unit1 to unit2)
  - Stock adjustments with reason tracking
- **Stock Movement History**: Complete audit trail of all stock movements
- **Multi-unit Support**: Handle products with two different units (e.g., pieces and boxes)

### Technical Features
- **Authentication & Authorization**: JWT-based secure authentication
- **Real-time Updates**: Live stock level updates
- **Responsive Design**: Mobile-friendly interface
- **Production Ready**: Proper error handling, logging, and validation
- **API Documentation**: Swagger/OpenAPI documentation
- **Database Migrations**: Automated database schema management

## Technology Stack

### Backend
- **Go 1.21+**: High-performance backend API
- **Gin Framework**: Fast HTTP web framework
- **PostgreSQL**: Robust relational database
- **JWT Authentication**: Secure token-based auth
- **Database Migrations**: golang-migrate for schema management
- **Swagger**: API documentation

### Frontend
- **React 18**: Modern React with hooks
- **TypeScript**: Type-safe development
- **Vite**: Fast build tool and dev server
- **Tailwind CSS**: Utility-first CSS framework
- **React Query**: Server state management
- **React Hook Form**: Form handling with validation
- **Zustand**: Lightweight state management
- **React Router**: Client-side routing

## Project Structure

```
wms/
├── backend/                 # Go backend
│   ├── cmd/server/         # Application entry point
│   ├── internal/           # Private application code
│   │   ├── config/         # Configuration management
│   │   ├── database/       # Database connection and migrations
│   │   ├── handlers/       # HTTP request handlers
│   │   ├── middleware/     # HTTP middleware
│   │   ├── models/         # Data models
│   │   ├── repositories/   # Data access layer
│   │   └── services/       # Business logic layer
│   ├── migrations/         # Database migration files
│   ├── go.mod             # Go module definition
│   └── .env.example       # Environment variables template
├── frontend/               # React frontend
│   ├── src/
│   │   ├── components/     # Reusable UI components
│   │   ├── pages/          # Page components
│   │   ├── services/       # API service layer
│   │   ├── stores/         # State management
│   │   ├── types/          # TypeScript type definitions
│   │   └── utils/          # Utility functions
│   ├── public/             # Static assets
│   └── package.json        # Node.js dependencies
└── README.md              # This file
```

## Getting Started

### Prerequisites
- Go 1.21 or higher
- Node.js 18 or higher
- PostgreSQL 12 or higher

### Backend Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd wms/backend
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Set up PostgreSQL database**
   ```sql
   CREATE DATABASE wms_db;
   CREATE USER wms_user WITH PASSWORD 'your_password';
   GRANT ALL PRIVILEGES ON DATABASE wms_db TO wms_user;
   ```

5. **Run database migrations**
   ```bash
   go run cmd/server/main.go
   # Migrations run automatically on startup
   ```

6. **Start the backend server**
   ```bash
   go run cmd/server/main.go
   ```

The backend API will be available at `http://localhost:8080`

### Frontend Setup

1. **Navigate to frontend directory**
   ```bash
   cd ../frontend
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Start the development server**
   ```bash
   npm run dev
   ```

The frontend will be available at `http://localhost:3000`

## API Documentation

Once the backend is running, you can access the Swagger API documentation at:
`http://localhost:8080/swagger/index.html`

### Key API Endpoints

- **Authentication**
  - `POST /api/v1/auth/login` - User login
  - `POST /api/v1/auth/register` - User registration

- **Products**
  - `GET /api/v1/products` - List all products
  - `POST /api/v1/products` - Create new product
  - `PUT /api/v1/products/{id}` - Update product
  - `DELETE /api/v1/products/{id}` - Delete product

- **Stock Operations**
  - `GET /api/v1/stock` - Get stock summary
  - `GET /api/v1/stock/card` - Get stock movement history
  - `POST /api/v1/stock/in` - Stock in operation
  - `POST /api/v1/stock/out` - Stock out operation
  - `POST /api/v1/stock/transfer` - Transfer between warehouses
  - `POST /api/v1/stock/adjust` - Stock adjustment

## Database Schema

The system uses the following main entities:

- **Products**: Product catalog with dual-unit support
- **Warehouses**: Storage locations
- **Stock**: Current stock levels per product/warehouse
- **Stock Movements**: Complete audit trail of all stock changes
- **Users**: System users with role-based access

## Development

### Running Tests
```bash
# Backend tests
cd backend
go test ./...

# Frontend tests
cd frontend
npm test
```

### Building for Production
```bash
# Backend
cd backend
go build -o wms-server cmd/server/main.go

# Frontend
cd frontend
npm run build
```

## Deployment

### Docker Deployment
```bash
# Build and run with Docker Compose
docker-compose up -d
```

### Manual Deployment
1. Build the backend binary
2. Build the frontend static files
3. Set up PostgreSQL database
4. Configure environment variables
5. Run database migrations
6. Start the services

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support and questions, please open an issue in the GitHub repository.