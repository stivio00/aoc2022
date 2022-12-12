package main

import (
	"testing"
)

func TestWebPageSample(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "2-4,6-8",
			expected: false,
		},
		{
			input:    "5-7,7-9",
			expected: false,
		},
		{
			input:    "2-8,3-7",
			expected: true,
		},
		{
			input:    "6-6,4-6",
			expected: true,
		},
		{
			input:    "2-6,4-8",
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			pair := ParseLine(tC.input)
			result := pair.FullContains()
			if result != tC.expected {
				t.Errorf("FullConatins in  %v returns %v, expected %v \n", pair, result, tC.expected)
			}
		})
	}
}
