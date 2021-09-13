package main

import "fmt"

func main() {
	nyStore := newNYPizzaStore()
	pizza := nyStore.orderPizza("cheese")
	fmt.Printf("Ethan ordered %s pizz\n\n", pizza.getName())

	chicagoStore := newChicagoPizzaStore()
	cPizza := chicagoStore.orderPizza("cheese")
	fmt.Printf("Joel ordered %s pizz\n\n", cPizza.getName())
}
