package main

import (
	"fmt"
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
	heat := make(chan bool)
	go func() {
		do("preheat", 18)
		heat <- true
	}()
	go do("grease pan", 3)
	go do("mix ingredients", 15)
	<-heat
	do("bake", 30)
	fmt.Println("baking took", (time.Since(start) * pace))
	// END OMIT
}
