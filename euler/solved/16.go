package solved

import (
	"fmt"
	"math/big"
)


func Problem16() int {
	n := big.NewInt(2)
	n.Exp(big.NewInt(2), big.NewInt(1000), nil)
	fmt.Println(n)
	
	sum := 0
	for _, c := range fmt.Sprintf("%v", n) {
		sum += int(c - '0')
	}
	fmt.Printf("Sum: %d\n", sum)
	return sum
}
