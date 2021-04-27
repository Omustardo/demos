package pandigital

import "fmt"

// Pandigital returns if the provided number is a pandigital number for the provided range.
// For example, 12345 is pandigital with min 1 and max 5.
func Pandigital(n int64, min, max int) bool {
	s := fmt.Sprintf("%d", n)
	if len(s) != (max - min + 1) {
		return false
	}
	digits := make(map[rune]bool)
	for _, d := range s {
		if digits[d] {
			return false
		}
		if d-'0' > int32(max) || d-'0' < int32(min) {
			return false
		}
		digits[d] = true
	}
	return true
}
