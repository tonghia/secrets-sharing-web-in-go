package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tonghia/secrets-sharing-web-go/handlers"
)

func main() {

	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	// Create and setup handlers - we create a new ServeMux object here and pass
	// that as a parameter to the SetupHandlers() function
	mux := http.NewServeMux()
	handlers.SetupHandlers(mux)

	err := http.ListenAndServe(listenAddr, mux)
	if err != nil {
		log.Fatalf("Server could not start listening on %s. Error: %v", listenAddr, err)
	}

}
