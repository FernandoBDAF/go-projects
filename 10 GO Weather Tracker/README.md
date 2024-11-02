# GO Weather Tracker

A simple weather tracking API built with Go that fetches real-time weather data from OpenWeatherMap API.

## Features

- 🌡️ Real-time weather data retrieval
- 🌍 City-based weather lookup
- 🚀 Simple HTTP endpoints
- 🔑 API key configuration
- 🌐 Temperature in Kelvin

## Prerequisites

- Go 1.x
- OpenWeatherMap API key
- Basic understanding of HTTP requests

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/go-weather-tracker.git
    ```
2. Navigate to the project directory:
    ```bash
    cd go-weather-tracker
    ```
3. Set up OpenWeatherMap API key:
   - Sign up at [OpenWeatherMap](https://openweathermap.org/)
   - Get your API key
   - Create `.apiConfig` file in project root:
    ```json
    {
        "OpenWeatherMapApiKey": "your_api_key_here"
    }
    ```

## API Endpoints

### 1. Health Check


GET /hello

Returns a simple "Hello, World!" message.

### 2. Weather Lookup

GET /weather/{city}

Returns current weather data for the specified city.

Examples:
- Single word city: `/weather/london`
- Multi-word city: `/weather/san+francisco`

## Usage

1. Start the server:
```bash
go run main.go
```

2. The server will start on port 8888

3. Access the endpoints:
- http://localhost:8888/hello
- http://localhost:8888/weather/london

## Project Structure

```
├── main.go
├── .apiConfig
├── README.md
└── .gitignore
```

## Data Structures

The weather data is structured as follows:
```go
startLine: 21
endLine: 30
```

## Configuration

The API configuration is loaded from `.apiConfig` file:
```go
startLine: 32
endLine: 45
```

## Security Notes

- Never commit your `.apiConfig` file to version control
- Add `.apiConfig` to your `.gitignore`
- Rotate your API key if exposed
- Consider using environment variables for production

## Error Handling

The application includes error handling for:
- Invalid city names
- API configuration loading errors
- HTTP request errors
- JSON parsing errors
