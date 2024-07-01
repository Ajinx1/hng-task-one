package repositories

import (
	"net"
	"net/http"
	"strings"
)

type IPRepository interface {
	GetClientIP(r *http.Request) string
}

type ipRepository struct{}

func NewIPRepository() IPRepository {
	return &ipRepository{}
}

func (r *ipRepository) GetClientIP(req *http.Request) string {
	xff := req.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	ip := req.RemoteAddr
	if strings.Contains(ip, ":") {
		ip, _, _ = net.SplitHostPort(req.RemoteAddr)
	}

	if ip == "::1" {
		ip = "127.0.0.1"
	}

	return ip
}
