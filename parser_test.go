package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_ValidInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"123+456", []string{"123", "+", "456"}},
		{" 7 * 8 - 9 ", []string{"7", "*", "8", "-", "9"}},
		{"10/2+5", []string{"10", "/", "2", "+", "5"}},
		// więcej dobrych przypadków
	}

	for _, tc := range cases {
		result, err := parser(tc.input)
		if err != nil {
			t.Fatalf("dla wejścia %q nie spodziewano się błędu, otrzymano: %v", tc.input, err)
		}
		// Porównanie długości wyników
		if len(result) != len(tc.expected) {
			t.Errorf("dla wejścia %q otrzymano %v (długość %d), oczekiwano %v (długość %d)",
				tc.input, result, len(result), tc.expected, len(tc.expected))
			continue
		}
		// Porównanie kolejnych tokenów
		for i := range result {
			if result[i] != tc.expected[i] {
				t.Errorf("dla wejścia %q token[%d]=%q, oczekiwano %q", tc.input, i, result[i], tc.expected[i])
			}
		}
	}
}

// func TestParser_EmptyInput(t *testing.T) {
// 	defer func() {
// 		if r := recover(); r == nil {
// 			t.Errorf("oczekiwano panic dla pustego wejścia")
// 		}
// 	}()
// 	_, _ = parser("")
// }

func TestParser_InvalidInput(t *testing.T) {
	cases := []struct {
		input       string
		errContains string
	}{
		{"+1+2", "start with"},    // zacznij od cyfry
		{"1+2-", "end with"},      // koniec cyfrą
		{"1a+2", "unsupported"},   // nieobsługiwany znak 'a'
		{"1 +\t2", "unsupported"}, // znak tabulacji '\t' nie jest usuwany przez RemoveWhiteSpaces
	}

	for _, tc := range cases {
		result, err := parser(tc.input)
		assert.Nil(t, result, "powinno zwracać nil przy błędzie, wejście: %q", tc.input)
		if assert.Error(t, err, "oczekiwano błędu dla wejścia %q", tc.input) {
			assert.Contains(t, err.Error(), tc.errContains, "wejście: %q", tc.input)
		}
	}
}
