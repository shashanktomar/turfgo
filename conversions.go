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
func DistanceToRads(distance float64, unit Unit) float64 {
	return distance / radius[unit]
}

// RadsToDistance convert a distance measurement (assuming a spherical Earth) from radians to a more friendly unit.
func RadsToDistance(radians float64, unit Unit) float64 {
	return radians * radius[unit]
}

// DistanceToDegrees convert a distance measurement (assuming a spherical Earth) from a real-world unit into degrees
func DistanceToDegrees(distance float64, unit Unit) float64 {
	return RadsToDegree(DistanceToRads(distance, unit))
}

// ConvertDistance converts a distance to the requested unit.
func ConvertDistance(distance float64, originalUnit Unit, finalUnit Unit) float64 {
	return RadsToDistance(DistanceToRads(distance, originalUnit), finalUnit)
}

// BearingToAngle converts any bearing angle from the north line direction (positive clockwise)
// and returns an angle between 0-360 degrees (positive clockwise), 0 being the north line
func BearingToAngle(bearing float64) float64 {
	angle := math.Mod(bearing, 360)
	if angle < 0 {
		angle += 360
	}
	return angle
}
