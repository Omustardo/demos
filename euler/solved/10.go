package solved

import (
	"fmt"
)

func Problem10() int {
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
	for primes[len(primes)-1] < 2e6 {
		if isPrime(i) {
			primes = append(primes, i)
		}
		i += 2
	}
	sum := int64(0)
	for i := 0; i < len(primes)-1; i++ {
		sum += primes[i]
	}
	fmt.Println(sum)
	return int(sum)
}
