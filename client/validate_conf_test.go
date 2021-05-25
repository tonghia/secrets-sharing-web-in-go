package main

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {

	type testConfig struct {
		conf appConfig
		errs []error
	}

	testCases := []testConfig{
		testConfig{
			conf: appConfig{
				apiURL:    "http://localhost:8080/api",
				action:    "create",
				plainText: "Super Secret",
			},
			errs: []error{},
		},
		testConfig{
			conf: appConfig{
				apiURL:    "http://localhost:8080/api",
				action:    "view",
				plainText: "Super Secret",
			},
			errs: []error{
				errors.New("view: Secret ID not specified"),
			},
		},
		testConfig{
			conf: appConfig{},
			errs: []error{
				errors.New("API URL not specified"),
				errors.New("No action specified"),
			},
		},
	}
	for _, tc := range testCases {
		errs := validateConf(tc.conf)
		if len(tc.errs) == 0 && len(errs) != 0 {
			t.Errorf("Expected no errors, got: %v", errs)
		}

		if len(tc.errs) != 0 {
			for i, e := range tc.errs {
				if errs[i] == nil || errs[i].Error() != e.Error() {
					t.Errorf("Expected error: %v, Got: %v", e, errs[i])
				}
			}
		}
	}
}
