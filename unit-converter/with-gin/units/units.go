package units

var LENGTH_UNITS_ALLOW = map[string]struct{}{"millimeter": {}, "centimeter": {}, "meter": {}, "kilometer": {}, "inch": {}, "foot": {},
	"yard": {}, "mile": {}}
var WEIGHT_UNITS_ALLOW = map[string]struct{}{"milligram": {}, "gram": {}, "kilogram": {}, "ounce": {}, "pound": {}}
var TEMPERATURE_UNITS_ALLOW = map[string]struct{}{"Celsius": {}, "Fahrenheit": {}, "Kelvin": {}}

// var Relation = map[string]map[string]func(){
// 	"Celsius": {
// 		"Kevin":      celsiusToKevin,
// 		"Fahrenheit": celsiusToFahrenheit,
// 	},
// } Hacer mapa de mapa de funciones
