package turfgo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRadsToDegree(t *testing.T) {
	Convey("Should convert rads to degree", t, func() {
		So(RadsToDegree(1), ShouldAlmostEqual, 57.295, threeDecimalPlaces)
	})
}

func TestDegreeToRads(t *testing.T) {
	Convey("Should convert degree to rads", t, func() {
		So(DegreeToRads(1), ShouldAlmostEqual, 0.017, threeDecimalPlaces)
	})
}

func TestDistanceToRads(t *testing.T) {
	Convey("Should covert distance to rads", t, func() {
		d := DistanceToRads(1, Miles)
		So(d, ShouldAlmostEqual, 0.0002525252525252525, twelveDecimalPlaces)
	})
}

func TestRadsToDistance(t *testing.T) {
	Convey("Should covert rads to distance", t, func() {
		d := RadsToDistance(1, Miles)
		So(d, ShouldAlmostEqual, 3960, twelveDecimalPlaces)
	})
}

func TestDistanceToDegrees(t *testing.T) {
	Convey("Should covert distance to degrees", t, func() {
		d := DistanceToDegrees(1, Miles)
		So(d, ShouldAlmostEqual, 0.014468631190172304, twelveDecimalPlaces)
	})
}

func TestBearingToAngle(t *testing.T) {
	Convey("Should return angle between 0-360", t, func() {
		So(BearingToAngle(40), ShouldAlmostEqual, 40, twelveDecimalPlaces)
		So(BearingToAngle(410), ShouldAlmostEqual, 50, twelveDecimalPlaces)
		So(BearingToAngle(360+360+270.124), ShouldAlmostEqual, 270.124, twelveDecimalPlaces)
		So(BearingToAngle(-105), ShouldAlmostEqual, 255, twelveDecimalPlaces)
		So(BearingToAngle(-200), ShouldAlmostEqual, 160, twelveDecimalPlaces)
		So(BearingToAngle(-360-34.6), ShouldAlmostEqual, 325.4, twelveDecimalPlaces)
		So(BearingToAngle(-395), ShouldAlmostEqual, 325, twelveDecimalPlaces)

	})
}

func TestConvertDistance(t *testing.T) {
	Convey("Should convert distance between units", t, func() {
		So(ConvertDistance(1000, Meters, Kilometers), ShouldAlmostEqual, 1, twelveDecimalPlaces)
		So(ConvertDistance(1, Kilometers, Miles), ShouldAlmostEqual, 0.6213714106386318, twelveDecimalPlaces)
		So(ConvertDistance(1, Miles, Kilometers), ShouldAlmostEqual, 1.6093434343434343, twelveDecimalPlaces)
		So(ConvertDistance(1, NauticalMiles, Kilometers), ShouldAlmostEqual, 1.851999843075488, twelveDecimalPlaces)
		So(ConvertDistance(1, Meters, Centimeters), ShouldAlmostEqual, 100, twelveDecimalPlaces)
	})

}

func TestConvertArea(t *testing.T) {
	Convey("Should convert distance between units", t, func() {
		a, err := ConvertArea(0, Meters, Kilometers)
		So(err, ShouldBeNil)
		So(a, ShouldAlmostEqual, 0, twelveDecimalPlaces)

		a, err = ConvertArea(1000, Meters, Kilometers)
		So(err, ShouldBeNil)
		So(a, ShouldAlmostEqual, 0.001, twelveDecimalPlaces)

		a, err = ConvertArea(1, Kilometers, Miles)
		So(err, ShouldBeNil)
		So(a, ShouldAlmostEqual, 0.386, twelveDecimalPlaces)

		a, err = ConvertArea(1, Miles, Kilometers)
		So(err, ShouldBeNil)
		So(a, ShouldAlmostEqual, 2.5906735751295336, twelveDecimalPlaces)

		a, err = ConvertArea(1, Meters, Centimeters)
		So(err, ShouldBeNil)
		So(a, ShouldAlmostEqual, 10000, twelveDecimalPlaces)

		a, err = ConvertArea(100, Meters, Yards)
		So(err, ShouldBeNil)
		So(a, ShouldAlmostEqual, 119.59900459999999, twelveDecimalPlaces)

		a, err = ConvertArea(100, Meters, Feet)
		So(err, ShouldBeNil)
		So(a, ShouldAlmostEqual, 1076.3910417, twelveDecimalPlaces)

		a, err = ConvertArea(100000, Feet, Kilometers)
		So(err, ShouldBeNil)
		So(a, ShouldAlmostEqual, 0.009290303999749462, twelveDecimalPlaces)
	})

	Convey("Should return error if area is negative", t, func() {
		_, err := ConvertArea(-100, Meters, Kilometers)
		So(err.Error(), ShouldEqual, "area can't be negative")
	})

	Convey("Should return error if original unit is wrong", t, func() {
		_, err := ConvertArea(100, Radians, Kilometers)
		So(err.Error(), ShouldEqual, "invalid unit")
	})

	Convey("Should return error if final unit is wrong", t, func() {
		_, err := ConvertArea(100, Meters, Radians)
		So(err.Error(), ShouldEqual, "invalid unit")
	})

}
