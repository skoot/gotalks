package main

import "fmt"

func repeat(s string, n int) chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			c <- s
		}
	}()
	return c
}

func main() {
	for n := range repeat("na ", 10) {
		fmt.Print(n)
	}
	fmt.Println("Batman!")
}
