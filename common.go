package turfgo

import (
	"math"
)

func isEqualLocation(point1 *Point, point2 *Point) bool {
	return IsEqualFloatPair(point1.Lat, point1.Lng, point2.Lat, point2.Lng, TwelveDecimalPlaces)
}

func translate(point *Point, horizontalDisplacement float64, verticalDisplacement float64) *Point {
	latDisplacementRad := verticalDisplacement / R[Meters]
	longDisplacementRad := horizontalDisplacement / (R[Meters] * math.Cos(DegreeToRads(point.Lat)))

	latDisplacement := RadsToDegree(latDisplacementRad)
	longDisplacement := RadsToDegree(longDisplacementRad)

	translatedPoint := NewPoint(point.Lat+latDisplacement, point.Lng+longDisplacement)

	return translatedPoint
}
