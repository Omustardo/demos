package solved

import (
	"fmt"
	"github.com/omustardo/demos/euler/prime"
	"math"
)

func Problem27() int {
	var primesUnder1000 []int
	for i := 2; i < 1000; i++ {
		if prime.IsPrime(int64(i)) {
			primesUnder1000 = append(primesUnder1000, i)
		}
	}

	// validateQuadPrimes determines if the given quadratic function (based on a,b: n*n+a*n+b) results in primes for every 'n' from [0,start]
	// For example, validateQuadPrimes(1,41, 30) == true but validateQuadPrimes(1,41, 41) == false
	validateQuadPrimes := func(a, b, start int) bool {
		i := start
		for {
			if i == -1 {
				return true
			}
			if prime.IsPrime(int64(i*i + a*i + b)) {
				i--
			} else {
				return false
			}
		}
	}

	// checkQuadPrimes returns the number of consecutive primes resulting from inputs to the quadratic function n*n+a*n+b, starting with 0 and incrementing by 1.
	// For example, checkQuadPrimes(1,41, 0) == 40
	checkQuadPrimes := func(a, b, start int) int {
		i := start
		for {
			if prime.IsPrime(int64(i*i + a*i + b)) {
				i++
			} else {
				return i
			}
		}
	}

	var max, maxA, maxB int
	for a := -999; a <= 999; a++ {
		for _, b := range primesUnder1000 { // b must be prime, otherwise f(0) wouldn't be prime.
			// when n == b, it's guaranteed to not result in a prime, so if b < max we don't need to check.
			if math.Abs(float64(b)) < float64(max) {
				continue
			}

			// If it can't beat the current best, don't bother.
			if !validateQuadPrimes(a, b, max+1) {
				continue
			}

			c := checkQuadPrimes(a, b, max+1)
			if c > max {
				fmt.Printf("New Best: [n*n + %dn + %d]: %d\n", a, b, c)
				max = c
				maxA = a
				maxB = b
			}
		}
	}

	fmt.Printf("MAX FOUND: [n*n + %dn + %d]: %d\n", maxA, maxB, max)

	return maxA * maxB
}
