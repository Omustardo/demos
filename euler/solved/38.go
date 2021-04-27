package solved

import (
	"fmt"
	"strconv"
)

func Problem38() int {

	// pandigital returns whether the provided string is a 9 digit number containing each digit [1-9] exactly once.
	pandigital := func(s string) bool {
		if len(s) != 9 {
			return false
		}
		var digits [10]bool
		for _, c := range s {
			digits[c-'0'] = true
		}
		for i := 1; i < 10; i++ {
			if !digits[i] {
				return false
			}
		}
		return true
	}

	var max int
	// anything above 1e5 can't be pandigital product since it would be at least two 5 digit numbers
	for i := 100; i < 1e5; i++ {
		var out string
		for n := 1; len(out) < 9; n++ {
			out = fmt.Sprintf("%s%d", out, i*n) // this is not particularly efficient... but it's fine.
		}
		if pandigital(out) {
			fmt.Println(i, out)
			max, _ = strconv.Atoi(out)
		}
	}

	return max
}
