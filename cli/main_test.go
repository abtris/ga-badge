package main

import "testing"

func TestGenerateBadge(t *testing.T) {
	tests := []struct {
		name string
		url   string
		branch string
		label string
		output  string
	}{
		{
			name: "Simple ok",
			url: "https://github.com/abtris/ga-badge/actions/workflows/node.js.yml",
			branch:"master",
			label: "Build Status",
			output: "[![Build Status](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg)](https://github.com/abtris/ga-badge/actions)",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, _ := generateBadge(test.url, test.branch, test.label)
			if actual != test.output {
				t.Errorf("Expected %v and real %v)", test.output, actual)
			}
		})
	}
}
