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

func DegreesToRads(first float64, second float64)  (float64, float64){
  return DegreeToRad(first), DegreeToRad(second)
}
