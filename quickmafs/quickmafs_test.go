package quickmafs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name           string
		a, b, expected int
	}{
		{"simple", 1, 2, 3},
		{"all_zeros", 0, 0, 0},
		{"negative", -1, -2, -3},
		{"different_signs", 1, -2, -1},
		{"huge", 1152921504606846976, 1152921504606846976, 2305843009213693952},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestSub(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name           string
		a, b, expected int
	}{
		{"simple", 1, 2, -1},
		{"all_zeros", 0, 0, 0},
		{"negative", -1, -2, 1},
		{"huge", 1152921504606846976, 2305843009213693952, -1152921504606846976},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Sub(tc.a, tc.b)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestMult(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name           string
		a, b, expected int
	}{
		{"simple", 1, 2, 2},
		{"by_zero", 1234, 0, 0},
		{"all_zeros", 0, 0, 0},
		{"negative", -2, -3, 6},
		{"different_signs", 1, -2, -2},
		{"huge", 1152921504606846976, 2, 2305843009213693952},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Mult(tc.a, tc.b)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestDiv(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name     string
		a, b     int
		expected int
		err      error
	}{
		{"simple", 10, 2, 5, nil},
		{"negative", -10, -2, 5, nil},
		{"different_signs", -10, 2, -5, nil},
		{"huge", 1152921504606846976, 2, 576460752303423488, nil},
		{"div_by_zero", 1, 0, 0, ErrDivByZero},
		{"with_reminder", 11, 3, 3, nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := Div(tc.a, tc.b)
			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestPrimes(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name     string
		count    int
		expected []int
	}{
		{"simple", 5, []int{2, 3, 5, 7, 11}},
		{"count_zero", 0, nil},
		{"huge", 100, []int{
			2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79,
			83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
			179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269,
			271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373,
			379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467,
			479, 487, 491, 499, 503, 509, 521, 523, 541},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Primes(tc.count)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestFactorize(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name     string
		n        int
		expected []int
	}{
		{"simple", 30, []int{2, 3, 5}},
		{"zero", 0, nil},
		{"one", 1, []int{1}},
		{"repeated", 32, []int{2, 2, 2, 2, 2}},
		{"large_prime", 82372946003, []int{82372946003}},
		{"large_non_prime", 82372946004, []int{2, 2, 3, 3, 3, 17, 41, 61, 17939}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Factorize(tc.n)
			assert.Equal(t, tc.expected, got)
		})
	}
}
