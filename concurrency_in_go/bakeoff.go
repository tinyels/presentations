package main

import (
	"fmt"
	"time"
)

const pace time.Duration = 750

func wait(duration time.Duration) {
	time.Sleep(duration * time.Minute / pace)
}
func do(task string, duration time.Duration, baker string) {
	fmt.Println(baker, "=>", task)
	wait(duration)
	fmt.Println(baker, "<=", task)
}
func bakeCake(baker string) {
	fmt.Println(baker, "is starting")
	go do("preheat oven", 18, baker)
	do("Grease pan", 3, baker)
	do("mix ingredients", 15, baker)
	do("bake", 30, baker)
}
func main() {
	start := time.Now()
	go bakeCake("Nancy")
	go bakeCake("Luis")
	go bakeCake("Martha")
	fmt.Println("baking took", (time.Since(start) * pace))
}
