package solved

import (
	"fmt"
	"strconv"
)

func Problem24() int {

	// 0123456789
	// 0123456798 swap 8,9
	// 0123456879
	// 0123456897
	// 0123456978
	// 0123456987
	// 0123457689

	var p []string
	chars := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var gen func(string, []int)
	gen = func(s string, chars []int) {
		if len(s) == 10 {
			p = append(p, s)
			return
		}
		tmp := make([]int, len(chars)-1)
		for i := range chars {
			c := chars[i]
			copy(tmp[:i], chars[:i])
			copy(tmp[i:], chars[i+1:])
			gen(fmt.Sprintf("%s%d", s, c), tmp)
		}
	}
	gen("", chars)

	// fmt.Println(len(p)) // this generates over 3m combinations. We could make gen() stop at 1m, but why bother.
	n, _ := strconv.Atoi(p[999999])
	fmt.Println(n)
	return n
}
