package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type CreateSecretResponse struct {
	Id string `json:"id"`
}

type GetSecretResponse struct {
	Data string `json:"data"`
}

func createSecret(apiURL string, plainText string) (CreateSecretResponse, error) {
	s := CreateSecretResponse{}
	var jsonBody = fmt.Sprintf(`{"plain_text":"%s"}`, plainText)
	resp, err := http.Post(apiURL, "application/json", strings.NewReader(jsonBody))
	if err != nil {
		return s, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return s, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return s, fmt.Errorf("%s", string(data))
	}

	err = json.Unmarshal(data, &s)
	return s, err
}

func getSecret(apiURL string, secretID string) (GetSecretResponse, error) {
	s := GetSecretResponse{}
	resp, err := http.Get(apiURL + "/" + secretID)
	if err != nil {
		return s, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return s, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return s, fmt.Errorf("%s", string(data))
	}

	err = json.Unmarshal(data, &s)
	return s, err
}
