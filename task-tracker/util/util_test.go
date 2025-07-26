package util

import "testing"

func TestCheckValidStatus(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"done", true},
		{"todo", true},
		{"in-progress", true},
		{"invalid", false},
		{"", false},
	}

	for _, tt := range tests {
		result := CheckValidStatus(tt.input)
		if result != tt.expected {
			t.Errorf("CheckValidStatus(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
