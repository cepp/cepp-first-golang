package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, channel chan int) {
	time.Sleep(time.Second)
	channel <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	channel <- 3 * base

	time.Sleep(3 * time.Second)
	channel <- 4 * base
}

func main() {
	channel := make(chan int)
	go doisTresQuatroVezes(2, channel)

	a, b := <-channel, <-channel // recebendo dados do canal
	fmt.Println(a, b)

	fmt.Println(<-channel)

	// chamando a leitura do canal que não tenha retorno gera um deadlock
	//fmt.Println(<-channel)
}
