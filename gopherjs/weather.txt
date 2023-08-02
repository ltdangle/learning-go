package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type WeatherResponse struct {
	City     string  `json:"city"`
	Temp     float64 `json:"temp"`
	Wind     float64 `json:"wind"`
	Humidity float64 `json:"humidity"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// Prepare parameters.
	apiKey := os.Getenv("WEATHER_API_KEY")
	city := r.URL.Query().Get("city")

	// Call api.
	var apiResponse string
	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1", apiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		respByte, _ := io.ReadAll(resp.Body)
		apiResponse = string(respByte)
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(apiResponse), &result)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Convert api result into our json response.
	var weather = WeatherResponse{
		Temp:     result["current"].(map[string]interface{})["temp_c"].(float64),
		Wind:     result["current"].(map[string]interface{})["wind_kph"].(float64),
		Humidity: result["current"].(map[string]interface{})["humidity"].(float64),
	}

	json.NewEncoder(w).Encode(weather)
}

func main() {
	http.HandleFunc("/weather", weatherHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
