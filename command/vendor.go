// Receiver
package main

import "fmt"

type light struct {
	roomName string
}

func (l *light) on()  {
	fmt.Printf("%s light is turned on\n", l.roomName)
}

func (l *light) off()  {
	fmt.Printf("%s light is turned off\n", l.roomName)
}

type garage struct {}

func (g *garage) up()  {
	fmt.Printf("The garage is open\n")
}
