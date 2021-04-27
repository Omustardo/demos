package divisor

import (
	"math"
	"sort"
)

var known = map[int64][]int64{
	1: {},
	2: {1},
}

// ProperDivisors returns numbers d between 1 and n-1 for which n%d==0.
// For example, 220 -> [1 2 4 5 10 11 20 22 44 55 110]
func ProperDivisors(n int64) []int64 {
	if n < 0 {
		return nil
	}
	if d, ok := known[n]; ok {
		tmp := make([]int64, len(d))
		copy(tmp, d)
		return tmp
	}

	// Find all divisors < sqrt(n)
	var out []int64
	for i := int64(1); i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			out = append(out, i)
		}
	}
	// Use found divisors to get remaining ones.
	for i := int64(len(out) - 1); i > 0; i-- {
		if out[i]*out[i] != n { // avoid adding square roots twice
			out = append(out, n/out[i])
		}
	}

	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })

	known[n] = out
	tmp := make([]int64, len(out))
	copy(tmp, out)
	return tmp
}

func SummedProperDivisors(n int64) int64 {
	ds := ProperDivisors(n)
	var sum int64
	for _, d := range ds {
		sum += d
	}
	return sum
}

// Abundant numbers have the sum of proper divisors > themselves.
// Returns whether the input is abundant, and the sum of the proper divisors.
func Abundant(n int64) (bool, int64) {
	s := SummedProperDivisors(n)
	return s > n, s
}

// Perfect numbers have the sum of proper divisors == themselves.
func Perfect(n int64) bool {
	return SummedProperDivisors(n) == n
}

// Deficient numbers have the sum of proper divisors < themselves.
// Returns whether the input is deficient, and the sum of the proper divisors.
func Deficient(n int64) (bool, int64) {
	s := SummedProperDivisors(n)
	return s < n, s
}
