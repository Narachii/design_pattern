package main

import "fmt"

type flyBehaviour interface {
	fly()
}

type flyWithWings struct {}

func (fw *flyWithWings) fly()  {
	fmt.Println("I am flying")
}

type flyNoWay struct {}

func (fn *flyNoWay) fly()  {
	fmt.Println("I can't fly")
}

type flyRocketPowered struct {}

func (f *flyRocketPowered) fly()  {
	fmt.Println("I am flying with a rocket!")
}
