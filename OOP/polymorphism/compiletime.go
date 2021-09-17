package main

import "fmt"

type maths struct {}

func (m *maths) add(numbers ...int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func main() {
	m := &maths{}

	fmt.Printf("Result: %d\n", m.add(2, 3))
	fmt.Printf("Result: %d\n", m.add(2, 3, 4))
}
// Go doesn't support Method Overloading. below program will raise a compilation error
/*
package main

type maths struct{}

func (m *maths) add(a, b int) int {
	return a + b
}

func (m *maths) add(a, b, c int) int {
	return a + b + c
}

func main() {
	m := &maths{}
}

Output >>
(*maths).add redeclared in this block
        previous declaration at
*/
