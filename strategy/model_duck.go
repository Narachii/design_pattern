package main

import "fmt"

type modelDuck struct {
	*abstractDuck
}
func newModelDuck() *modelDuck {
	d := &abstractDuck{
		display: func() {
			fmt.Println("I am a real Model duck")
		},
		flyingBehaviour: &flyNoWay{},
		quackingBehaviour: &muteQuack{},
	}

	return &modelDuck{d}
}
