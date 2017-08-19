package turfgo

import "math"

const(
	ThreeDecimalPlaces float64 = .001
	TwelveDecimalPlaces float64 = .000000000001
)

func RadsToDegree(rad float64)  float64{
	return rad * 180 / math.Pi
}

func DegreeToRads(degree float64)  float64{
	return degree * math.Pi / 180
}

func DegreesToRads(first float64, second float64)  (float64, float64){
	return DegreeToRads(first), DegreeToRads(second)
}
