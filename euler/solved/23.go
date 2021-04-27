package solved

import (
	"fmt"
	"github.com/omustardo/demos/euler/divisor"
)

func Problem23() int {
	// abundant numbers: have sum of proper divisors greater than selves

	// Find all abundant numbers below 28123
	var abund []int64
	for i := int64(12); i < 28123; i++ {
		if ok, _ := divisor.Abundant(i); ok {
			abund = append(abund, i)
		}
	}
	abundSet := make(map[int64]bool)
	for _, a := range abund {
		abundSet[a] = true
	}

	// Find numbers < 28123 that cannot be created by adding two abundant numbers
	var out int64
	for i := int64(1); i <= 28123; i++ {
		abundant := false
		for j := 0; j < len(abund) && abund[j] < i; j++ {
			diff := i - abund[j]
			if _, ok := abundSet[diff]; ok {
				fmt.Printf("ABUNDANT: %d = %d + %d\n", i, abund[j], diff)
				abundant = true
				break
			}
		}
		if !abundant {
			fmt.Printf("NOT ABUN: %d\n", i)
			out += i
		}
	}
	fmt.Println(out)
	return int(out)
}
