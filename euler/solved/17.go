package solved

import (
	"fmt"
	"bytes"
)

func Problem17() int {
	var (
		baseWords = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
			"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
		tens = []string{"zero", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	)
	
	var toWord func(int) string
	toWord = func(n int) string {
		if n == 1000 {
		return "onethousand"
	}
		if n < 20 {
		return baseWords[n]
	}
		if n < 100 && n % 10 == 0 {
		return tens[n/10]
	}
		if n >= 100 {
		if n % 100 == 0 {
		return baseWords[n/100] + "hundred"
	}
		return baseWords[n/100] + "hundredand" + toWord(n % 100)
	}
		if n % 10 == 0 {
		return tens[n / 10]
	}
		return tens[n/10] + baseWords[n%10]
	}
	
	var buf bytes.Buffer
	for i := 1; i <= 1000; i+=1 {
		w := toWord(i)
		fmt.Println(w)
		buf.WriteString(w)
	}
	fmt.Println(buf.Len())
	return buf.Len()
}

