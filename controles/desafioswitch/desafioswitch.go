package main

import "fmt"

func notaParaConceito(nota float64) string {
	switch {
	case nota > 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	case nota >= 0 && nota < 3:
		return "E"
	default:
		return "Nota inválida"
	}
}
func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(3))
	fmt.Println(notaParaConceito(2))
	fmt.Println(notaParaConceito(10.1))
}
