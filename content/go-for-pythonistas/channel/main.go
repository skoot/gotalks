package main

import (
	"fmt"
	"time"
)

func displayAll(numbers chan int, name string) {
	for s := range numbers {
		time.Sleep(10*time.Millisecond)
		fmt.Println(name, s)
	}
}

func main() {
	numbers := make(chan int)
	go displayAll(numbers, "Batman")
	go displayAll(numbers, "Robin")
	for i:= 0; i < 5; i++ {
		numbers <- i
	}
	time.Sleep(100*time.Millisecond)
}
