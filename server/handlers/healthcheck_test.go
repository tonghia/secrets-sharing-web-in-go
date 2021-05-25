package handlers

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckHandlerUnrecognizedMethod(t *testing.T) {

	req := httptest.NewRequest("GET", "http://test-server.com/healthcheck", nil)
	w := httptest.NewRecorder()
	healthCheckHandler(w, req)

	response := w.Result()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != 200 {
		t.Errorf("Expected Response Status to be: %d, Got: %d", 200, response.StatusCode)

	}
	expectedResponseBody := "ok"
	if string(body) != expectedResponseBody {
		t.Errorf("Expected Response body to be: %s, Got: %s", expectedResponseBody, string(body))
	}
}
