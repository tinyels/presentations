package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func funcA() {
	mu.Lock()
	fmt.Println("Hello, World")
	mu.Unlock()
}

func main() {
	mu.Lock()
	funcA()
	mu.Unlock()
}
