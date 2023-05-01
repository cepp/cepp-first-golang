package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iterações %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// função sendo chamada não concorrente
	//fale("Maria", "Por que você não fala comigo?", 3)
	//fale("João", "Estou esperando minha vez", 1)

	// go routines são executadas até que o programa esteja sendo executado
	//go fale("Maria", "Ei...", 500)
	//go fale("João", "Opa...", 500)

	// adicionando sleep para que dê tempo que as go routines sejam executadas
	go fale("Maria", "Entendi...", 10)
	go fale("João", "Parabéns", 5)
	time.Sleep(time.Second * 5)
	fmt.Println("Fim!")
}
