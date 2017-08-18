package math

import "math"

const(
  ThreeDecimalPlaces float64 = .001
  TwelveDecimalPlaces float64 = .000000000001
)

func RadToDegree(rad float64)  float64{
  return rad * 180 / math.Pi
}

func DegreeToRad(degree float64)  float64{
  return degree * math.Pi / 180
}

