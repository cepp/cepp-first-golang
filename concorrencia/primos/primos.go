package main

import (
	"fmt"
	"time"
)

func isPrimo(numero int) bool {
	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func primos(numero int, channel chan int) {
	inicio := 2
	for i := 0; i < numero; i++ {
		for primo := inicio; ; primo++ {
			channel <- primo
			inicio = primo + 1
			time.Sleep(time.Millisecond * 100)
			break
		}
	}
	close(channel)
}

func main() {
	channel := make(chan int, 30)
	go primos(cap(channel), channel)
	for primo := range channel {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("\nFim!")
}
