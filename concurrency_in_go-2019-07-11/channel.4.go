package main

import (
	"fmt"
	"runtime"
	"time"
)

const pace time.Duration = 750

func wait(duration time.Duration) {
	runtime.Gosched() // yield to another goroutine
	time.Sleep(duration * time.Minute / pace)
}

func do(task string, duration time.Duration) {
	fmt.Println("=>", task)
	wait(duration)
	fmt.Println("<=", task)
}

func main() {
	start := time.Now()
	batter := make(chan bool)
	pan := make(chan bool)
	heat := make(chan int)
	go func() {
		do("preheat", 18)
		heat <- 350
	}()
	go func() {
		do("grease pan", 3)
		pan <- true
	}()
	go func() {
		do("mix ingredients", 15)
		batter <- true
	}()
	// START OMIT
	cake := make(chan string)
	go func() {
		msg1 := <-heat
		fmt.Println("temperature is", msg1)
		<-batter
		fmt.Println("batter is ready")
		<-pan
		fmt.Println("pan is ready")
		do("bake", 30)
		cake <- "fudge cake is ready"
	}()
	fmt.Println(<-cake)
	// END OMIT
	fmt.Println("baking took", (time.Since(start) * pace))
}
