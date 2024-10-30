# Go Postgres Stock API

A RESTful API built with Go and PostgreSQL for managing stock information. This project demonstrates basic CRUD operations using Go's standard libraries and common packages.

## Features

- RESTful API endpoints for stock management
- PostgreSQL database integration
- CRUD operations (Create, Read, Update, Delete)
- Environment-based configuration
- Automatic database migration
- Error handling and logging
- JSON response formatting

## Prerequisites

- Go 1.x
- PostgreSQL
- [github.com/gorilla/mux](https://github.com/gorilla/mux) - For HTTP routing
- [github.com/lib/pq](https://github.com/lib/pq) - PostgreSQL driver
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - Environment configuration

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-postgres.git
cd go-postgres
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your environment variables:
```bash
cp .env.example .env
```
Edit the `.env` file with your PostgreSQL connection details:
```
POSTGRES_URL="postgresql://myuser:mypassword@localhost:5432/mydb?sslmode=disable"
```

## API Endpoints

| Method | Endpoint | Description | Request Body | Response |
|--------|----------|-------------|--------------|-----------|
| GET | `/api/stocks/{id}` | Get a single stock by ID | - | `{"stockid": 1, "name": "AAPL", "price": 150, "company": "Apple Inc."}` |
| GET | `/api/stock` | Get all stocks | - | `[{"stockid": 1, "name": "AAPL", "price": 150, "company": "Apple Inc."}, ...]` |
| POST | `/api/stock` | Create a new stock | `{"name": "AAPL", "price": 150, "company": "Apple Inc."}` | `{"id": 1, "message": "Stock created successfully"}` |
| PUT | `/api/stock/{id}` | Update a stock | `{"name": "AAPL", "price": 160, "company": "Apple Inc."}` | `{"id": 1, "message": "Stock updated successfully. Total rows/record affected 1"}` |
| DELETE | `/api/deletestock/{id}` | Delete a stock | - | `{"id": 1, "message": "Stock deleted successfully. Total rows/record affected 1"}` |

## Usage

1. Start the server:
```bash
go run main.go
```
The server will start on port 8888.

2. Example API calls:

Create a new stock:
```bash
curl -X POST http://localhost:8888/api/stock \
  -H "Content-Type: application/json" \
  -d '{"name":"AAPL","price":150,"company":"Apple Inc."}'
```

Get all stocks:
```bash
curl http://localhost:8888/api/stock
```

Get a specific stock:
```bash
curl http://localhost:8888/api/stocks/1
```

Update a stock:
```bash
curl -X PUT http://localhost:8888/api/stock/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"AAPL","price":160,"company":"Apple Inc."}'
```

Delete a stock:
```bash
curl -X DELETE http://localhost:8888/api/deletestock/1
```

## Project Structure

```
.
├── main.go           # Application entry point
├── router/
│   └── router.go     # Route definitions
├── middleware/
│   └── handlers.go   # Request handlers and database operations
├── models/
│   └── models.go     # Data models
├── .env              # Environment configuration
└── .env.example      # Example environment configuration
```

## Database Schema

The application automatically creates a `stocks` table with the following structure:

```sql
CREATE TABLE stocks (
    stockid SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    company VARCHAR(255) NOT NULL
);
```

## Error Handling

The API includes error handling for common scenarios:
- Database connection errors
- Invalid request parameters
- Missing environment variables
- Database query errors
- JSON encoding/decoding errors

## Development

To run the project in development mode:

1. Ensure PostgreSQL is running
2. Configure your `.env` file
3. Run the application:
```bash
go run main.go
```

For hot reloading during development, you can use tools like:
```bash
go install github.com/cosmtrek/air@latest
air
```

## Testing

To run the tests:
```bash
go test ./...
```

## Acknowledgments

- [Gorilla Mux](https://github.com/gorilla/mux) for HTTP routing
- [lib/pq](https://github.com/lib/pq) for PostgreSQL driver
- [godotenv](https://github.com/joho/godotenv) for environment configuration
```
