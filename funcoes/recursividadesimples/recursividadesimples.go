package main

import "fmt"

func fatorial(numero uint) uint {
	switch {
	case numero == 0:
		return 1
	default:
		return numero * fatorial(numero-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
