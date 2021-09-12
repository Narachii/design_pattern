package main

import "fmt"

func main()  {
	fmt.Println("check New Model Duck")
	modelDuck := newModelDuck()
	modelDuck.display()
	modelDuck.performFly()
	modelDuck.performQuack()
	modelDuck.swim()

	fmt.Println("Check Mallar Duck")
	modelDuck.setFlyingBehaviour(&flyRocketPowered{})
	modelDuck.performFly()

	mallarDuck := newMallarDuck()
	mallarDuck.display()
	mallarDuck.performFly()
	mallarDuck.performQuack()
	mallarDuck.swim()

	mallarDuck.setQuakingBehaviour(&squeak{})
	mallarDuck.performQuack()
}
