package utils

import (
	"net/http"
	"time"
)

const (
	HTTP_CLIENT_MAX_CONNECTIONS = 20
	HTTP_CLIENT_REQUEST_TIMEOUT = 5
)

var HttpClient *http.Client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: HTTP_CLIENT_MAX_CONNECTIONS,
	},
	Timeout: time.Duration(HTTP_CLIENT_REQUEST_TIMEOUT) * time.Second,
}
