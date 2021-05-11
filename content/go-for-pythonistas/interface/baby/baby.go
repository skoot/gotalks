package baby

import "fmt"

type Baby struct {
	Name string
	Sex string
}

func (b Baby) Reveal() string {
	return fmt.Sprintf("%s is a %s!", b.Name, b.Sex)
}

func (b Baby) Cry() {}
func (b Baby) Sleep() {}
