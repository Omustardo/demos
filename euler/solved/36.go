package solved

import "fmt"

func Problem36() int {
	palindrome := func(s string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-i-1] {
				return false
			}
		}
		return true
	}
	palindrome2 := func(n int) bool {
		return palindrome(fmt.Sprintf("%b", n))
	}
	palindrome10 := func(n int) bool {
		return palindrome(fmt.Sprintf("%d", n))
	}

	var sum int
	for i := 0; i < 1e6; i++ {
		if palindrome2(i) && palindrome10(i) {
			sum += i
		}
	}

	return sum
}
