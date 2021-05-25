package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var secretHash string = "bcc286bbbe4353e6a97ae169729ed4a5"

func dummySecretHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id":"%s"}`, secretHash)
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, `{"data":"my super secret"}`)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func TestPerformAction(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", dummySecretHandler)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	secretData := "my super secret"
	expectedSecretCreationResponse := ts.URL + "/" + secretHash

	c := appConfig{
		apiURL:    ts.URL,
		action:    "create",
		plainText: secretData,
	}

	o, err := performAction(c)
	if err != nil {
		t.Fatal(err)
	}

	if o != expectedSecretCreationResponse {
		t.Fatalf("Expected: %s, Got: %s", expectedSecretCreationResponse, o)
	}

	c = appConfig{
		apiURL:   ts.URL,
		action:   "view",
		secretID: secretHash,
	}

	o, err = performAction(c)
	if err != nil {
		t.Fatal(err)
	}

	if o != secretData {
		t.Fatalf("Expected: %s, Got: %s", secretData, o)
	}

}

func TestPerformInvalidAction(t *testing.T) {

	c := appConfig{
		action: "invalid",
	}

	_, err := performAction(c)
	if err == nil {
		t.Fatalf("Expected non-nil error, got: %v", err)
	}
	if err.Error() != "Invalid action" {
		t.Fatalf("Expected error to be: Invalid action, Got: %v", err)
	}
}
