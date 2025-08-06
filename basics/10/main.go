package main

import "fmt"

func main() {
	var i, j int = 1, 2

	// var構文と同じ．関数内でのみ使える
	k := 3

	c, python, java := true, false, "No"

	fmt.Println(i, j, k, c, python, java)
}
