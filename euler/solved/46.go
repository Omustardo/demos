package solved

import (
	"github.com/omustardo/demos/euler/prime"
	"math"
)

func Problem46() int {
	var primes []int64

	for i := int64(3); ; i += 2 {
		if prime.IsPrime(i) {
			primes = append(primes, i)
			continue
		}

		proved := false
		for _, p := range primes {
			diff := i - p
			// sum of prime and twice a square, so div by 2 and check if square
			diff /= 2 // note that diff is always even since it's an odd number minus a prime
			root := math.Sqrt(float64(diff))
			if root == math.Floor(root) {
				proved = true
				break
			}
		}
		if !proved {
			return int(i)
		}
	}

	return 0
}
