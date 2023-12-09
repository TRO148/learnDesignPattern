package main

type WeatherData struct {
	observers   []Observer
	temperature float32
	humidity    float32
	pressure    float32
}

func (wd *WeatherData) register(o Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) deregister(o Observer) {
	wd.observers = filter(wd.observers, func(observer Observer) bool {
		return observer == o
	})
}

func (wd *WeatherData) notify() {
	for _, w := range wd.observers {
		w.update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) setData(temperature float32, humidity float32, pressure float32) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.notify()
}

func filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}
