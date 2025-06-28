package main

import (
	"fmt"
)

func task1() {
	mu1.Lock()
	mu2.Lock()
	fmt.Println("task1 done")
	mu2.Unlock()
	mu1.Unlock()
}

func task2() {
	mu1.Lock()
	mu2.Lock()
	fmt.Println("task2 done")
	mu2.Unlock()
	mu1.Unlock()
}
