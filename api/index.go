package handlers

import (
	"encoding/json"
	"fmt"
	"hng-task-one/schemas"
	"hng-task-one/services"
	"net/http"
	"strings"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	visitorName := r.URL.Query().Get("visitor_name")
	if visitorName == "" {
		visitorName = "Guest"
	} else {
		visitorName = strings.Trim(visitorName, `"`)
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

	greeting := fmt.Sprintf("Hello, %s!, the temperature is %.0f degrees Celsius in %s", visitorName, temperature, location)

	response := schemas.HelloResponse{
		ClientIP: clientIP,
		Location: location,
		Greeting: greeting,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
