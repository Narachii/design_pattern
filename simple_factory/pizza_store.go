package main

import "fmt"

type pizzaStore struct {
	factory pizzaFactory
}

func newPizzaStore(factory pizzaFactory) *pizzaStore {
	return &pizzaStore{
		factory: factory,
	}
}

func (ps *pizzaStore) orderPizza(pizzaType string) {
	pizza, err := ps.factory.createPizza(pizzaType)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		pizza.prepare()
		pizza.bake()
		pizza.cut()
		pizza.box()
	}
}
