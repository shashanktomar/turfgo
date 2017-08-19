package turfgo

import (
  . "github.com/smartystreets/goconvey/convey"
  "testing"
)

func TestRadsToDegree(t *testing.T) {
  Convey("Should convert rads to degree", t, func() {
    So(RadsToDegree(1), ShouldAlmostEqual, 57.295, ThreeDecimalPlaces)
  })
}

func TestDegreeToRads(t *testing.T) {
  Convey("Should convert degree to rads", t, func() {
    So(DegreeToRads(1), ShouldAlmostEqual, 0.017, ThreeDecimalPlaces)
  })
}

func TestDistanceToRads(t *testing.T){
  Convey("Should covert distance to rads", t, func() {
    d, err := DistanceToRads(1, Miles)
    So(err, ShouldBeNil)
    So(d, ShouldAlmostEqual, 0.0002525252525252525, TwelveDecimalPlaces)
  })

  Convey("Should throw error if unit is invalid", t, func() {
    _, err := DistanceToRads(1, "invalidUnit")
    So(err.Error(), ShouldEqual, invalidUnitError("invalidUnit").Error())
  })
}

func TestRadsToDistance(t *testing.T){
  Convey("Should covert rads to distance", t, func() {
    d, err := RadsToDistance(1, Miles)
    So(err, ShouldBeNil)
    So(d, ShouldAlmostEqual, 3960, TwelveDecimalPlaces)
  })

  Convey("Should throw error if unit is invalid", t, func() {
    _, err := RadsToDistance(1, "invalidUnit")
    So(err.Error(), ShouldEqual, invalidUnitError("invalidUnit").Error())
  })
}

func TestDistanceToDegrees(t *testing.T){
  Convey("Should covert distance to degrees", t, func() {
    d, err := DistanceToDegrees(1, Miles)
    So(err, ShouldBeNil)
    So(d, ShouldAlmostEqual, 0.014468631190172304, TwelveDecimalPlaces)
  })

  Convey("Should throw error if unit is invalid", t, func() {
    _, err := DistanceToDegrees(1, "invalidUnit")
    So(err.Error(), ShouldEqual, invalidUnitError("invalidUnit").Error())
  })
}
