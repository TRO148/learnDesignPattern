package main

import "fmt"

type StatisticsDisplay struct {
	max float32
	min float32
}

func (sd *StatisticsDisplay) update(weatherData *WeatherData) {
	temp := weatherData.temp
	humidity := weatherData.humidity
	pressure := weatherData.pressure
	sd.max = max(temp, max(humidity, pressure))
	sd.min = min(temp, min(humidity, pressure))
	sd.display()
}

func (sd *StatisticsDisplay) display() {
	fmt.Printf("Avg/Max/Min: %0.2f/%0.2f/%0.2f\n", (sd.max+sd.min)/2, sd.max, sd.min)
}
