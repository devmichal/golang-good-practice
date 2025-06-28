package main

import (
	"fmt"
	"time"
)

func listen(ch <-chan int) {
	for val := range ch {
		fmt.Println("Got:", val)
	}
}

func main() {
	ch := make(chan int)

	go listen(ch) // Gorutyna czeka na dane, ale ich nie dostanie

	// main kończy się bez zamknięcia kanału ani wysłania czegokolwiek
	time.Sleep(1 * time.Second)
}
