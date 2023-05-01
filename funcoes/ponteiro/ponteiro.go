package main

import "fmt"

func inc1(numero int) {
	numero++
}

func inc2(numero *int) {
	*numero++
}

func main() {
	numero := 1

	inc1(numero)
	fmt.Println(numero)

	inc2(&numero)
	fmt.Println(numero)
}
