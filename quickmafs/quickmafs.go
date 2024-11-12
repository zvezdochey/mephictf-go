package quickmafs

import "errors"

var (
	ErrDivByZero = errors.New("division by zero")
)

// Adds two integers.
func Add(a, b int) int {
	return 0
}

// Subtracts b from a.
func Sub(a, b int) int {
	return 0
}

// Multiplies two integers.
func Mult(a, b int) int {
	return 0
}

// Divides a by b.
//
// Returns ErrDivByZero if b is zero.
func Div(a, b int) (int, error) {
	return 0, errors.New("not implemented")
}

// Returns first count primes.
func Primes(count int) []int {
	return nil
}

// Returns prime factors of n in ascending order.
func Factorize(n int) []int {
	return nil
}
