package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestSetupParseFlags(t *testing.T) {

	type testConfig struct {
		args                   []string
		err                    error
		expectedConf           appConfig
		expectedOutputContains string
	}

	testCases := []testConfig{
		testConfig{
			args: []string{"-url", "http://localhost:8080/api", "-action", "create", "-data", "My super secret"},
			err:  nil,
			expectedConf: appConfig{
				apiURL:    "http://localhost:8080/api",
				action:    "create",
				plainText: "My super secret",
			},
		},
		testConfig{
			args: []string{"-url", "http://localhost:8080/api", "-action", "view", "-id", "abgtc12131a"},
			err:  nil,
			expectedConf: appConfig{
				apiURL:   "http://localhost:8080/api",
				action:   "view",
				secretID: "abgtc12131a",
			},
		},
		testConfig{
			args:         []string{"foo"},
			err:          errors.New("No positional parameters expected"),
			expectedConf: appConfig{},
		},
		testConfig{
			args:                   []string{"-h"},
			err:                    errors.New("flag: help requested"),
			expectedConf:           appConfig{},
			expectedOutputContains: "Usage of secret-client:",
		},
	}
	byteBuf := new(bytes.Buffer)
	for _, tc := range testCases {
		c, err := setupParseArgs(byteBuf, tc.args)
		if tc.err == nil && err != nil {
			t.Errorf("Expected non-nil error, got: %v", err)
		}

		if tc.err != nil {
			if err == nil || err.Error() != tc.err.Error() {
				t.Errorf("Expected error: %v, Got: %v", tc.err, err)
			}
		}

		if c != tc.expectedConf {
			t.Errorf("Expected:%#v Got: %#v", c, tc.expectedConf)
		}

		if len(tc.expectedOutputContains) != 0 {
			gotOutput := byteBuf.String()
			if strings.Index(gotOutput, tc.expectedOutputContains) == -1 {
				t.Errorf("Expected output: %s, Got: %s", tc.expectedOutputContains, gotOutput)
			}
		}
	}
}
