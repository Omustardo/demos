package solved

import (
	"github.com/omustardo/demos/euler/prime"
)

// An insight to this problem is that we can immediately ignore any primes which contain 0,2,4,5,6,8.
// They are guaranteed to have non-prime rotations when those digits are in the 1's place.
// With that knowledge, you could find all numbers < 1e6 made up of only 1,3,7,9 and only test them for primality.
//
// Runtime is already very quick though so it wasn't necessary to cut down the input primes.

func Problem35() int {
	// returns all rotations of a given input. Note that it doesn't do anything special for duplicates.
	// (196) -> [196 619 961]
	// (11) -> [11 11]
	rotations := func(n int64) []int64 {
		if n < 10 {
			return []int64{n}
		}
		if n < 100 {
			return []int64{n, (n%10)*10 + n/10}
		}

		var digits []int64
		tmp := n
		for tmp > 0 {
			digits = append(digits, tmp%10)
			tmp /= 10
		}
		// reverse the digits so they're in the correct order.
		for i := 0; i < len(digits)/2; i++ {
			digits[i], digits[len(digits)-1-i] = digits[len(digits)-1-i], digits[i]
		}

		out := []int64{n}
		// Start at 1. Since we already know the input is part of output.
		for start := 1; start < len(digits); start++ {
			curr := int64(digits[start])
			for i := start + 1; i%len(digits) != start; i++ {
				i %= len(digits)
				curr = curr*10 + digits[i]
			}
			out = append(out, curr)
		}
		return out
	}

	var count int
	for _, p := range prime.List(1e6) {
		rotPrime := true
		for _, rot := range rotations(p) {
			if !prime.IsPrime(rot) {
				rotPrime = false
				break
			}
		}
		if rotPrime {
			// fmt.Printf("%d is a circular prime! %v\n", p, rotations(p))
			count++
		}
	}

	return count
}
