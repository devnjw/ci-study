package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 5 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 5},
		{-1, 2},
		{0, 3},
		{-5, -2},
		{99999, 100002},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
