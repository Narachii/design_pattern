package main

import "fmt"

type observer interface {
	update(temp float64, humidity float64, pressure float64)
	getID() int64
}

type currentConditionDisplay struct {
	id int64
	temperature float64
	humidity float64
	pressure float64
	weatherData subject
}

func newCurrentConditionDisplay(s subject, id int64) *currentConditionDisplay {
	d := &currentConditionDisplay{
		id: id,
		weatherData: s,
	}
	s.registerObserver(d)

	return d
}

func (cd *currentConditionDisplay) update(temperature float64, humidity float64, pressure float64)  {
	cd.temperature = temperature
	cd.humidity = humidity
	cd.pressure = pressure
	cd.display()
}

func (cd *currentConditionDisplay) getID()  int64 {
	return cd.id
}

func (cd currentConditionDisplay) display()  {
	fmt.Printf("Current conditions: Temperature:%.2f, Humidity:%.2f\n", cd.temperature, cd.humidity)
}
