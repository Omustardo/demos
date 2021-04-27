package solved

import (
	"fmt"
	"time"
	
	"github.com/omustardo/demos/euler/divisor"
)

func Problem21() int {
	// n->sum(divisors(n))
	// if arr[arr[n]] == arr[n] then they are amicable
	
	fmt.Println(divisor.ProperDivisors(25))
	fmt.Println(divisor.ProperDivisors(6))
	fmt.Println(divisor.ProperDivisors(652))
	fmt.Println(divisor.ProperDivisors(496))
	
	arr := make([]int64, 10001)
	for n := int64(0); n < 10000; n++ {
		ds := divisor.ProperDivisors(n)
		var sum int64
		for _, d := range ds {
			sum += d
		}
		fmt.Println(n,":", sum, ":", ds)
		arr[n] = sum
	}
	
	var sum int64
	for i := int64(2); i < 10000; i++ {
		if arr[i] < 10000 && arr[i] < i && i == arr[arr[i]] {
			fmt.Println(i, ",", arr[i], "are amicable")
			sum += i + arr[i]
		}
	}
	
	fmt.Println("Total:", sum)
	return int(sum)
}

func main() {
	start := time.Now()
	Problem21()
	fmt.Printf("Elapsed Time: %v", time.Now().Sub(start))
}

