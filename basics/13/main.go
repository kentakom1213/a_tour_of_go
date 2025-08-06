package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4

	// 型変換
	f := math.Sqrt(float64(x*x + y*y))

	z := uint(f)

	fmt.Println(f)
	fmt.Println(x, y, z)
}
