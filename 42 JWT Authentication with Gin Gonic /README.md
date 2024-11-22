# JWT Authentication API with Gin Gonic and MongoDB

A secure REST API implementation using Go (Gin Gonic framework) with JWT authentication and MongoDB as the database. This API provides user authentication, authorization, and user management functionalities.

## Tech Stack
- Go 1.21.4
- Gin Gonic Framework
- MongoDB
- JWT (JSON Web Tokens)
- Docker

## Project Structure

.
├── controllers/
│ └── user.go # User-related controllers (signup, login, etc.)
├── database/
│ └── dbConnection.go # MongoDB connection configuration
├── helpers/
│ ├── auth.go # Authentication helper functions
│ └── token.go # JWT token management
├── middleware/
│ └── auth.go # Authentication middleware
├── routes/
│ ├── auth.go # Authentication routes
│ └── user.go # User routes
├── main.go # Application entry point
└── Dockerfile # MongoDB container configuration

## API Endpoints

### Authentication Routes
1. **Sign Up**
   - Path: `/users/signup`
   - Method: POST
   - Description: Register a new user
   - Required fields: email, password, first_name, last_name, phone, user_type

2. **Login**
   - Path: `/users/login`
   - Method: POST
   - Description: Authenticate user and receive JWT tokens
   - Required fields: email, password

### Protected Routes
1. **Get All Users** (Admin only)
   - Path: `/users`
   - Method: GET
   - Description: Retrieve all users with pagination
   - Headers Required: token

2. **Get User**
   - Path: `/users/:user_id`
   - Method: GET
   - Description: Get specific user details
   - Headers Required: token

3. **API Test Endpoints**
   - `/api-1`: Protected route (Admin only)
   - `/api-2`: Protected route (All authenticated users)

## Security Features

### JWT Token System
The application implements a dual-token system:
- Access Token (12 hours expiration)
- Refresh Token (100 hours expiration)

### Password Security
- Passwords are securely hashed using bcrypt
- Password verification on login
- Secure password storage in MongoDB

### Middleware Authentication
- Token validation for protected routes
- Role-based access control
- Request validation

## Setup and Installation

### Prerequisites
- Go 1.21.4 or higher
- Docker
- MongoDB
- Git

### Environment Variables
Create a `.env` file in the root directory:

```env
PORT=8000
MONGO_URL=mongodb://localhost:27017
SECRET_KEY=your-secret-key
```

### Running with Docker

1. Start MongoDB container:
```bash
docker run -d -p 27017:27017 --name mongodb my-mongodb
```

2. Build and run the application:
```bash
go mod download
go run main.go
```

### Development Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd <project-directory>
```

2. Install dependencies:
```bash
go mod download
```

3. Set up VS Code debugging:
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Go JWT Debug",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/main.go",
            "env": {
                "PORT": "8888",
                "MONGO_URL": "mongodb://localhost:27017",
                "SECRET_KEY": "your-secret-key"
            },
            "buildFlags": "-gcflags='all=-N -l'",
            "showLog": true
        }
    ]
}
```

## API Usage

### Creating a New User
```bash
curl -X POST http://localhost:8000/users/signup \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "password": "securepassword",
    "phone": "1234567890",
    "user_type": "USER"
  }'
```

### Login
```bash
curl -X POST http://localhost:8000/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

### Accessing Protected Routes
```bash
curl -X GET http://localhost:8000/users \
  -H "token: YOUR_JWT_TOKEN"
```

## Error Handling
The API implements comprehensive error handling for:
- Invalid authentication
- Database errors
- Input validation
- Token validation
- Authorization errors

## Rate Limiting and Security
- JWT token expiration
- Role-based access control (USER/ADMIN)
- Request validation using gin validator
- Secure password hashing
- MongoDB security best practices

## Testing
To test the API endpoints:
1. Create a user using the signup endpoint
2. Login to receive JWT token
3. Use the token in the Authorization header for protected routes

## Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Acknowledgments
- Gin Gonic Framework
- MongoDB Go Driver
- JWT Go Implementation

