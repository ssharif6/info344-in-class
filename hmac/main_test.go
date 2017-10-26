package main

import (
	"testing"
	"strings"
)

func TestSign(t *testing.T) {
	//TODO: write unit test cases for sign()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader

	cases := []struct {
		input      string
		signingKey string
		output     string
	}{
		{
			input:      "valid case",
			signingKey: "passwordKey",
			output:     "WbfBXEvODe6CYD7AkoM10QYetpmRvD+4CgOV26KtLJw=",
		},
	}
	for _, c := range cases {
		output, err := sign(c.signingKey, strings.NewReader(c.input))
		if err != nil {
			t.Errorf("Error signing %v", err)
		}
		 if output != c.output {
		 	t.Errorf("Expected %s but got %s", c.output, output)
		 }

	}
}

func TestVerify(t *testing.T) {
	//TODO: write until test cases for verify()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
}
