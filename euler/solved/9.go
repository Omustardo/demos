package solved

import (
	"fmt"
)

func Problem9() int {
	for a := 1; a < 500; a++ {
		for b := a + 1; b < 500; b++ {
			for c := b + 1; c < 500; c++ {
				if a*a+b*b == c*c && a+b+c == 1000 {
					fmt.Println(a * b * c)
					fmt.Println(a, b, c)
					return a * b * c
				}
			}
		}
	}
	return -1
}
