//RunTime Polymorphism means that a call is resolved at runtime.
////It is achieved in GO using interfaces.
package main

import "fmt"

type taxSystem interface {
	calculateTax() int
}

type indianTax struct {
	taxPercentage int
	income int
}

func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income int
}

func (i *singaporeTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type usaTax struct {
	taxPercentage int
	income int
}

func (i *usaTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

func mains() {
	indianTax := &indianTax{
		taxPercentage: 30,
		income: 1000,
	}
	singaporeTax := &singaporeTax{
		taxPercentage: 10,
		income: 2000,
	}
	usaTax := &usaTax{
		taxPercentage: 40,
		income:        500,
	}

	taxSystems := []taxSystem{indianTax, singaporeTax, usaTax}
	totalTax := calculateTotalTax(taxSystems) //This is where runtime polymorphism happens
	fmt.Printf("Total Tax is %d\n", totalTax)
}

func calculateTotalTax(taxSystems []taxSystem) int {
	totalTax := 0
	for _, t := range taxSystems {
		totalTax += t.calculateTax()
	}
	return totalTax
}
