package units

var ToMeters = map[string]func(float64) float64{
	"millimeter": func(v float64) float64 { return v / 1000 },
	"centimeter": func(v float64) float64 { return v / 100 },
	"meter":      func(v float64) float64 { return v },
	"kilometer":  func(v float64) float64 { return v * 1000 },
	"inch":       func(v float64) float64 { return v * 0.0254 },
	"foot":       func(v float64) float64 { return v * 0.3048 },
	"yard":       func(v float64) float64 { return v * 0.9144 },
	"mile":       func(v float64) float64 { return v * 1609.344 },
}

var FromMeters = map[string]func(float64) float64{
	"millimeter": func(v float64) float64 { return v * 1000 },
	"centimeter": func(v float64) float64 { return v * 100 },
	"meter":      func(v float64) float64 { return v },
	"kilometer":  func(v float64) float64 { return v / 1000 },
	"inch":       func(v float64) float64 { return v / 0.0254 },
	"foot":       func(v float64) float64 { return v / 0.3048 },
	"yard":       func(v float64) float64 { return v / 0.9144 },
	"mile":       func(v float64) float64 { return v / 1609.344 },
}

var ToKilogram = map[string]func(float64) float64{
	"milligram": func(v float64) float64 { return v / 1e6 },
	"gram":      func(v float64) float64 { return v / 1000 },
	"kilogram":  func(v float64) float64 { return v },
	"ounce":     func(v float64) float64 { return v * 0.0283495 },
	"pound":     func(v float64) float64 { return v * 0.453592 },
}

var FromKilogram = map[string]func(float64) float64{
	"milligram": func(v float64) float64 { return v * 1e6 },
	"gram":      func(v float64) float64 { return v * 1000 },
	"kilogram":  func(v float64) float64 { return v },
	"ounce":     func(v float64) float64 { return v / 0.0283495 },
	"pound":     func(v float64) float64 { return v / 0.453592 },
}

var ToCelsius = map[string]func(float64) float64{
	"Celsius":    func(v float64) float64 { return v },
	"Fahrenheit": func(v float64) float64 { return (v - 32) * 5 / 9 },
	"Kelvin":     func(v float64) float64 { return v - 273.15 },
}

var FromCelsius = map[string]func(float64) float64{
	"Celsius":    func(v float64) float64 { return v },
	"Fahrenheit": func(v float64) float64 { return v*9/5 + 32 },
	"Kelvin":     func(v float64) float64 { return v + 273.15 },
}
