package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("content/go-for-pythonistas/hello/hero.go")
	if err != nil {
		log.Fatalf("file open error: %s", err)
	}

	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	fmt.Println(string(bytes))
}
