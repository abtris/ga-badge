package main

import (
	"io/fs"
	"os"
	"testing"
)

func TestGenerateBadge(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		branch      string
		label       string
		output      string
		outputError error
	}{
		{
			name:        "Simple ok",
			url:         "https://github.com/abtris/ga-badge/actions/workflows/node.js.yml",
			branch:      "master",
			label:       "Build Status",
			output:      "[![Build Status](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg)](https://github.com/abtris/ga-badge/actions)",
			outputError: nil,
		},
		{
			name:        "Simple branch ok",
			url:         "https://github.com/abtris/ga-badge/actions/workflows/node.js.yml",
			branch:      "test",
			label:       "Build Status",
			output:      "[![Build Status](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg?branch=test)](https://github.com/abtris/ga-badge/actions)",
			outputError: nil,
		},
		{
			name:        "Simple label ok",
			url:         "https://github.com/abtris/ga-badge/actions/workflows/node.js.yml",
			branch:      "test",
			label:       "Build",
			output:      "[![Build](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg?branch=test)](https://github.com/abtris/ga-badge/actions)",
			outputError: nil,
		},
		{
			name:        "Simple fail",
			url:         "XXX",
			branch:      "test",
			label:       "Build",
			output:      "",
			outputError: errWrongURL,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, actualError := generateBadge(test.url, test.branch, test.label)
			if actual != test.output {
				t.Errorf("Expected %v and real %v)", test.output, actual)
			}
			if actualError != test.outputError {
				t.Errorf("Expected error '%v' and real '%v'", test.outputError, actualError)
			}
		})
	}
}

func TestInitWorkflow(t *testing.T) {
	tests := []struct {
		name           string
		lang           string
		outputError    error
		outputFileName string
	}{
		{name: "Simple Node template", lang: "node", outputError: nil, outputFileName: "node.yaml"},
		{name: "Simple Node.js template", lang: "node.js", outputError: nil, outputFileName: "node.js.yaml"},
		{name: "Simple Go template", lang: "go", outputError: nil, outputFileName: "go.yaml"},
	}
	templates, _ := fs.ReadDir(files, "templates")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualFileName, err := initWorkflow(test.lang, "/tmp/.github/workflows", templates)
			defer os.RemoveAll("/tmp/.github/workflows")
			if err != nil {
				t.Errorf("Error %v", err)
			}
			if actualFileName != test.outputFileName {
				t.Errorf("Expected %v and real %v)", test.outputFileName, actualFileName)
			}
		})
	}
}
