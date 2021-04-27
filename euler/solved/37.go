package solved

import (
	"fmt"
	"github.com/omustardo/demos/euler/prime"
	"math"
)

// We know the prime numbers must both start and end with digits 3,5,7 or they wouldn't be prime.
// The middle digits can be any of: 1,3,7,9

func Problem37() int {

	var sum int
	for _, p := range prime.List(1e6) {
		if p < 10 {
			continue // 2,3,5,7 not considered truncatable primes
		}

		left, right := p, p
		for i := 0; i < len(fmt.Sprintf("%d", p)); i++ {
			if !prime.IsPrime(left) || !prime.IsPrime(right) {
				break
			}
			left /= 10
			right = p % int64(math.Pow(10, float64(i+1)))
		}
		if left == 0 {
			fmt.Printf("%d is a truncatable prime\n", p)
			sum += int(p)
		}
	}

	return sum
}
