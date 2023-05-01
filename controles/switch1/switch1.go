package main

import "fmt"

func notaParaConceito(notaFloat float64) string {
	nota := int(notaFloat)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(7.9))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(3.1))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(11.1))
}
