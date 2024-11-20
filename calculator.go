package calculator

import (
	"errors"
)

// Divide divides two integers and returns the result and an error if dividing by zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Square returns the square of the given integer.
func Square(a int) int {
	return a * a
}
