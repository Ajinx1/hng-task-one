package repositories

import (
	"encoding/json"
	"fmt"
	"hng-task-one/schemas"
	"net/http"
	"os"
)

type GeolocationRepository interface {
	GetLocation(ip string) (string, error)
}

type geolocationRepository struct{}

func NewGeolocationRepository() GeolocationRepository {
	return &geolocationRepository{}
}

func (r *geolocationRepository) GetLocation(ip string) (string, error) {
	apiKey := os.Getenv("IPSTACK_API_KEY")

	if apiKey == "" {
		return "", fmt.Errorf("IPSTACK_API_KEY is not set")
	}

	resp, err := http.Get(fmt.Sprintf("https://api.ipgeolocation.io/ipgeo?apiKey=%v&ip=%v", apiKey, ip))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result schemas.LocationResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.City, nil
}
