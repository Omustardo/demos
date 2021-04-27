package solved

import ()

// a,b,c are integers
// a + b + c = p
// a^2 + b^2 = c^2
// right triangle, so sin/cos/tan might be handy?

func Problem39() int {

	type Triangle struct {
		a, b, c int
		p       int
	}

	const maxPerimeter = 1000

	var solutions []Triangle
	for p := 12; p <= maxPerimeter; p++ {
		var curSolutions []Triangle

		for a := 1; a < maxPerimeter; a++ {
			for b := a + 1; b < maxPerimeter; b++ {
				c := p - (a + b)
				if c < 0 {
					continue // not a triangle
				}
				if c*c != a*a+b*b {
					continue // not right triangle
				}
				curSolutions = append(curSolutions, Triangle{a, b, c, p})
			}
		}

		if len(curSolutions) > len(solutions) {
			// fmt.Printf("Perimeter %d has %d solutions: %v\n", p, len(curSolutions), curSolutions)
			solutions = curSolutions
		}
	}

	return len(solutions)
}
