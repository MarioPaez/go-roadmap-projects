package length

import (
	"fmt"
	"unit-tracker-gin/units"
)

func ValidateUnits(to, from string) error {
	units := units.LENGTH_UNITS_ALLOW

	if _, ok := units[to]; !ok {
		return fmt.Errorf("the unit '%s' is not supported. only support %s", to, units)
	}
	if _, exist := units[from]; !exist {
		return fmt.Errorf("the unit '%s' is not supported. only support %s", from, units)
	}

	return nil
}
