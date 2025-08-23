package main

import (
	"api-weather/weather"
)

func main() {
	if err := weather.WeatherService(); err != nil {
		panic(err)
	}
}
