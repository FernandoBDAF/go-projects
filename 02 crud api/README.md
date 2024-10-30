# Movie API

A simple RESTful API built with Go that simulates a movie database using in-memory storage with slices.

## Features
- CRUD operations for movies
- In-memory data storage using slices
- JSON response handling
- Random ID generation for new movies

## Data Models

### Movie
```go
type Movie struct {
    ID       string    `json:"id"`
    Isbn     string    `json:"isbn"`
    Title    string    `json:"title"`
    Director *Director `json:"director"`
}
```

### Director
```go
type Director struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
}
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | /movies  | Get all movies |
| GET    | /movies/{id} | Get a movie by ID |
| POST   | /movies  | Create a new movie |
| PUT    | /movies/{id} | Update a movie |
| DELETE | /movies/{id} | Delete a movie |

## Server Configuration
- Default Port: 8888

## Running the Application

1. Install dependencies:
```bash
go get -u github.com/gorilla/mux
```

2. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8888`

## Future Improvements

1. Add error handling
   - Proper HTTP status codes
   - Error response structures
   - Input validation

2. Add middleware
   - Request logging
   - CORS support
   - Request validation

3. Add data persistence
   - Database integration
   - Data validation

4. Improve security
   - Authentication
   - Rate limiting

5. Add testing
   - Unit tests
   - Integration tests

6. Add documentation
   - API documentation
   - Code comments