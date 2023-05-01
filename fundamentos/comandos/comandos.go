package main

import "fmt"

func main() {
	/**
	expõe a documentação do go na porta 6060
	godoc -http=:6060
	*/
	/**
	verifica se o arquivo possui algum problema
	go vet arquivo.go
	*/
	fmt.Printf("Outro program em %s!", "Golang")
}
