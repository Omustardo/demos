package solved

import (
	"fmt"
	"math"
)

func Problem5() int {
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 // all taken care of by larger factors
	var factors = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i := 21; i < math.MaxInt32; i++ {
		divisible := true
		for _, f := range factors {
			if i%f != 0 {
				divisible = false
				break
			}
		}
		if divisible {
			fmt.Println(i, "is evenly divisible by 1-20")
			return i
		}
	}
	return -1
}
