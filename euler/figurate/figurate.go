// Package figurate provides tools for dealing with Figurate Numbers.
// https://en.wikipedia.org/wiki/Figurate_number
package figurate

import "math"

func IsTriangular(x int) (bool, int) {
	n := (math.Sqrt(8*float64(x)+1) + 1) / 2
	if n == math.Floor(n) {
		return true, int(n)
	}
	return false, 0
}

// IsPentagonal returns whether the given value is a pentagonal number and which pentagonal number it is.
// Note that this isn't for generalized pentagonal numbers (including 0 and negative numbers).
func IsPentagonal(x int) (bool, int) {
	n := (math.Sqrt(24*float64(x)+1) + 1) / 6
	if n == math.Floor(n) {
		return true, int(n)
	}
	return false, 0
}

func IsHexagonal(x int) (bool, int) {
	n := (math.Sqrt(8*float64(x)+1) + 1) / 4
	if n == math.Floor(n) {
		return true, int(n)
	}
	return false, 0
}

func Triangular(n int) int {
	return n * (n + 1) / 2
}
func Pentagonal(n int) int {
	return n * (2*n - 1) / 2
}
func Hexagonal(n int) int {
	return n * (2*n - 1)
}
