package main

import (
	"fmt"
	"net/http"
)

type SecretHandler struct{}

func (h *SecretHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "secret handler")
}

type HealthCheckHandler struct{}

func (h *HealthCheckHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "healthcheck")
}

func main() {
	secretHandler := SecretHandler{}
	healthCheckHandler := HealthCheckHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/secret", &secretHandler)
	http.Handle("/healthcheck", &healthCheckHandler)

	server.ListenAndServe()
}
