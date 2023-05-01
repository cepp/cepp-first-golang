package main

import "fmt"

func main() {
	// mapas devem ser inicializados
	//var aprovados := map[int]string
	aprovados := make(map[int]string)

	aprovados[12345678891] = "Maria"
	aprovados[98765432109] = "Pedro"
	aprovados[87538903773] = "Ana"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[87538903773])
	delete(aprovados, 87538903773)
	fmt.Println(aprovados[87538903773])
}
