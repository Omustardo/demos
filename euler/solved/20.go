package solved

import (
	"fmt"
	"math/big"
)

func Problem20() int {
	n := big.NewInt(1)
	
	for i := int64(1); i <= 100; i++ {
		n.Mul(big.NewInt(i), n)
	}
	fmt.Println(n)
	s := fmt.Sprintf("%v",n)
	sum := 0
	for _, c := range s {
		sum += int(c - '0')
	}
	fmt.Println(sum)
return sum
	}

