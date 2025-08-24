package server

import (
	"api-weather/handler"
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func InitServer(redisClient *redis.Client, ctx context.Context) {
	weatherHandler := handler.NewWeatherService(redisClient, ctx)
	http.HandleFunc("/weather", weatherHandler.GetWeatherHandler)

	http.ListenAndServe(":8080", nil)
}
