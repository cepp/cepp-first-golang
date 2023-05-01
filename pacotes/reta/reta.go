package main

import "math"

// Inicializando com a letra maiúscula a instrução passa a ser pública (visível dentro e fora do pacote)
// Inicializando com letra minúscula é privado (visível apenas dentro do pacote)

// Exemplos
// Ponto - será público
// ponto ou _Ponto - será privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
