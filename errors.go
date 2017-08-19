package turfgo

import "fmt"

var unitError = "%s is not a valid unit. Allowed units are mi(miles), km(kilometers), d(degrees) and r(radians)"

func invalidUnitError(unit string) error {
	return fmt.Errorf(unitError, unit)
}
