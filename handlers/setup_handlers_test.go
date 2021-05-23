package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupHandler(t *testing.T) {
	mux := http.NewServeMux()
	SetupHandlers(mux)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/healthcheck")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if string(resp) != "ok" {
		t.Errorf("Expected GET /healthcheck request to return ok, Got: %s", string(resp))
	}

	// Test for GET request to /
	res, err = http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	expectedRespBody := "{\"data\":\"\"}"
	if string(resp) != expectedRespBody {
		t.Errorf("Expected GET / request to return: %s, Got: %s", expectedRespBody, string(resp))
	}

	// Test for POST request to /
	res, err = http.Post(ts.URL, "application/json", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if string(resp) != "Invalid Request\n" {
		t.Errorf("Expected POST / request to return: Invalid Request, Got: %s", string(resp))
	}

}
