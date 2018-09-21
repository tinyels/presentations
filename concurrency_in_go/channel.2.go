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
}

func main() {
	start := time.Now()
	cake := make(chan string)
	batter := make(chan bool)
	pan := make(chan bool)
	heat := make(chan int)
	go func() {
		do("grease pan", 3)
		pan <- true
	}()
	go func() {
		do("mix ingredients", 15)
		batter <- true
	}()
	go func() {
		do("preheat", 18)
		heat <- 350
	}()
	// START OMIT
	go func() {
		for i := 0; i < 3; i++ {
			select {
			case msg1 := <-heat:
				fmt.Println("temp is", msg1)
			case <-batter:
				fmt.Println("batter is ready")
			case <-pan:
				fmt.Println("pan is ready")
			}
		}
		do("bake", 30)
		cake <- "fudge cake is ready"
	}()
	// END OMIT
	fmt.Println(<-cake)
	fmt.Println("baking took", (time.Since(start) * pace))
}
