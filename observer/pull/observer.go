package main

type Observer interface {
	update(weatherData *WeatherData)
}
