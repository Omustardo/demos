package solved

import (
	"fmt"
	"github.com/omustardo/demos/euler/pandigital"
	"strconv"
)

func Problem43() int {

	ssDivisible := func(n int) bool {
		s := fmt.Sprintf("%d", n)
		if len(s) < 4 {
			return true
		}
		arr := []int{2, 3, 5, 7, 11, 13, 17}

		// Subsets of s must be divisible by elements in the array, but since s isn't necessarily len()==10 the
		// number of elements checked needs to be adjusted.
		for i, p := range arr[0 : len(s)-3] {
			n, _ := strconv.Atoi(s[i+1 : i+4]) // 3 digit slices
			if n%p != 0 {
				return false
			}
		}
		return true
	}

	// yield pushes all 10 digit numbers into the given channel.
	yield := func(c chan<- int) {
		defer func() { close(c) }()

		var digits [10]bool // True = digit is in use. All must be true to yield.

		// permute pushes all permutations of 10 digit numbers of the channel.
		// len is the number of base10 digits.
		// digits keeps track of the digits used. True means used, and all must be true to push a number to the channel.
		var permute func(curr, len int, digits [10]bool)
		permute = func(curr, len int, digits [10]bool) {
			if len == 10 {
				c <- curr
				return
			}
			// Reduce total number of checks drastically by checking divisibility early.
			// This improves runtime from ~4.5 seconds to ~33ms
			if !ssDivisible(curr) {
				return
			}

			for d, used := range digits {
				if d == 0 && len == 0 {
					continue // leading zero doesn't result in a 10 digit number
				}
				if !used {
					digits[d] = true
					permute(curr*10+d, len+1, digits)
					digits[d] = false
				}
			}
		}

		permute(0, 0, digits)
	}

	c := make(chan int)
	go yield(c)

	var sum int
	for n := range c {
		if !pandigital.Pandigital(int64(n), 0, 9) {
			continue
		}
		if ssDivisible(n) {
			fmt.Printf("%d is sub-string divisible\n", n)
			sum += n
		}
	}

	return int(sum)
}
