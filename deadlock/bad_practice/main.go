package main

import (
	"fmt"
	"sync"
	"time"
)

var mu1 sync.Mutex
var mu2 sync.Mutex

func task1() {
	mu1.Lock()
	defer mu1.Unlock()

	time.Sleep(100 * time.Millisecond) // sztuczne opóźnienie
	mu2.Lock()
	defer mu2.Unlock()

	fmt.Println("task1 done")
}

func task2() {
	mu2.Lock()
	defer mu2.Unlock()

	time.Sleep(100 * time.Millisecond)
	mu1.Lock()
	defer mu1.Unlock()

	fmt.Println("task2 done")
}

func main() {
	go task1()
	go task2()

	time.Sleep(1 * time.Second)
}
