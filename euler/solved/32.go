package solved

import "fmt"

func Problem32() int {

	products := make(map[int]bool)

	shareDigits := func(a, b int) bool {
		var seen [10]bool
		for _, n := range []int{a, b} {
			for n > 0 {
				if seen[n%10] {
					return true
				}
				seen[n%10] = true
				n /= 10
			}
		}
		return false
	}

	pandigital := func(a, b, c int) bool {
		var seen [10]bool
		for _, n := range []int{a, b, c} {
			for n > 0 {
				if seen[n%10] {
					return false
				}
				seen[n%10] = true
				n /= 10
			}
		}
		if seen[0] {
			return false
		}
		for i := 1; i < 10; i++ {
			if !seen[i] {
				return false
			}
		}
		return true
	}

	for a := 2; a < 5000; a++ {
		for b := a + 1; b < 5000; b++ {
			if shareDigits(b, a) {
				continue
			}
			if b%10 == 1 {
				continue
			}
			c := a * b
			if !products[c] && pandigital(a, b, c) {
				products[c] = true
				fmt.Printf("%d * %d = %d\n", a, b, c)
			}
		}
	}

	var sum int
	for p := range products {
		sum += p
	}
	fmt.Println("Sum: ", sum)
	return sum
}
