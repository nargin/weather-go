# Weather API in Go

A simple and efficient weather API service built with Go that provides real-time weather data and forecasts.

## Features

- Real-time weather data for any location
- 7-day weather forecasts
- Current conditions (temperature, humidity, wind, pressure)
- UV index and visibility information
- RESTful API endpoints
- JSON response format

## Installation

```bash
git clone https://github.com/yourusername/weather-go.git
cd weather-go
go mod download
```

## Configuration

Create a `.env` file in the project root:

```
WEATHER_API_KEY=your_api_key_here
PORT=8080
```

## Usage

Start the server:

```bash
go run main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Get Current Weather
```
GET /api/weather/current?city={city}&country={country}
```

### Get Weather Forecast
```
GET /api/weather/forecast?city={city}&country={country}&days={1-7}
```

### Get Weather by Coordinates
```
GET /api/weather/coordinates?lat={latitude}&lon={longitude}
```

## Example Response

```json
{
  "location": {
    "city": "London",
    "country": "UK",
    "latitude": 51.5074,
    "longitude": -0.1278
  },
  "current": {
    "temperature": 18.5,
    "feels_like": 17.2,
    "humidity": 65,
    "condition": "Partly Cloudy",
    "wind_speed": 12.5
  }
}
```

## Technologies Used

- Go 1.21+
- Gin Web Framework
- Weather data from OpenWeatherMap API

## License

MIT
