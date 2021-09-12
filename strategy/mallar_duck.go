package main

import "fmt"

type mallarDuck struct {
	*abstractDuck
}

func newMallarDuck() *mallarDuck {
	d := &abstractDuck{
		display: func() {
			fmt.Println("test")
		},
		flyingBehaviour: &flyWithWings{},
		quackingBehaviour: &quack{},
	}

	return &mallarDuck{d}
}
