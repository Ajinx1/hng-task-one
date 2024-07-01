package repositories

import (
	"encoding/json"
	"fmt"
	"hng-task1/schemas"
	"net/http"
	"os"
)

type WeatherRepository interface {
	GetTemperature(location string) (float64, error)
}

type weatherRepository struct{}

func NewWeatherRepository() WeatherRepository {
	return &weatherRepository{}
}

func (r *weatherRepository) GetTemperature(location string) (float64, error) {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")

	if apiKey == "" {
		return 0, fmt.Errorf("OPENWEATHERMAP_API_KEY is not set")
	}

	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", location, apiKey))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result schemas.WeatherResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Main.Temp, nil
}
