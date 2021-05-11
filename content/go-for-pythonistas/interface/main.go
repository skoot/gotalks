package main

import (
	"fmt"
	"github.com/skoot/gotalks/content/go-for-pythonistas/interface/baby"
	"github.com/skoot/gotalks/content/go-for-pythonistas/interface/hero"
)

type revealer interface {
	Reveal() string
}

func display(r revealer) {
	fmt.Println(r.Reveal())
}

func main() {
	batman := hero.Hero{HeroName: "Batman", RealName: "Bruce Wayne"}
	olivier := baby.Baby{Name: "Olivier", Sex: "boy"}
	display(batman)
	display(olivier)
}
