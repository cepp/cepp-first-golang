package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Carlos":  11325.45,
		"Eduardo": 15456.78,
		"Pantoja": 1200.0,
	}

	funcsESalarios["Pina"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
