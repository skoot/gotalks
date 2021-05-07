package main

import "fmt"

func main() {
	// Explicit type declaration.
	var heroName string
	heroName = "Batman"

	// Most of the time though, the type can be inferred
	// and we can use this shorter syntax using the walrus operator.
	realName := "Bruce Wayne"

	fmt.Printf("%s is %s!", heroName, realName)
}
