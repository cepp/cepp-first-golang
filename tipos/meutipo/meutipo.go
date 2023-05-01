package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(nota nota) string {
	if nota.entre(9.0, 10.0) {
		return "A"
	} else if nota.entre(7.0, 8.99) {
		return "B"
	} else if nota.entre(5.0, 7.99) {
		return "C"
	} else if nota.entre(3.0, 4.99) {
		return "D"
	} else if nota > 10.0 {
		return "Nota inválida"
	}
	return "E"
}

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(3))
	fmt.Println(notaParaConceito(2))
	fmt.Println(notaParaConceito(10.1))
}
