package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("batman_sucks.txt")
	if err != nil {
		log.Fatalf("file open error: %s", err)
	}

	fmt.Println(file)
}