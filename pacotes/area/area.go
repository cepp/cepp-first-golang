package area

import "math"

// Pi é uma proporção numérica definida plea relação entre o perímetro de um circunferência e seu diâmetro
const Pi = 3.1416

// Circ calcula a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _TrianguloEq(base, altura float64) float64 {
	return Rect(base, altura) / 2
}
