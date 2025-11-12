# Warehouse Control Service

A robust REST API service built in Go for managing warehouse inventory and user authentication. This service provides secure endpoints for item management with role-based access control and comprehensive audit logging.

## Features

- **User Authentication**: Secure login and registration with JWT-based authentication
- **Role-Based Access Control**: Different permission levels (admin, manager, viewer)
- **Item Management**: Full CRUD operations for warehouse items
- **Audit Logging**: Automatic tracking of all item changes in the database for internal analytics
- **PostgreSQL Integration**: Reliable database storage with connection pooling
- **Docker Support**: Containerized deployment with Docker Compose

## API Endpoints

### Authentication
- `POST /warehouse-control/api/auth/login` - User login
- `POST /warehouse-control/api/auth/register` - User registration (admin only)

### Items (Requires Authentication)
- `POST /warehouse-control/api/items/` - Create new item
- `GET /warehouse-control/api/items/` - Retrieve all items
- `PUT /warehouse-control/api/items/:id` - Update existing item
- `DELETE /warehouse-control/api/items/:id` - Delete item

## Quick Start

### Prerequisites
- Go 1.25.3 or later
- Docker and Docker Compose
- PostgreSQL (handled via Docker)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/avraam311/warehouse-control.git
cd warehouse-control
```

2. Create a `.env` file with your database configuration:
```env
DB_HOST=db
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_SSL_MODE=disable
JWT_SECRET=your_jwt_secret
GOOSE_DRIVER=postgres
GOOSE_MIGRATION_DIR=/migrations
```

3. Start the services:
```bash
make up
```

This will build and start the application along with PostgreSQL database and run migrations.

### Alternative: Build and Run Locally

1. Ensure PostgreSQL is running locally or adjust connection settings
2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/main.go
```

## Configuration

The service uses a YAML configuration file (`config/local.yaml`) for server and database settings. Environment variables can override these settings.

Key configuration options:
- Server port (default: 8080)
- Database connection pool settings
- Authentication roles

## Database Schema

The service uses PostgreSQL with the following main tables:
- `users` - User accounts with roles
- `item` - Warehouse items
- `item_history` - Audit trail of all item changes

All item modifications are automatically logged to the `item_history` table for internal analytics and compliance purposes. Note that while the history is maintained in the database, there are no public API endpoints for retrieving this historical data - it is used solely for internal analysis.

## Development

### Available Commands
- `make lint` - Run code linting and vetting
- `make up` - Start services with Docker Compose
- `make buildup` - Rebuild and start services
- `make down` - Stop services and remove volumes

### Project Structure
```
cmd/                    # Application entry point
config/                 # Configuration files
internal/
  api/                  # HTTP handlers and server setup
  models/               # Data structures
  repository/           # Database layer
  service/              # Business logic
migrations/             # Database migrations
```

## Security

- JWT-based authentication for API access
- Role-based permissions for different user types
- CORS middleware for cross-origin requests
- Input validation using struct tags

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request
