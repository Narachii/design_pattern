package main

func main()  {
	weatherData := newWeatherData()
	newCurrentConditionDisplay(weatherData, 1)
	newCurrentConditionDisplay(weatherData, 2)
	weatherData.setMeasurements(27, 65, 30.4)
	weatherData.setMeasurements(28, 70, 29.2)
	weatherData.setMeasurements(26, 90, 29.2)

}
