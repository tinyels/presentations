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

func block(dead chan string) {
	//	dead <- "is anyone there"
}

func main() {
	cs := make(chan int)
	dead := make(chan string)
	go preheat(cs, 450)
	go block(dead)
	for s := range cs {
		fmt.Println("temperature is: ", s)
	}

	fmt.Println(<-dead)
}
