package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando ao canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
