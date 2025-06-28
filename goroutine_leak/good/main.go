package main

import (
	"fmt"
)

func listen(ch <-chan int) {
	for val := range ch {
		fmt.Println("Got:", val)
	}
	fmt.Println("Channel closed, exiting goroutine.")
}

func main() {
	ch := make(chan int)

	go listen(ch)

	ch <- 1
	ch <- 2
	close(ch) // Zamyka kanał – listen kończy pętlę
}
