package main

import (
	"fmt"
	"time"
)

func display(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go display("Batman")
	go display("Jocker")
	time.Sleep(100 * time.Millisecond)
}
