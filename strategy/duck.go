package main

import "fmt"

type abstractDuck struct {
	display func()

	flyingBehaviour flyBehaviour
	quackingBehaviour quackBehaviour
}

func (d *abstractDuck) performFly()  {
	d.flyingBehaviour.fly()
}

func (d *abstractDuck) performQuack()  {
	d.quackingBehaviour.quack()
}

func (d *abstractDuck) swim()  {
	fmt.Println("All ducks float, even decoys!")
}

func (d *abstractDuck) setFlyingBehaviour(fb flyBehaviour)  {
	d.flyingBehaviour = fb
}

func (d *abstractDuck) setQuakingBehaviour(qb quackBehaviour)  {
	d.quackingBehaviour = qb
}
