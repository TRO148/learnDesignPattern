package main

import "fmt"

type ForecastDisplay struct {
	headindex float32
}

func (fd *ForecastDisplay) update(weatherData *WeatherData) {
	fd.headindex = 16.923 + (0.185212 * weatherData.temp) + (5.37941 * weatherData.humidity) - (0.100254 * weatherData.temp * weatherData.humidity) + (0.00941695 * (weatherData.temp * weatherData.temp)) + (0.00728898 * (weatherData.humidity * weatherData.humidity)) + (0.000345372 * (weatherData.temp * weatherData.temp * weatherData.humidity)) - (0.000814971 * (weatherData.temp * weatherData.humidity * weatherData.humidity)) + (0.0000102102 * (weatherData.temp * weatherData.temp * weatherData.humidity * weatherData.humidity)) - (0.000038646 * (weatherData.temp * weatherData.temp * weatherData.temp)) + (0.0000291583 * (weatherData.humidity * weatherData.humidity * weatherData.humidity)) + (0.00000142721 * (weatherData.temp * weatherData.temp * weatherData.temp * weatherData.humidity)) + (0.000000197483 * (weatherData.temp * weatherData.humidity * weatherData.humidity * weatherData.humidity)) - (0.0000000218429 * (weatherData.temp * weatherData.temp * weatherData.temp * weatherData.humidity * weatherData.humidity)) + 0.000000000843296*(weatherData.temp*weatherData.temp*weatherData.temp*weatherData.humidity*weatherData.humidity*weatherData.humidity) - (0.0000000000481975 * (weatherData.temp * weatherData.temp * weatherData.temp * weatherData.humidity * weatherData.humidity * weatherData.humidity))
	fd.display()
}

func (fd *ForecastDisplay) display() {
	fmt.Printf("Forecast: %0.2f\n", fd.headindex)
}
