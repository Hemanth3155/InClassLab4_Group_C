package calculator

import (
	"fmt"
	"testing"
)

// TestDivide tests the Divide function for valid and invalid cases.
func TestDivide(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
		err      bool
	}{
		{10, 2, 5, false},   // normal division
		{5, 0, 0, true},     // division by zero
		{-10, 2, -5, false}, // negative number
		{10, -2, -5, false}, // negative divisor
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d/%d", tt.a, tt.b), func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if tt.err && err == nil {
				t.Errorf("expected an error but got none")
			} else if !tt.err && err != nil {
				t.Errorf("unexpected error: %v", err)
			} else if result != tt.expected {
				t.Errorf("expected %d but got %d", tt.expected, result)
			}
		})
	}
}

// TestSquare tests the Square function.
func TestSquare(t *testing.T) {
	tests := []struct {
		a        int
		expected int
	}{
		{2, 4},       // positive number
		{-2, 4},      // negative number
		{0, 0},       // zero
		{100, 10000}, // large number
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Square(%d)", tt.a), func(t *testing.T) {
			result := Square(tt.a)
			if result != tt.expected {
				t.Errorf("expected %d but got %d", tt.expected, result)
			}
		})
	}
}
