package common

import (
	"fmt"
	"time"
)

const pace time.Duration = 750

func Wait(duration time.Duration) {
	time.Sleep(duration * time.Minute / pace)
}
func Do(task string, duration time.Duration) {
	fmt.Println("=>", task)
	Wait(duration)
	fmt.Println("<=", task)
}
