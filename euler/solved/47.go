package solved

import (
	"github.com/omustardo/demos/euler/prime"
)

func Problem47() int {
	// how many distinct factors to look for, in the same number of consecutive integers
	distinct := 3

	//fmt.Println(prime.Factors(23 * 16))
	//fmt.Println(prime.Factors(23 * 16 * 2152))
	//fmt.Println(prime.Factors(3 * 3 * 2 * 23))

	factors := make(map[int64]int) // all factors of the set of four consecutive numbers
	add := func(m map[int64]int, f []int64) {
		for _, v := range f {
			m[v]++
		}
	}
	clear := func(m map[int64]int) {
		for k := range m {
			delete(m, k)
		}
	}

	for i := 2; ; i++ {
		for j := 0; j < distinct; j++ {
			if f := prime.Factors(int64(i + j)); len(f) != distinct {
				clear(factors)
				break
			} else {
				add(factors, f)
				if len(factors) != (j+1)*distinct {
					clear(factors)
					break
				}
			}
		}
		if len(factors) == distinct*distinct {
			return i
		}
	}
	return 0
}
