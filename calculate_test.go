package main

import (
	"errors"
	"testing"
)

// Zakładam, że calculate znajduje się w tym samym pakiecie co test
func TestCalculate_ValidExpressions(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected float64
	}{
		{"SimpleAddition", []string{"2", "+", "3"}, 5},
		{"SimpleSubtraction", []string{"10", "-", "4"}, 6},
		{"SimpleMultiplication", []string{"3", "*", "4"}, 12},
		{"SimpleDivision", []string{"8", "/", "2"}, 4},
		{"FloatAddition", []string{"2.5", "+", "1.5"}, 4},
		{"FloatDivision", []string{"7.5", "/", "2.5"}, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := calculate(tc.input)
			if err != nil {
				t.Fatalf("Unexpected error for %v: %v", tc.input, err)
			}
			if result != tc.expected {
				t.Errorf("Expected %f, got %f", tc.expected, result)
			}
		})
	}
}

func TestCalculate_InvalidExpressions(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		expectedErr error
	}{
		{"MissingOperator", []string{"1", "2", "3"}, errors.New("function 'calculate' returned in a way it shouldn't (i think)")},
		{"MissingOperands", []string{"+", "2"}, errors.New("not enough arguments")},
		{"TooFewElements", []string{"2"}, errors.New("not enough arguments")},
		{"InvalidLeftOperand", []string{"a", "+", "3"}, nil}, // error will be printed, but not returned
		{"InvalidRightOperand", []string{"3", "+", "b"}, nil},
		{"DivisionByZero", []string{"5", "/", "0"}, nil}, // no error returned, but Go returns +Inf
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := calculate(tc.input)
			if tc.expectedErr != nil {
				if err == nil || err.Error() != tc.expectedErr.Error() {
					t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			}
		})
	}
}
