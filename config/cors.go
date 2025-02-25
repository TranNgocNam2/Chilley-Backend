package config

import (
	"net/http"
	"time"
)

var (
	AllowMethods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions}
	AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	CorsMaxAge   = 12 * time.Hour
)
