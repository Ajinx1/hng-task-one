package services

import (
	"hng-task-one/repositories"
)

type GeolocationService interface {
	GetLocation(ip string) (string, error)
}

type geolocationService struct {
	repo repositories.GeolocationRepository
}

func NewGeolocationService() GeolocationService {
	return &geolocationService{
		repo: repositories.NewGeolocationRepository(),
	}
}

func (s *geolocationService) GetLocation(ip string) (string, error) {
	return s.repo.GetLocation(ip)
}
