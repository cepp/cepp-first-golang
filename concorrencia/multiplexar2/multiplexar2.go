package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			channel <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()
	return channel
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				channel <- s
			case s := <-entrada2:
				channel <- s
			}
		}
	}()
	return channel
}

func main() {
	channel := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-channel, <-channel)
	fmt.Println(<-channel, <-channel)
	fmt.Println(<-channel, <-channel)
}
