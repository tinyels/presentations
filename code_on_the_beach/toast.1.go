package main

import (
	"fmt"
)

func main() {
	// START OMIT
	messages := make(chan string)
	//sending data
	messages <- "toast is ready"
	//receiving data
	var msg = <-messages
	fmt.Println(msg)
	// END OMIT
}
