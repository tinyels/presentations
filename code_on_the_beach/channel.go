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

func main() {
	// START OMIT
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	heat := make(chan bool)
	go func() {
		do("preheat", 18)
		heat <- true
	}()
	go do("grease pan", 3)
	go do("mix ingredients", 15)
	go func() {
		<-heat
		do("**** bake", 30)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("baking took", (time.Since(start) * pace))
	// END OMIT
}
