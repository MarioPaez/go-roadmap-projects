package main

import (
	"api-weather/cache"
	"api-weather/server"
	"context"
)

func main() {

	redisClient := cache.Setup()
	ctx := context.Background()
	server.InitServer(redisClient, ctx)
}
