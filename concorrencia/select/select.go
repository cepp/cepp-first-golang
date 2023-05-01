package main

import (
	"fmt"
	gohtmltitle "github.com/cepp/go-html-title"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	channel1 := gohtmltitle.Titulo(url1)
	channel2 := gohtmltitle.Titulo(url2)
	channel3 := gohtmltitle.Titulo(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-channel1:
		return t1
	case t2 := <-channel2:
		return t2
	case t3 := <-channel3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		//default:
		//	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido("https://github.com", "https://www.google.com", "https://www.linkedin.com")
	fmt.Println(campeao)
}
