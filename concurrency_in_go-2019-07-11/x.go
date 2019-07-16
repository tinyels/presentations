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
	start := time.Now()
	do("preheat oven", 18)
	do("Grease pan", 3)
	do("mix ingredients", 15)
	do("bake", 30)
	fmt.Println("baking took", (time.Since(start) * pace))
}
