package mymath

import (
	"time"
	"math/rand"
	"math"
)

// Generates a slice of size, size filled with random numbers
func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
// WithBranch uses control structures to return the absolute value of an
// integer.
func AbsWithBranch(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
func AbsInt64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
// WithStdLib uses the standard library's math package to compute the
// absolute value on an integer.
//
// We expect test for correctness to fail on large numbers that overflow
// float64.
func AbsWithStdLib(n int64) int64 {
	return int64(math.Abs(float64(n)))
}

// WithTwosComplement uses a trick from Henry S. Warren's incredible book,
// Hacker's Delight. It utilizes Two's Complement arithmetic to compute the
// absolute value of an integer.
func AbsWithTwosComplement(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

// WithASM uses the Two's Complement trick, but implemented in Assembly to
// compute the absolute value of a signed integer.
//func WithASM(n int64) int64

