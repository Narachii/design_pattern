package main

import "fmt"

func main() {
	beverage := &espresso{}
	fmt.Printf("%s $%.2f\n", beverage.description(), beverage.cost())

	// Make a DarkRoast object.
	darkRoast := &darkRoast{}

	// Wrap it with a Moch
	singleMocha := &mocha{
		beverage: darkRoast,
	}

	// Wrap it in a cecond Mocha
	doubleMocha := &mocha{
		beverage: singleMocha,
	}

	// Wrap it in a Whip
	doubleMochaWhip := &whip{
		beverage: doubleMocha,
	}
	fmt.Printf("%s $%.2f\n", doubleMochaWhip.description(), doubleMocha.cost())

	// Finally give us a HouseBlend with Soy, Mocha, and Whip.
	soyMochaWhipHouseBlend := &whip{
		&mocha{
			&soy{
				&houseBlend{},
			},
		},
	}

	fmt.Printf("%s $%.2f\n", soyMochaWhipHouseBlend.description(), soyMochaWhipHouseBlend.cost())
}
