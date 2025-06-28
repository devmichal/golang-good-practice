package main

import (
	"fmt"
	"time"
)

func worker() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Odzyskano z paniki:", r)
		}
	}()

	panic("coś poszło nie tak")
}

func main() {
	go worker()
	time.Sleep(1 * time.Second)
	fmt.Println("Program działa dalej, gorutyna obsłużyła panikę")
}
