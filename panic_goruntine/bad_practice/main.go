package main

import (
	"fmt"
	"time"
)

func worker() {
	panic("coś poszło nie tak")
}

func main() {
	go worker()
	time.Sleep(1 * time.Second)
	fmt.Println("Program działa dalej mimo paniki w gorutynie")
}
