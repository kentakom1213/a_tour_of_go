package main

import (
	"fmt"
	"math"
)

// ニュートン法を行う
func newton(f, df func(float64) float64, initValue, eps float64, maxIter int) float64 {
	x := initValue
	newX := initValue
	for i := 0; i < maxIter; i++ {
		newX = x - f(x)/df(x)
		if math.Abs(x-newX) < eps {
			return newX
		}
		x = newX
	}
	return newX
}

// ニュートン法により平方根を求める
func sqrt(x float64) float64 {
	return newton(
		func(p float64) float64 { return p*p - x },
		func(p float64) float64 { return 2 * p },
		x,
		1e-7,
		1000,
	)
}

func main() {
	fmt.Printf("√1 = %f\n", sqrt(1))
	fmt.Printf("√2 = %f\n", sqrt(2))
	fmt.Printf("√4 = %f\n", sqrt(4))
	fmt.Printf("√5 = %f\n", sqrt(5))
	fmt.Printf("√100 = %f\n", sqrt(100))
}
