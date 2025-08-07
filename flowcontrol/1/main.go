package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0.0

	for i := 1; i <= 10000; i++ {
		sum += 1 / float64(i*i)

		if i%1000 == 0 {
			fmt.Printf("i = %d: %f\n", i, sum)
		}
	}

	// ans
	ans := math.Pi * math.Pi / 6.0

	fmt.Printf("π^2 / 6 = %f\n", ans)
}
