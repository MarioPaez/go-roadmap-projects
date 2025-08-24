package cache

import (
	"api-weather/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func CheckIfCityExist(city string, client *redis.Client, ctx context.Context) bool {
	exists, err := client.Exists(ctx, city).Result()
	if err != nil {
		log.Fatal("error trying to check if the key exist in redis cache")
	}
	fmt.Println("valor de exists", exists)
	return exists > 0
}

func SaveCity(city string, result model.Information, client *redis.Client, ctx context.Context) {
	resultBytes, err := json.Marshal(result)
	if err != nil {
		log.Fatal("could not possible to do the marshal to save the city", err)
	}
	if _, err := client.Set(ctx, city, resultBytes, time.Hour*12).Result(); err != nil {
		log.Fatal("could not possible to save the object into redis cache", err)
	}
	fmt.Println("weather information saved successfully into redis cache")
}

func GetCity(city string, client *redis.Client, ctx context.Context) string {
	value, err := client.Get(ctx, city).Result()
	if err != nil {
		log.Fatal("error trying to retrieve if the information in redis cache")
	}
	return value
}
