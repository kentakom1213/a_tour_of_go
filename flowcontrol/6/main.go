package main

import (
	"fmt"
	"math"
)

func saturating_pow(x, n, lim float64) float64 {
	// vを定義
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		saturating_pow(3, 2, 10),
		saturating_pow(3, 3, 10),
	)
}
