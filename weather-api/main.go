package main

import (
	"fmt"
	"os"
)

const KEY = "API_WEATHER_KEY"

func main() {
	apiKey := os.Getenv(KEY)
	fmt.Println(apiKey)
}
