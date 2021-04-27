package solved

import (
	"fmt"
)

func Problem6() int {
	var sum0, sum1 int64
	for i := int64(1); i <= 100; i++ {
		sum0 += i * i
		sum1 += i
	}
	sum1 *= sum1
	fmt.Println(sum1 - sum0)
	return int(sum1 - sum0)
}
