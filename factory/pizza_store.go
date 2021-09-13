package main

import "fmt"

// Abstract Interface
type iPizzaStore interface {
	orderPizza(pizzaType string) iPizza
	createPizza(pizzaType string) (iPizza, error)
}

// Abstract Concrete Type
type aPizzaStore struct {
	createPizza func(pizzaType string) (iPizza, error)
}

func (a *aPizzaStore) orderPizza(pizzaType string) iPizza {
	pizza, err := a.createPizza(pizzaType)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		pizza.prepare()
		pizza.bake()
		pizza.cut()
		pizza.box()
	}

	return pizza
}

/**
 * Each subclass implements the createPizza() method
 * and make use of the orderPizza() method in the parent class
 */
type nyPizzaStore struct {
	*aPizzaStore
}

func newNYPizzaStore() iPizzaStore {
	basePizzaStore := &aPizzaStore{}

	nyPizzaStore := &nyPizzaStore{basePizzaStore}
	nyPizzaStore.aPizzaStore.createPizza = nyPizzaStore.createPizza

	return nyPizzaStore
}

func (ps *nyPizzaStore) createPizza(pizzaType string) (iPizza, error) {
	var pizza iPizza
	switch pizzaType {
	case "cheese":
		pizza = newNYStyleCheesePizza()
		//	case "veggie":
		//		pizza = newNYStyleVeggiePizza()
		//	case "pepperoni":
		//		pizza = newNYStylePepperoniPizza()
	}

	if pizza == nil {
		return nil, fmt.Errorf("Invalid pizza type")
	}

	return pizza, nil
}

type chicagoPizzaStore struct {
	*aPizzaStore
}

func newChicagoPizzaStore() iPizzaStore {
	basePizzaStore := &aPizzaStore{}
	chicagoPizzaStore := &chicagoPizzaStore{basePizzaStore}
	chicagoPizzaStore.aPizzaStore.createPizza = chicagoPizzaStore.createPizza

	return chicagoPizzaStore
}

func (cp *chicagoPizzaStore) createPizza(pizzaType string) (iPizza, error) {
	var pizza iPizza
	switch pizzaType {
	case "cheese":
		pizza = newChicagoStyleCheesePizza()
		//	case "veggie":
		//		pizza = newNYStyleVeggiePizza()
		//	case "pepperoni":
		//		pizza = newNYStylePepperoniPizza()
	}

	if pizza == nil {
		return nil, fmt.Errorf("Invalid pizza type")
	}

	return pizza, nil
}
