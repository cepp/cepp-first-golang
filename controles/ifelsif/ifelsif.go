package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota > 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 8 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else if nota > 10 {
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
