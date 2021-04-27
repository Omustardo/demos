package solved

import "fmt"

func Problem33() int {

	// shareDigit takes two 2-digit integers that are not equal.
	// It returns the single digit they have in common
	shareDigit := func(a, b int) (int, bool) {
		if a < 10 || b < 10 || a > 99 || b > 99 || a == b {
			fmt.Println("Inputs must be unequal 2 digit integers. Got:", a, b)
			return 0, false
		}
		if a%10 == b%10 {
			return a % 10, true
		}
		if a/10 == b%10 {
			return a / 10, true
		}
		if a%10 == b/10 {
			return a % 10, true
		}
		if a/10 == b/10 {
			return a / 10, true
		}
		return 0, false
	}

	// removeDigit takes a two digit integer and a single digit integer. It removes the one digit integer from the
	// two digit integer.
	// For example: (51, 5) -> 1
	removeDigit := func(a, n int) int {
		if a < 10 || a > 99 || n > 9 || n < 0 {
			fmt.Println("Inputs must be a two digit integer and a single digit integer. Got:", a, n)
			return 0
		}
		if !(a%10 == n || a/10 == n) {
			fmt.Printf("%d expected to be a digit in %d\n", n, a)
			return a
		}
		if a%10 == n {
			return a / 10
		}
		if a/10 == n {
			return a % 10
		}
		panic("should never happen")
	}

	type Fraction struct {
		a, b int
	}
	found := make(map[Fraction]bool)

	for a := 10; a < 100; a++ {
		for b := a + 1; b < 100; b++ {
			if a%10 == 0 || b%10 == 0 {
				continue // only non-trivial examples are desires (so 10/20 == 1/2 isn't what we care about)
			}
			if n, ok := shareDigit(a, b); ok {
				c := removeDigit(a, n)
				d := removeDigit(b, n)

				if float32(a)/float32(b) == float32(c)/float32(d) {
					fmt.Printf("Found %d/%d == %d/%d\n", a, b, c, d)
					found[Fraction{a, b}] = true
				}
			}
		}
	}
	if len(found) != 4 {
		fmt.Println("Expected 4 fractions, got:", found)
	}
	product := Fraction{1, 1}
	for f := range found {
		product.a *= f.a
		product.b *= f.b
	}

	fmt.Printf("Product is: %v\n", product)
	// reduce to lowest common terms
	for reduced := true; reduced; {
		reduced = false
		for factor := 2; factor <= product.a; factor++ {
			if product.a%factor == 0 && product.b%factor == 0 {
				product.a /= factor
				product.b /= factor
				fmt.Println("Reduced by dividing by", factor)
				reduced = true
			}
		}
	}
	fmt.Printf("Reduced Product is: %v\n", product)
	return product.b
}
