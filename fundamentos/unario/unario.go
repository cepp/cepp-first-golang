package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas forma postfix
	x++
	y--

	fmt.Printf("X: %v, Y: %v", x, y)
}
