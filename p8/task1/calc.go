package main

import (
	"errors"
)

var errDivisionByZero error = errors.New("division by zero")

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errDivisionByZero
	}
	return a / b, nil
}
