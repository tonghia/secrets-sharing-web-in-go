package main

import (
	"errors"
	"fmt"
)

func performAction(conf appConfig) (string, error) {
	switch conf.action {
	case "create":
		s, err := createSecret(conf.apiURL, conf.plainText)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s/%s", conf.apiURL, s.Id), nil

	case "view":
		s, err := getSecret(conf.apiURL, conf.secretID)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(s.Data), nil
	default:
		return "", errors.New("Invalid action")
	}
}
