package main

import (
	"fmt"
)

type Hero struct {
	heroName string
	realName string
}

func (h Hero) reveal() string {
	return fmt.Sprintf("%s is %s", h.heroName, h.realName)
}

func main() {
	batman := Hero{heroName: "Batman", realName: "Bruce Wayne"}

	fmt.Println(batman.reveal())
}
