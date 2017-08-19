package turfgo

func isEqualLocation(point1 *Point, point2 *Point) bool {
	return IsEqualFloatPair(point1.Lat, point1.Lng, point2.Lat, point2.Lng, TwelveDecimalPlaces)
}
