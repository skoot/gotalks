package main

import (
	"fmt"
)

type Hero struct {
	heroName string
	realName string
}

func setRealName(hero *Hero, name string) {
	hero.realName = name
}

func main() {
	batman := Hero{heroName: "Batman"}

	setRealName(&batman, "Bruce Wayne")

	fmt.Println(batman.heroName, "is", batman.realName)
}
