package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	adidasShort := adidasFactory.makeShort()
	adidasShoe := adidasFactory.makeShoe()

	nikeFactory, _ := getSportsFactory("nike")
	nikeShort := nikeFactory.makeShort()
	nikeShoe := nikeFactory.makeShoe()
	printShortDetails(adidasShort)
	printShoeDetails(adidasShoe)

	printShortDetails(nikeShort)
	printShoeDetails(nikeShoe)

	adidasShoe.setLogo("new adidas Shoe")
	adidasShoe.setSize(24)
	printShoeDetails(adidasShoe)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShortDetails(s iShort) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}
