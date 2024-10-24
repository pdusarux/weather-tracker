package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type WeatherData struct {
	Current struct {
		Temp_c    float64 `json:"temp_c"`
		Temp_f    float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
		Humidity    int     `json:"humidity"`
		Wind_kph    float64 `json:"wind_kph"`
		Wind_mph    float64 `json:"wind_mph"`
		Feelslike_c float64 `json:"feelslike_c"`
		Feelslike_f float64 `json:"feelslike_f"`
	} `json:"current"`
	Location struct {
		Name      string  `json:"name"`
		Country   string  `json:"country"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		Localtime string  `json:"localtime"`
	} `json:"location"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	apiKey := os.Getenv("API_KEY")

	r.GET("/weather/:city", func(c *gin.Context) {
		city := c.Param("city")

		url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, city)

		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}

		if resp.StatusCode != http.StatusOK {
			c.JSON(resp.StatusCode, gin.H{
				"error":    "API request failed",
				"status":   resp.Status,
				"response": string(body),
			})
			return
		}

		var weatherData WeatherData
		if err := json.Unmarshal(body, &weatherData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":    "Failed to parse weather data",
				"response": string(body),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    weatherData,
		})
	})

	r.Run()
}
