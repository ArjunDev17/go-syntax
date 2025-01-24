package main

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},           // valid email
		{"invalid-email", false},             // invalid email
		{"test@.com", false},                 // invalid email
		{"user@domain.co", true},             // valid email
		{"@domain.com", false},               // invalid email
		{"test@subdomain.example.com", true}, // valid email
	}

	for _, test := range tests {
		t.Run(test.email, func(t *testing.T) {
			actual := ValidateEmail(test.email)
			if actual != test.expected {
				t.Errorf("expected %v, but got %v for email %s", test.expected, actual, test.email)
			}
		})
	}
}
