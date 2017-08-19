package turfgo

import "math"

// RadsToDegree convert a radians (assuming a spherical Earth) into degrees
func RadsToDegree(rad float64) float64 {
	return rad * 180 / math.Pi
}

// DegreeToRads convert degrees (assuming a spherical Earth) into radians
func DegreeToRads(degree float64) float64 {
	return degree * math.Pi / 180
}

// DegreesToRads convert a pair of floats (assuming a spherical Earth) into a pair of radians
func DegreesToRads(first float64, second float64) (float64, float64) {
	return DegreeToRads(first), DegreeToRads(second)
}

// DistanceToRads convert a distance measurement (assuming a spherical Earth) from a real-world unit into radians
// Valid units: miles, nauticalmiles, inches, yards, meters, metres, kilometers, centimeters, feet
func DistanceToRads(distance float64, unit string) (float64, error) {
	f, ok := Factors[unit]
	if !ok {
		return -1, invalidUnitError(unit)
	}
	return distance / f, nil
}

// RadsToDistance convert a distance measurement (assuming a spherical Earth) from radians to a more friendly unit.
// Valid units: miles, nauticalmiles, inches, yards, meters, kilometers, centimeters, feet
func RadsToDistance(radians float64, unit string) (float64, error) {
	f, ok := Factors[unit]
	if !ok {
		return -1, invalidUnitError(unit)
	}
	return radians * f, nil
}

// DistanceToDegrees convert a distance measurement (assuming a spherical Earth) from a real-world unit into degrees
// Valid units: miles, nauticalmiles, inches, yards, meters, centimeters, kilometres, feet
func DistanceToDegrees(distance float64, unit string) (float64, error) {
	d, err := DistanceToRads(distance, unit)
	if err != nil {
		return -1, err
	}
	return RadsToDegree(d), nil
}
