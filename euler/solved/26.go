package solved

import "fmt"

func Problem26() int {
	// I believe we only need to look at prime numbers since they should be the only decimals with repetition,
	// besides numbers that have those primes as factors (e.g. 1/3==.333 1/9=.111).
	// I'm not quite sure how to determine when we have a repetition, except by doing division step by step
	// and determining when the value we're at and the current remainder are the same as something that
	// we've seen before...

	type DigitRemainder struct {
		digit, remainder int
	}

	// digits returns the digits past the decimal obtained by dividing 1 by the input.
	// e.g. 8 -> [1 2 5]
	digits := func(n int) []int {
		m := make(map[DigitRemainder]bool)
		x := 1
		var digits []int
		for i := 0; i < 10000; i++ {
			if x == 0 {
				return digits
			}
			if x < n {
				x = x * 10
			}
			dr := DigitRemainder{digit: x, remainder: x / n}
			if _, ok := m[dr]; ok { // already seen this, so we know it's a repeat
				return digits
			}
			m[dr] = true
			digits = append(digits, x/n)
			x = x % n
		}
		return digits
	}

	var maxN int
	var max []int
	for n := 7; n < 1000; n++ {
		d := digits(n)
		if len(d) > len(max) {
			maxN = n
			max = d
		}
		// fmt.Printf("1/%d has len %d: %v\n", n, len(d), d)
	}

	fmt.Printf("Longest series from 1/%d has len %d: %v\n", maxN, len(max), max)
	return maxN
}
