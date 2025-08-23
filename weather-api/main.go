package main

import (
	"api-weather/redis"
	"api-weather/weather"
	"context"
	"log"
)

func main() {

	redisClient := redis.Setup()
	ctx := context.Background()

	if err := weather.WeatherService(redisClient, ctx); err != nil {
		log.Fatal(err)
	}

}
