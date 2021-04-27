package solved

import (
	"fmt"
	"math/big"
)

func Problem25() int {

	a := big.NewInt(1) // i=1
	b := big.NewInt(1) // i=2
	i := 2

	// 10^999 is the smallest number with 1000 digits
	limit := big.NewInt(1).Exp(big.NewInt(10), big.NewInt(999), nil)

	// loop while limit is > b
	for limit.Cmp(b) == 1 {
		i++
		c := new(big.Int)
		c.Add(b, a)
		a = b
		b = c
		//if i < 10 || i%100 == 0 {
		//	fmt.Println(b)
		//}
	}

	//fmt.Println(b)
	fmt.Println(i)
	return i
}
