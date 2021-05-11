package hero

import "fmt"

type Hero struct {
	HeroName string
	RealName string
}

func (h Hero) Reveal() string {
	return fmt.Sprintf("%s is %s...", h.HeroName, h.RealName)
}

func (h Hero) Fight() {}
