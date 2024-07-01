package services

import (
	"hng-task-one/repositories"
)

type WeatherService interface {
	GetTemperature(location string) (float64, error)
}

type weatherService struct {
	repo repositories.WeatherRepository
}

func NewWeatherService() WeatherService {
	return &weatherService{
		repo: repositories.NewWeatherRepository(),
	}
}

func (s *weatherService) GetTemperature(location string) (float64, error) {
	return s.repo.GetTemperature(location)
}
