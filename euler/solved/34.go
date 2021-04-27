package solved

import "fmt"

func Problem34() int {
	// pre-compute factorials for 0 through 9.
	factorial := make(map[int]int)
	factorial[0] = 1 // 0! = 1
	for i := 1; i < 10; i++ {
		factorial[i] = factorial[i-1] * i
	}

	// factorialSum returns the sum of the factorial of the digits of the given number.
	// e.g. (145) -> 145 because that's equal to (1! + 4! + 5!)
	factorialSum := func(n int) int {
		sum := 0
		m := n
		for m > 0 {
			sum += factorial[m%10]
			m /= 10
		}
		return sum
	}

	// based on problem statement, 1 and 2 are not included
	var found []int
	for i := 3; i < 1e5; i++ {
		if factorialSum(i) == i {
			found = append(found, i)
			fmt.Println(i)
		}
	}
	fmt.Println(found)

	var sum int
	for _, n := range found {
		sum += n
	}
	return sum
}
