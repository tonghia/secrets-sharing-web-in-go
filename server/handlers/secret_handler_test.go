package handlers

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestSecretHandlerUnrecognizedMethod(t *testing.T) {

	// Create a secret
	req := httptest.NewRequest("PUT", "http://test-server.com", nil)
	w := httptest.NewRecorder()
	secretHandler(w, req)

	response := w.Result()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != 405 {
		t.Errorf("Expected Response Status to be: %d, Got: %d", 405, response.StatusCode)

	}
	expectedResponseBody := "Method not allowed\n"
	if string(body) != expectedResponseBody {
		t.Errorf("Expected Response body to be: %s, Got: %s", expectedResponseBody, string(body))
	}
}
