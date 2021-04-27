package prime

import (
	"sort"
)

var (
	primeSlice = []int64{2, 3, 5}
	primeSet   = map[int64]bool{2: true, 3: true, 5: true}

	factorCache = make(map[int64]map[int64]int64)
)

func IsPrime(n int64) bool {
	if n <= 1 {
		return false
	}

	// If it's a known prime, return immediately.
	if _, ok := primeSet[n]; ok {
		return true
	}

	for _, p := range primeSlice {
		// If the input is divisible by a known prime, then it's not prime by definition.
		if n%p == 0 {
			return false
		}
		// If the input is less than the square of a known prime, then it can't be prime.
		if p*p >= n {
			break
		}
	}

	// If prime is unknown, generate primes up to and above it. Only up to would be enough, but using
	// going above mitigates the situation where user asks for ever increasing primes.
	// For example, IsPrime(1000); IsPrime(1100); IsPrime(1200) would each re-regenerate all primes.
	generate(int64(n * 2))

	return primeSet[n]
}

// generates all primes <= n
func generate(n int64) {
	primeSlice = nil
	b := make([]bool, n+1)
	for i := int64(2); i <= n; i++ {
		if b[i] == true {
			continue
		}
		primeSlice = append(primeSlice, i)
		primeSet[i] = true
		for k := i * i; k <= n; k += i {
			b[k] = true
		}
	}
}

// List returns all prime numbers <= n.
func List(max int64) []int64 {
	// generate primes in case they weren't calculated yet. If already generated this is a no-op.
	generate(max)

	i := sort.Search(len(primeSlice), func(i int) bool { return primeSlice[i] >= max })
	// Copy to avoid callers modifying internal prime data.
	out := make([]int64, i)
	copy(out, primeSlice)
	return out
}

func add(m map[int64]int64, f map[int64]int64) {
	for k := range f {
		m[k]++
	}
}

func dupe(m map[int64]int64) map[int64]int64 {
	n := make(map[int64]int64)
	for k, v := range m {
		n[k] = v
	}
	return n
}

// Factors returns all of the prime factors of the given number, mapped to their counts.
//func Factors(n int64) map[int64]int64 {
//	// generate primes in case they weren't calculated yet. If already generated this is a no-op.
//	generate(n)
//
//	// already found factors
//	if m, ok := factorCache[n]; ok {
//		return dupe(m)
//	}
//
//	curr := n
//	m := make(map[int64]int64)
//	for _, p := range primeSlice {
//		if p > curr {
//			break
//		}
//		if f, ok := factorCache[curr]; ok {
//			add(m, f)
//			factorCache[n] = m
//			return dupe(m)
//		}
//
//		for p <= curr && curr%p == 0 {
//			curr /= p
//			m[p]++
//		}
//	}
//	factorCache[n] = m
//	return dupe(m)
//}

// https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/
// Get all prime factors of a given number n
func Factors(n int64) []int64 {
	var out []int64

	// Get the number of 2s that divide n
	for n%2 == 0 {
		out = append(out, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := int64(3); i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			out = append(out, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		out = append(out, n)
	}
	return out
}
