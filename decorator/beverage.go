package main
/*
Use interface to implement abstract class
 */

//Abstract Interface
type beverage interface {
	description() string
	cost() float64
}
// Abstract Concrete Type
type espresso struct {}

func (e *espresso) description() string{
	return "Espresso"
}

func (e *espresso) cost() float64 {
	return 1.99
}

type houseBlend struct {}

func (h *houseBlend) description() string {
	return "HouseBlend coffee"
}

func (h *houseBlend) cost() float64{
	return 0.89
}

type darkRoast struct{}

func (dr *darkRoast) description() string {
	return "Dark Roast Coffee"
}

func (dr *darkRoast) cost() float64 {
	return 0.99
}

type decaf struct{}

func (d *decaf) description() string {
	return "Decaf"
}

func (d *decaf) cost() float64 {
	return 1.05
}
