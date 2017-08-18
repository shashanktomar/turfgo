package math

import "math"

func IsEqualFloat(first float64, second float64, epsilon float64) bool {
	return math.Abs(first-second) < epsilon
}

func IsEqualFloatPair(p1X float64, p1Y float64, p2X float64, p2Y float64, epsilon float64) bool {
	return math.Abs(p1X-p2X) < epsilon && math.Abs(p1Y-p2Y) < epsilon
}
