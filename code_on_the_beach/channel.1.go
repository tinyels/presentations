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
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	// START OMIT
	heat := make(chan bool)
	batter := make(chan bool)
	pan := make(chan bool)
	go func() {
		<-heat
		<-batter
		<-pan
		do("**** bake", 30)
		wg.Done()
	}()
	go func() {
		do("preheat", 18)
		heat <- true
	}()
	go func() {
		do("grease pan", 3)
		pan <- true
	}()
	go func() {
		do("mix ingredients", 15)
		batter <- true
	}()
	// END OMIT
	wg.Wait()
	fmt.Println("baking took", (time.Since(start) * pace))
}
