package main

import (
	"fmt"
	"sync"
	"time"
)

const pace time.Duration = 750

func wait(duration time.Duration) {
	time.Sleep(duration * time.Minute / pace)
}

func do(task string, duration time.Duration) {
	fmt.Println("=>", task)
	wait(duration)
	fmt.Println("<=", task)
}

// START OMIT
var oven sync.Mutex

func useOven(task string, duration time.Duration) {
	oven.Lock()
	do(task, duration)
	oven.Unlock()
}

func main() {
	start := time.Now()
	go useOven("preheat oven", 18)
	do("Grease pan", 3)
	do("mix ingredients", 15)
	useOven("bake", 30)
	fmt.Println("baking took", (time.Since(start) * pace))
}

// END OMIT
