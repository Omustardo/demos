package solved

import (
	"fmt"
)

func Problem40() int {
	// This problem has nothing to do with decimals really. Just count the digits.
	// 9 1's
	//  9 so far
	// 90 10's
	//  9+90*2 so far
	// 900 3 digit numbers
	//  9+90*2+900*3
	// 9000 4 digit numbers
	//  9+90*2+900*3+9000*4

	// or just do it directly:
	var buf []rune
	var i int
	for len(buf) <= 1e6 {
		buf = append(buf, []rune(fmt.Sprintf("%d", i))...)
		i++
	}

	prod := 1
	for i := 1; i <= 1e6; i *= 10 {
		d := buf[int(i)] - '0'
		fmt.Printf("i: %d\n", d)
		prod *= int(d)
	}

	return prod
}
