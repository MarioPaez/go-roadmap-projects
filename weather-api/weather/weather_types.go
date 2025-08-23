package weather

const (
	KEY                  = "API_WEATHER_KEY"
	URL                  = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/{city}/today?key={API_KEY}"
	CITY_PLACE_HOLDER    = "{city}"
	API_KEY_PLACE_HOLDER = "{API_KEY}"
)

type Information struct {
	Address           string            `json:"address"`
	CurrentConditions CurrentConditions `json:"currentConditions"`
	Description       string            `json:"description"`
	Latitude          float32           `json:"latitude"`
	Longitude         float32           `json:"longitude"`
	ResolveAddress    string            `json:"resolvedAddress"`
	Timezone          string            `json:"timezone"`
}

type CurrentConditions struct {
	Conditions  string  `json:"conditions"`
	Datetime    string  `json:"datetime"`
	Humidity    float32 `json:"humidity"`
	Pressure    float32 `json:"pressure"`
	Temperature float32 `json:"temp"`
	Sunrise     string  `json:"sunrise"`
	Sunset      string  `json:"sunset"`
	WindSpeed   float32 `json:"windspeed"`
}
