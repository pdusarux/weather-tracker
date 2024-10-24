# Weather API Service

A simple weather API service built with Go and Gin framework

## Features

- Fetch current weather data for a specified city
- Utilize WeatherAPI.com to retrieve up-to-date weather information
- Return data in JSON format

## Setup

To use this feature, you need to sign up for a free account at [WeatherAPI.com](https://www.weatherapi.com/my/) and obtain an API key. Once you have the key, add it to your .env file as:

```
API_KEY=your_api_key_here
```

## Installation

1. Clone this project:

   ```
   git clone https://github.com/pdusarux/weather-tracker.git
   cd weather-tracker
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Create a `.env` file in the project's root directory and add your API key:
   ```
   API_KEY=your_weatherapi_com_key_here
   ```

## Usage

1. Run the server:

   ```
   go run main.go
   ```

2. Call the API using the `/weather/:city` endpoint:
   ```
   http://localhost:8080/weather/London
   ```

## API Endpoint

- `GET /weather/:city`: Get current weather data for the specified city

  Example response:

  ```json
  {
    "message": "success",
    "data": {
      "current": {
        "temp_c": 15,
        "temp_f": 59,
        "condition": {
          "text": "Partly cloudy",
          "icon": "//cdn.weatherapi.com/weather/64x64/day/116.png"
        },
        "humidity": 68,
        "wind_kph": 17.6,
        "wind_mph": 10.9,
        "feelslike_c": 13.9,
        "feelslike_f": 57
      },
      "location": {
        "name": "London",
        "country": "United Kingdom",
        "lat": 51.52,
        "lon": -0.11,
        "localtime": "2023-04-20 15:30"
      }
    }
  }
  ```
