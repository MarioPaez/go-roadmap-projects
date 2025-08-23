package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func WeatherService() error {
	url := buildUrl()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	printInformation(resp)

	return nil
}

func buildUrl() string {
	apiKey := os.Getenv(KEY)
	url := strings.Replace(URL, CITY_PLACE_HOLDER, "Milano", 1)
	return strings.Replace(url, API_KEY_PLACE_HOLDER, apiKey, 1)
}

func printInformation(resp *http.Response) {
	var result Information

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}
	pretty, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pretty))
}
