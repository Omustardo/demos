package solved

import (
	"fmt"
	"github.com/omustardo/demos/euler/figurate"
)

func Problem45() int {
	for h := 144; ; h++ {
		hex := figurate.Hexagonal(h)
		if ok, p := figurate.IsPentagonal(hex); ok {
			if ok, t := figurate.IsTriangular(hex); ok {
				fmt.Printf("Hex#%d = Pent#%d = Tri#%d = %d\n", h, p, t, hex)
				return hex
			}
		}
	}
	return 0
}
