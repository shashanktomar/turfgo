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
		So(BearingToAngle(360 + 360 + 270.124), ShouldAlmostEqual, 270.124, twelveDecimalPlaces)
		So(BearingToAngle(- 105), ShouldAlmostEqual, 255, twelveDecimalPlaces)
		So(BearingToAngle(- 200), ShouldAlmostEqual, 160, twelveDecimalPlaces)
		So(BearingToAngle(- 360 - 34.6), ShouldAlmostEqual, 325.4, twelveDecimalPlaces)
		So(BearingToAngle(- 395), ShouldAlmostEqual, 325, twelveDecimalPlaces)

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
