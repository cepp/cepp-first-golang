package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <- chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
	channel := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)

			read, _ := regexp.Compile("<title>(.*?)<\\/title>")
			channel <- read.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return channel
}

func main() {
	titulo1 := titulo("https://www.google.com", "https://github.com")
	titulo2 := titulo("https://www.linkedin.com", "https://gitlab.com")
	fmt.Println("Primeiros:", <-titulo1, "|", <-titulo2)
	fmt.Println("Segundos:", <-titulo1, "|", <-titulo2)
}
