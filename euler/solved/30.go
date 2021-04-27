package solved

import (
	"fmt"
	"math"
)

func Problem30() int {
	// Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.

	const numDigits = 5

	sumOfDigitPowers := func(n int) bool {
		tmp := n
		var sum int
		for n > 0 {
			digit := n % 10
			n /= 10
			sum += int(math.Pow(float64(digit), numDigits))
		}
		return sum == tmp
	}

	sum := 0
	// check numbers up to 400k. The highest can't be above ~350k because 999999 -> (9^5)*6==354k
	for i := 2; i < 400000; i++ {
		if sumOfDigitPowers(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println("Sum:", sum)
	return sum
}
