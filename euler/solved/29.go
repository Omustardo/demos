package solved

import (
	"fmt"
	"math/big"
)

func Problem29() int {
	m := make(map[string]bool)
	for a := int64(2); a <= 100; a++ {
		for b := int64(2); b <= 100; b++ {
			n := big.NewInt(0).Exp(big.NewInt(a), big.NewInt(b), nil).String()
			fmt.Println("a^b=", n)
			m[n] = true
		}
	}
	fmt.Println(len(m))
	return len(m)
}
