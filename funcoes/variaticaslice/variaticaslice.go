package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	listaAprovados := []string{"Maria", "Joaquina", "De Amaral", "Pereira", "Góes"}
	imprimirAprovados(listaAprovados...)
}
