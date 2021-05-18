package main

import (
	"fmt"
	"net/http"
)

func secretHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "secret handler")
}

func healthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "healthcheck")
}

func main() {
	http.HandleFunc("/", secretHandler)
	http.HandleFunc("/healthcheck", healthCheckHandler)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
