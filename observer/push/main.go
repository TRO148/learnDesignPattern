package main

func main() {
	wd := &WeatherData{}
	ccd := &CurrentConditionDisplay{}
	fd := &ForecastDisplay{}
	sd := &StatisticsDisplay{}

	wd.register(ccd)
	wd.register(fd)
	wd.register(sd)
	wd.setData(80, 65, 30.4)
	wd.setData(82, 70, 29.2)
	wd.setData(78, 90, 29.2)
}
