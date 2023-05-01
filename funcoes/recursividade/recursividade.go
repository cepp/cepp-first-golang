package main

import "fmt"

func fatorial(numero int) (int, error) {
	switch {
	case numero < 0:
		return -1, fmt.Errorf("número inválido: %d", numero)
	case numero == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(numero - 1)
		return numero * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
