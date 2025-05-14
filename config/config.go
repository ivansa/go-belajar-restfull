package config

import (
	"go-belajar-restfull/middleware"
	"net/http"
)

func NewHttpServer(middleware *middleware.ApiMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: middleware,
	}
}
