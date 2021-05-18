package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func secretHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "secret handler")
}

func healthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "healthcheck")
}

func main() {
	http.HandleFunc("/", secretHandler)
	http.HandleFunc("/healthcheck", healthCheckHandler)

	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
