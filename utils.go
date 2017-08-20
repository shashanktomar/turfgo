package turfgo

import (
	"math"
)

const (
	threeDecimalPlaces  float64 = .001
	twelveDecimalPlaces float64 = .000000000001
)

func isEqualFloat(first float64, second float64, epsilon float64) bool {
	return math.Abs(first-second) < epsilon
}

func isEqualFloatPair(p1X float64, p1Y float64, p2X float64, p2Y float64, epsilon float64) bool {
	return math.Abs(p1X-p2X) < epsilon && math.Abs(p1Y-p2Y) < epsilon
}

func isEqualLocation(point1 *Point, point2 *Point) bool {
	return isEqualFloatPair(point1.Lat, point1.Lng, point2.Lat, point2.Lng, twelveDecimalPlaces)
}
