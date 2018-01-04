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
		{".eth.", ".eth."},
		{"wealdtech.eth", "wealdtech.eth"},
		{".wealdtech.eth", ".wealdtech.eth"},
		{"subdomain.wealdtech.eth", "subdomain.wealdtech.eth"},
		{"*.wealdtech.eth", "*.wealdtech.eth"},
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

func TestDomainPart(t *testing.T) {
	tests := []struct {
		input  string
		part   int
		output string
		err    bool
	}{
		{"", 1, "", false},
		{"", 2, "", true},
		{"", -1, "", false},
		{"", -2, "", true},
		{".", 1, "", false},
		{".", 2, "", false},
		{".", 3, "", true},
		{".", -1, "", false},
		{".", -2, "", false},
		{".", -3, "", true},
		{"ETH", 1, "eth", false},
		{"ETH", 2, "", true},
		{"ETH", -1, "eth", false},
		{"ETH", -2, "", true},
		{".ETH", 1, "", false},
		{".ETH", 2, "eth", false},
		{".ETH", 3, "", true},
		{".ETH", -1, "eth", false},
		{".ETH", -2, "", false},
		{".ETH", -3, "", true},
		{"wealdtech.eth", 1, "wealdtech", false},
		{"wealdtech.eth", 2, "eth", false},
		{"wealdtech.eth", 3, "", true},
		{"wealdtech.eth", -1, "eth", false},
		{"wealdtech.eth", -2, "wealdtech", false},
		{"wealdtech.eth", -3, "", true},
		{".wealdtech.eth", 1, "", false},
		{".wealdtech.eth", 2, "wealdtech", false},
		{".wealdtech.eth", 3, "eth", false},
		{".wealdtech.eth", 4, "", true},
		{".wealdtech.eth", -1, "eth", false},
		{".wealdtech.eth", -2, "wealdtech", false},
		{".wealdtech.eth", -3, "", false},
		{".wealdtech.eth", -4, "", true},
		{"subdomain.wealdtech.eth", 1, "subdomain", false},
		{"subdomain.wealdtech.eth", 2, "wealdtech", false},
		{"subdomain.wealdtech.eth", 3, "eth", false},
		{"subdomain.wealdtech.eth", 4, "", true},
		{"subdomain.wealdtech.eth", -1, "eth", false},
		{"subdomain.wealdtech.eth", -2, "wealdtech", false},
		{"subdomain.wealdtech.eth", -3, "subdomain", false},
		{"subdomain.wealdtech.eth", -4, "", true},
		{"a.b.c", 1, "a", false},
		{"a.b.c", 2, "b", false},
		{"a.b.c", 3, "c", false},
		{"a.b.c", 4, "", true},
		{"a.b.c", -1, "c", false},
		{"a.b.c", -2, "b", false},
		{"a.b.c", -3, "a", false},
		{"a.b.c", -4, "", true},
	}

	for _, tt := range tests {
		result, err := DomainPart(tt.input, tt.part)
		if err != nil && !tt.err {
			t.Errorf("Failure: %v, %v => error (unexpected)\n", tt.input, tt.part)
		}
		if err == nil && tt.err {
			t.Errorf("Failure: %v, %v => no error (unexpected)\n", tt.input, tt.part)
		}
		if tt.output != result {
			t.Errorf("Failure: %v, %v => %v (expected %v)\n", tt.input, tt.part, result, tt.output)
		}
	}
}
