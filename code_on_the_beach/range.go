package main

import (
	"fmt"
)

func preheat(cs chan int, maxTemp int) {
	for currentTemp := 0; currentTemp <= maxTemp; currentTemp += 25 {
		cs <- currentTemp
	}
	close(cs)
}

func main() {
	cs := make(chan int)
	go preheat(cs, 450)
	for s := range cs {
		fmt.Println("temperature is: ", s)
	}
}
