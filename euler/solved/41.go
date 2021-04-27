package solved

import (
	"fmt"
	"github.com/omustardo/demos/euler/prime"
)

func Problem41() int {
	primes := prime.List(1e9)
	fmt.Println("Done calculating all 9 digit primes")

	for i := len(primes) - 1; i >= 0; i-- {
		var digits [10]bool
		pandigital := true
		s := fmt.Sprintf("%d", primes[i])
		for _, d := range s {
			if digits[d-'0'] || d-'0' == 0 {
				pandigital = false
				break
			}
			digits[d-'0'] = true
		}
		for j := 1; pandigital && j <= len(s); j++ {
			if !digits[j] {
				pandigital = false
			}
		}

		if pandigital {
			fmt.Println("Largest pandigital prime is", primes[i])
			return int(primes[i])
		}
	}

	return -1
}
