package main

func main() {
	factory := &simplePizzaFactory{}
	pizzaStore := newPizzaStore(factory)
	pizzaStore.orderPizza("pepperoni")
}
