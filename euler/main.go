package main

import (
	"fmt"
	"time"

	"github.com/omustardo/demos/euler/solved"
)

func main() {
	start := time.Now()

	go func() {
		i := 1
		for {
			time.Sleep(time.Second * 10)
			fmt.Printf("%d seconds elapsed\n", i*10)
			i++
		}
	}()

	fmt.Println("Solution:", solved.Problem47())

	fmt.Printf("Elapsed Time: %v", time.Now().Sub(start))
}
