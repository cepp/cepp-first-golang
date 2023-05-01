package main

import (
	"fmt"
	title "github.com/cepp/go-html-title"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	channel := make(chan string)
	go encaminhar(entrada1, channel)
	go encaminhar(entrada2, channel)
	return channel
}

func main() {
	channel := juntar(
		title.Titulo("https://www.google.com", "https://github.com"),
		title.Titulo("https://www.linkedin.com", "https://gitlab.com"),
	)
	fmt.Println(<-channel, "|", <-channel)
	fmt.Println(<-channel, "|", <-channel)
}
