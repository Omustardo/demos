package solved

import (
	"fmt"
)

func Problem14() int {
	// number to length of chain
	var known = make(map[int]int)

	count := func(n int) int {
		temp := n
		count := 1 // sequence includes the starting number
		for n != 1 {
			if c, ok := known[n]; ok {
				count += c
				break
			}
			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}
			count++
		}
		known[temp] = count
		return count
	}

	max := 1
	maxN := 1

	for i := 2; i < 1e6; i++ {
		c := count(i)
		if c > max {
			max = c
			maxN = i
		}
		// fmt.Printf("Count of %d is %d\n", i, c)
	}
	fmt.Printf("Max sequence starts with %d and has %d elements\n", maxN, max)
	return maxN
}
