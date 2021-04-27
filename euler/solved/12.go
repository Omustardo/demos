package solved

import (
	"fmt"
)

func Problem12() int {

	var primes = []int64{2, 3, 5, 7, 11, 13}

	isPrime := func(n int64) bool {
		for _, p := range primes {
			if n%p == 0 {
				return false
			}
		}
		return true
	}

	var i int64 = 15
	fmt.Printf("Calculating primes up to %f\n", 2e5)
	for primes[len(primes)-1] < 2e5 {
		if isPrime(i) {
			primes = append(primes, i)
		}
		i += 2
	}

	countDivisors := func(n int64) int64 {
		// if you have the prime factorization of the number n, then to calculate how many divisors it has, you take all the exponents in the factorization, add 1 to each, and then multiply these "exponents + 1"s together.
		prod := int64(1)
		for i := range primes {
			if n == 0 || primes[i] > n {
				//fmt.Printf("primes[i]=%d  n=%d\n", primes[i], n)
				break
			}

			count := int64(0)
			for n%primes[i] == 0 {
				//fmt.Printf("got factor=%d\n", primes[i])
				count++
				n /= primes[i]
			}
			if count > 0 {
				//fmt.Printf("Got factor: %d. Found %d times\n", primes[i], count)
				prod *= count + 1
			}
			i++
		}
		return prod
	}

	tri := int64(1)
	cur := int64(1)
	divisors := countDivisors(tri)
	for divisors < 500 {
		if cur%100 == 0 {
			fmt.Printf("Tri#%d (%d) has %d divisors\n", cur, tri, divisors)
		}
		cur++
		tri += cur
		divisors = countDivisors(tri)
	}

	fmt.Printf("Tri#%d (%d) has %d divisors\n", cur, tri, divisors)
	return int(tri)
}
