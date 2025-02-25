package utils

import (
	"testing"
)

func TestRoundFloatToTwoDecimals(t *testing.T) {
	tests := []struct {
		name   string
		input  float64
		result float64
	}{
		{"Positive number", 16.17432, 16.17},
		{"Long positive number", 16.166666666666666632, 16.17},
		{"Negative number", -55.123322, -55.12},
		{"Long negative number", -99.1779666666666632, -99.18},
		{"Zero", 0.0, 0.0},
		{"Whole number", 1, 1},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {
			result := RoundFloatToTwoDecimals(entry.input)

			if result != entry.result {
				t.Errorf("expected %f, but received %f", entry.result, result)
			}
		})
	}
}
