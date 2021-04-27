package solved

import "fmt"

func Problem28() int {

	sum := 1
	start := 1
	for dims := 3; dims <= 1001; dims += 2 {
		sum += start + 1*(dims-1)
		sum += start + 2*(dims-1)
		sum += start + 3*(dims-1)
		sum += start + 4*(dims-1)

		start = start + 4*(dims-1)
		//fmt.Println("Ending at", start)
		//fmt.Println("Sum:", sum)
		//fmt.Println()
	}

	fmt.Println(sum)
	return sum
}
