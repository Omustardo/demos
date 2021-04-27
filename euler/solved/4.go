package solved

import (
	"fmt"
)

func Problem4() int {
	isPalindrome := func(n int64) bool {
		s := fmt.Sprintf("%d", n)
		for i := 0; i <= len(s)/2; i++ {
			if s[i] != s[len(s)-i-1] {
				return false
			}
		}
		return true
	}

	hasSmallFactor := func(n int64) bool {
		for i := 2; i < 1000; i++ {
			if n%int64(i) == 0 && n/int64(i) < 1000 {
				fmt.Println("Factors:", i, n/int64(i))
				return true
			}
		}
		return false
	}

	var n int64 = 999 * 999
	for n > 1 {
		if isPalindrome(n) && hasSmallFactor(n) {
			fmt.Println(n)
			return int(n)
		}
		n--
	}
	return -1
}
