package turfgo

import (
  "testing"
)

func TestRadsToDegree(t *testing.T){
  var expected float64 = 57.295
  result := RadsToDegree(1)
  if !IsEqualFloat(result, expected, ThreeDecimalPlaces){
    t.Errorf("Expected: %g, Actual: %g", expected, result)
  }
}

func TestDegreeToRads(t *testing.T){
  var expected float64 = 0.017
  result := DegreeToRads(1)
  if !IsEqualFloat(result, expected, ThreeDecimalPlaces){
    t.Errorf("Expected: %g, Actual: %g", expected, result)
  }
}
