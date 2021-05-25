package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

func setupParseArgs(w io.Writer, args []string) (appConfig, error) {
	conf := appConfig{}

	fs := flag.NewFlagSet("secret-client", flag.ContinueOnError)
	fs.SetOutput(w)

	fs.StringVar(&conf.action, "action", "", "Create/View secret (create|view)")
	fs.StringVar(&conf.plainText, "data", "", "Secret to store")
	fs.StringVar(&conf.secretID, "id", "", "ID of secret to fetch")

	fs.StringVar(&conf.apiURL, "url", "", "URL for the Secret Web Application API")

	err := fs.Parse(args)

	if err != nil {
		return conf, err
	}

	if fs.NArg() != 0 {
		return conf, errors.New("No positional parameters expected")
	}
	return conf, err
}

func validateConf(conf appConfig) []error {
	var validationErrors []error

	if len(conf.apiURL) == 0 {
		validationErrors = append(validationErrors, errors.New("API URL not specified"))
	}

	switch conf.action {
	case "create":
		if len(conf.plainText) == 0 {
			validationErrors = append(validationErrors, errors.New("create: Plain text not specified"))
		}

	case "view":
		if len(conf.secretID) == 0 {
			validationErrors = append(validationErrors, errors.New("view: Secret ID not specified"))
		}
	default:
		if len(conf.action) == 0 {
			validationErrors = append(validationErrors, errors.New("No action specified"))
		} else {
			validationErrors = append(validationErrors, fmt.Errorf("Invalid action specified: %v", conf.action))
		}
	}

	return validationErrors
}
