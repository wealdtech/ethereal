package ens

import (
	"testing"
)

func TestNormaliseDomain(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{".", "."},
		{"eth", "eth"},
		{"ETH", "eth"},
		{".eth", ".eth"},
		{"wealdtech.eth", "wealdtech.eth"},
		{".wealdtech.eth", ".wealdtech.eth"},
		{"subdomain.wealdtech.eth", "subdomain.wealdtech.eth"},
	}

	for _, tt := range tests {
		result := NormaliseDomain(tt.input)
		if tt.output != result {
			t.Errorf("Failure: %v => %v (expected %v)\n", tt.input, result, tt.output)
		}
	}
}

func TestTld(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{".", ""},
		{"eth", "eth"},
		{"ETH", "eth"},
		{".eth", "eth"},
		{"wealdtech.eth", "eth"},
		{".wealdtech.eth", "eth"},
		{"subdomain.wealdtech.eth", "eth"},
	}

	for _, tt := range tests {
		result := Tld(tt.input)
		if tt.output != result {
			t.Errorf("Failure: %v => %v (expected %v)\n", tt.input, result, tt.output)
		}
	}
}
