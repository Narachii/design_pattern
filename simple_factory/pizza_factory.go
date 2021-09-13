package main

import "fmt"

type pizzaFactory interface {
	createPizza(pizzaType string) (iPizza, error)
}

type simplePizzaFactory struct{}

func (sp *simplePizzaFactory) createPizza(pizzaType string) (iPizza, error) {
	var pizza iPizza
	switch pizzaType {
	case "cheese":
		pizza = newCheesePizza()
	case "greek":
		pizza = newGreekPizza()
	case "pepperoni":
		pizza = newPepperoniPizza()
	}

	if pizza == nil {
		return nil, fmt.Errorf("Invalid pizza type")
	}

	return pizza, nil
}
