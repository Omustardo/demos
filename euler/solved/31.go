package solved

import "fmt"

func Problem31() int {
	pence := []int{1, 2, 5, 10, 20, 50, 100, 200}
	target := 200

	var make200 func(sum, i int) int
	make200 = func(sum, i int) int {
		if sum == target {
			return 1
		}
		if sum > target {
			return 0
		}
		count := 0
		for j := i; j < len(pence); j++ {
			count += make200(sum+pence[j], j)
		}
		return count
	}

	count := make200(0, 0)
	fmt.Println(count)
	return count
}
