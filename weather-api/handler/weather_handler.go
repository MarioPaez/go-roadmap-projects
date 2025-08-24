package handler

import (
	"api-weather/cache"
	"api-weather/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

type WeatherService struct {
	RedisClient *redis.Client
	Ctx         context.Context
}

func NewWeatherService(redisClient *redis.Client, ctx context.Context) *WeatherService {
	return &WeatherService{
		RedisClient: redisClient,
		Ctx:         ctx,
	}
}

func (weatherService *WeatherService) GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "missing city", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Weather requested for city: %s", city)
	weatherService.getWeather(city)
}

func (weatherService *WeatherService) getWeather(city string) {
	url := buildUrl(city)
	var result model.Information
	if cache.CheckIfCityExist(city, weatherService.RedisClient, weatherService.Ctx) {
		val := cache.GetCity(city, weatherService.RedisClient, weatherService.Ctx)
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			log.Fatal("could not possible to unmarshal the response into Information struct")
		}
		fmt.Println("the city has been obtained by the cache")
	} else {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			log.Fatal("city does not exist")
		}

		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Fatal("could not possible to decoder the response into Information struct", err)
		}
		cache.SaveCity(city, result, weatherService.RedisClient, weatherService.Ctx)
		fmt.Println("the city has been obtained by the 3rd service")
	}

	printInformation(result)
}

func buildUrl(city string) string {
	apiKey := os.Getenv(model.KEY)
	url := strings.Replace(model.URL, model.CITY_PLACE_HOLDER, city, 1)
	return strings.Replace(url, model.API_KEY_PLACE_HOLDER, apiKey, 1)
}

func printInformation(result model.Information) {
	pretty, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pretty))
}
