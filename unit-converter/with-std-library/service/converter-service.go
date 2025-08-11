package service

import (
	"fmt"
	"std-library/model"
	"std-library/units"
)

type Converter struct {
	ToIntermediate   map[string]func(float64) float64
	FromIntermediate map[string]func(float64) float64
	SupportedUnits   []string
}

var Converters = map[string]Converter{
	"length": {
		ToIntermediate:   units.ToMeters,
		FromIntermediate: units.FromMeters,
		SupportedUnits:   []string{"millimeter", "centimeter", "meter", "kilometer", "inch", "foot", "yard", "mile"},
	},
	"weight": {
		ToIntermediate:   units.ToKilogram,
		FromIntermediate: units.FromKilogram,
		SupportedUnits:   []string{"milligram", "gram", "kilogram", "ounce", "pound"},
	},
	"temperature": {
		ToIntermediate:   units.ToCelsius,
		FromIntermediate: units.FromCelsius,
		SupportedUnits:   []string{"Celsius", "Fahrenheit", "Kelvin"},
	},
}

func DoConversion(c *model.Conversion) (float64, error) {
	converter, ok := Converters[string(c.Type)]
	if !ok {
		return -1, fmt.Errorf("unsupported conversion type: %s", c.Type)
	}

	toIntermediate, ok := converter.ToIntermediate[c.From]
	if !ok {
		return -1, fmt.Errorf("unit '%s' is not supported for type '%s'. Supported: %v", c.From, c.Type, converter.SupportedUnits)
	}

	fromIntermediate, ok := converter.FromIntermediate[c.To]
	if !ok {
		return -1, fmt.Errorf("unit '%s' is not supported for type '%s'. Supported: %v", c.To, c.Type, converter.SupportedUnits)
	}

	intermediateValue := toIntermediate(c.Value)
	return fromIntermediate(intermediateValue), nil
}
