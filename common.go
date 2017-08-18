package turfgo

import (
	"math"
	tm "github.com/shashanktomar/turfgo/math"
)

func isEqualLocation(point1 *Point, point2 *Point) bool {
	return tm.IsEqualFloatPair(point1.Lat, point1.Lng, point2.Lat, point2.Lng, tm.TwelveDecimalPlaces)
}

func translate(point *Point, horizontalDisplacement float64, verticalDisplacement float64) *Point {
	latDisplacementRad := verticalDisplacement / R["m"]
	longDisplacementRad := horizontalDisplacement / (R["m"] * math.Cos(tm.DegreeToRad(point.Lat)))

	latDisplacement := tm.RadToDegree(latDisplacementRad)
	longDisplacement := tm.RadToDegree(longDisplacementRad)

	translatedPoint := NewPoint(point.Lat+latDisplacement, point.Lng+longDisplacement)

	return translatedPoint
}
