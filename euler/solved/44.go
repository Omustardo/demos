package solved

import ()

func Problem44() int {

	// The formula for pentagonal numbers is f(n)=n(3n-1)/2
	// But it's simpler to look at it like this:
	// Sequence starts with 1, 5 and each number after that had 3*n+1 added to it.
	// So the differences between numbers in the sequence are: 4, 7, 10, 13, 16, 19, 22...

	pent := func(n int) int {
		return (n*n*3 - n) / 2
	}
	pentA := []int{0}
	pentM := make(map[int]bool)
	for i := 1; i < 10000; i++ {
		pentA = append(pentA, pent(i))
		pentM[pentA[i]] = true
	}

	for i := 1; i < len(pentA); i++ {
		for j := i + 1; j < len(pentA); j++ {
			diff := pentM[pentA[j]-pentA[i]]
			sum := pentM[pentA[j]+pentA[i]]
			if diff && sum {
				// fmt.Printf("[%d %d]: [%d %d] : %d\n", i, j, pentA[i], pentA[j], pentA[j]-pentA[i])
				return pentA[j] - pentA[i]
			}
		}
	}

	return 0
}
