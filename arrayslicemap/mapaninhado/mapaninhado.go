package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"C": {
			"Carlos Pina": 15456.78,
		},
		"E": {
			"Eduardo Pantoja": 10000.90,
		},
		"P": {
			"Pina Pantoja": 20000.90,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
