package services

import (
	"hng-task-one/repositories"
	"net/http"
)

type IPService interface {
	GetClientIP(r *http.Request) string
}

type ipService struct {
	repo repositories.IPRepository
}

func NewIPService() IPService {
	return &ipService{
		repo: repositories.NewIPRepository(),
	}
}

func (s *ipService) GetClientIP(r *http.Request) string {
	return s.repo.GetClientIP(r)
}
