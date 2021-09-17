package main

import "fmt"

type person struct {
	name string
	gender string
	age int
}

func main() {
	err := addPerson("Tina", "Female", 20)
	if err != nil {
		fmt.Println("PersonAdd Error:" + err.Error())
	}

}
func addPerson(args ...interface{})  error {
	if len(args) > 3 {
		return fmt.Errorf("Wrong number of argumens passed")
	}

	p := &person{}

	for i, arg := range args {
		switch i {
		case 0:
			name, ok := arg.(string)
			if !ok {
				return fmt.Errorf("Name is not passed as string")
			}
			p.name = name
		case 1:
			gender, ok := arg.(string)
			if !ok {
				return fmt.Errorf("Gender is not passed as string")
			}
			p.gender = gender
		case 2:
			age, ok := arg.(int)
			if !ok {
				return fmt.Errorf("Age is not passed as string")
			}
			p.age = age
		}
	}
	fmt.Printf("Person struct is %+v\n", p)
	return nil
}
