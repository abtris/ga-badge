package main

import (
	"testing"
)

func TestGenerateBadge(t *testing.T) {
	tests := []struct {
		name string
		url   string
		branch string
		label string
		output  string
		outputError error
	}{
		{
			name: "Simple ok",
			url: "https://github.com/abtris/ga-badge/actions/workflows/node.js.yml",
			branch:"master",
			label: "Build Status",
			output: "[![Build Status](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg)](https://github.com/abtris/ga-badge/actions)",
			outputError: nil,
		},
		{
			name: "Simple branch ok",
			url: "https://github.com/abtris/ga-badge/actions/workflows/node.js.yml",
			branch:"test",
			label: "Build Status",
			output: "[![Build Status](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg?branch=test)](https://github.com/abtris/ga-badge/actions)",
			outputError: nil,
		},
		{
			name: "Simple label ok",
			url: "https://github.com/abtris/ga-badge/actions/workflows/node.js.yml",
			branch:"test",
			label: "Build",
			output: "[![Build](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg?branch=test)](https://github.com/abtris/ga-badge/actions)",
			outputError: nil,
		},
		{
			name: "Simple fail",
			url: "XXX",
			branch:"test",
			label: "Build",
			output: "",
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

}
