package main

import "fmt"

type windows struct {
	printer printer
}

func (w *windows) print()  {
	fmt.Println("Print request for windows")
}

func (w *windows) setPrinter(p printer)  {
	w.printer = p
	fmt.Println("Set a New Printer for windows")
}
