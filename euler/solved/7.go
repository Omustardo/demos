package solved

import (
	"fmt"
)

func Problem7() int {
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
	for len(primes) < 10001 {
		if isPrime(i) {
			primes = append(primes, i)
		}
		i += 2
	}
	fmt.Println(primes)
	return int(primes[len(primes)-1])
}
