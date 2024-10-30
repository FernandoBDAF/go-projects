# Go Bookstore API

A RESTful API for managing books built with Go, using Gorilla Mux for routing and GORM with MySQL for data persistence.

## Features

- CRUD operations for books
- RESTful API endpoints
- MySQL database integration
- Structured project layout
- JSON response format

## Prerequisites

- Go 1.x
- MySQL 5.7+
- Docker (optional, for MySQL)

## Project Structure

```
.
├── cmd/
│   └── main/
│       └── main.go         # Application entry point
└── pkg/
    ├── config/
    │   └── app.go         # Database configuration
    ├── controllers/
    │   └── book-controller.go  # Request handlers
    ├── models/
    │   └── book.go        # Book model and database operations
    ├── routes/
    │   └── bookstore-routes.go # API routes definition
    └── utils/
        └── utils.go       # Utility functions
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/fbdaf/bookstore.git
cd bookstore
```

2. Install dependencies:
```bash
go mod download
```

3. Set up MySQL:
```bash
# Using Docker
docker run --name mysql-bookstore -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_DATABASE=bookstore -p 3306:3306 -d mysql:latest

# Or configure your existing MySQL instance and create a database named 'bookstore'
```

## Running the Application

1. Start the server:
```bash
go run cmd/main/main.go
```

The server will start on `localhost:8888`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/book` | Get all books |
| GET | `/book/{bookId}` | Get a specific book |
| POST | `/book` | Create a new book |
| PUT | `/book/{bookId}` | Update a book |
| DELETE | `/book/{bookId}` | Delete a book |

### Request Body Format (POST/PUT)

```json
{
    "name": "Book Name",
    "author": "Author Name",
    "publication": "Publication Name"
}
```

## Database Schema

The Book model includes:
- ID (auto-generated)
- Name
- Author
- Publication
- CreatedAt
- UpdatedAt
- DeletedAt

## Error Handling

The API includes basic error handling for:
- Invalid book IDs
- Database connection issues
- JSON parsing errors

## Development

The project uses:
- [Gorilla Mux](https://github.com/gorilla/mux) for routing
- [GORM](https://gorm.io) as ORM
- MySQL as the database