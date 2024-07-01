package handlers

import (
	"encoding/json"
	"fmt"
	"hng-task1/schemas"
	"hng-task1/services"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	visitorName := r.URL.Query().Get("visitor_name")
	if visitorName == "" {
		visitorName = "Guest"
	}

	var errResponse schemas.ErrorResponse

	ipService := services.NewIPService()
	clientIP := ipService.GetClientIP(r)

	locationService := services.NewGeolocationService()
	location, err := locationService.GetLocation(clientIP)
	if err != nil {
		errResponse = schemas.ErrorResponse{
			ErrorMsg: "Unable to get user location",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	weatherService := services.NewWeatherService()
	temperature, err := weatherService.GetTemperature(location)
	if err != nil {
		errResponse = schemas.ErrorResponse{
			ErrorMsg: "Unable to get the current weather condition",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	response := schemas.HelloResponse{
		ClientIP: clientIP,
		Location: location,
		Greeting: fmt.Sprintf("Hello, %s!, the temperature is %.2f degrees Celsius in %s", visitorName, temperature, location),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
