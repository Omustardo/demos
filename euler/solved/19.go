package solved

import (
	"fmt"
)

func Problem19() int {
	iter := func() func() (int, int, int) {
		year := 1901
		monthDays := []int{
			31, 28, 31, // J F M
			30, 31, 30, // A M J
			31, 31, 30, // J A S
			31, 30, 31, // O N D
		}
		var mo int
		day := 0

		return func() (int, int, int) {
			outD, outM, outY := day, mo, year

			day += monthDays[mo]
			if mo == 1 { // Feb
				if year%4 == 0 { // standard leap year
					if year%100 != 0 || year%400 == 0 { // ignore century years, except if divisible by 400
						day++
						fmt.Println("Leap Year accounted for in", year)
					}
				}
			}
			mo++

			if mo == 12 {
				mo = 0
				year++
			}

			return outD, outM, outY
		}
	}

	// Assume Jan 1 1901 is considered day 0 (it's a Tueday).
	i := iter()
	var total int
	day, mo, year := i()
	for year < 2001 {
		if (day+2)%7 == 0 {
			total++
			fmt.Printf("Got a 1st Sunday: %d/%d\n", mo, year)
		}
		fmt.Printf("%d %d\n", mo, year)
		day, mo, year = i()
	}
	fmt.Println(total)
	return total
}
