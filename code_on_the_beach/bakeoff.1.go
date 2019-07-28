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
	do("**** bake", 30, baker)
}
func main() {
	start := time.Now()
	var wg sync.WaitGroup // HL
	wg.Add(3)             // HL
	go func() {
		bakeCake("Nancy")
		wg.Done() // HL
	}()
	go func() {
		bakeCake("Luis")
		wg.Done() // HL
	}()
	go func() {
		bakeCake("Martha")
		wg.Done() // HL
	}()
	wg.Wait() // HL
	fmt.Println("baking took", (time.Since(start) * pace))
}
