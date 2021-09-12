package main

type subject interface {
	registerObserver(o observer)
	removerObserver(o  observer)
	notifyObservers()
}

type weatherData struct {
	observers []observer
	temperature float64
	humidity 	float64
	pressure 	float64
}

func newWeatherData() *weatherData {
	return &weatherData{}
}

func (w *weatherData) registerObserver(o observer)  {
	w.observers = append(w.observers, o)
}

func (w *weatherData) removerObserver(o  observer) {
	w.observers = removeFromslice(w.observers, o)
}

func (w *weatherData) notifyObservers()  {
	for _, observer := range w.observers {
		observer.update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *weatherData) measurementChanged()  {
	w.notifyObservers()
}

func (w *weatherData) setMeasurements(temperature float64, humidity float64, pressure float64)  {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure

	w.measurementChanged()
}


func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength -1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
