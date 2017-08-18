package turfgo

import (
	"errors"
	"fmt"
	"math"

	tm "github.com/shashanktomar/turfgo/math"
)

// Along takes a line and returns a point at a specified distance along the line.
// Units should be one of km, mi, r or d
func Along(lineString *LineString, distance float64, unit string) (*Point, error) {
	travelled := float64(0)
	points := lineString.getPoints()
	for i, point := range points {
		if distance >= travelled && i == len(points)-1 {
			return point, nil
		} else if travelled >= distance {
			overshot := distance - travelled
			if overshot == 0 {
				return point, nil
			}
			bearing, err := Bearing(points[i], points[i-1])
			if err != nil {
				return nil, err
			}
			direction := bearing - 180
			interpolated, err := Destination(points[i], overshot, direction, unit)
			if err != nil {
				return nil, err
			}
			return interpolated, nil

		} else {
			t, err := Distance(points[i], points[i+1], unit)
			if err != nil {
				return nil, err
			}
			travelled += t
		}
	}
	return nil, nil
}

// Bearing takes two points and finds the geographic bearing between them.
func Bearing(point1, point2 *Point) (float64, error) {
	if point1 == nil || point2 == nil {
		return -1, errors.New("points can't be nil")
	}

	lat1, lng1 := tm.DegreesToRads(point1.Lat, point1.Lng)
	lat2, lng2 := tm.DegreesToRads(point2.Lat, point2.Lng)
	a := math.Sin(lng2-lng1) * math.Cos(lat2)
	b := math.Cos(lat1)*math.Sin(lat2) - math.Sin(lat1)*math.Cos(lat2)*math.Cos(lng2-lng1)
	return tm.RadToDegree(math.Atan2(a, b)), nil
}

// Center takes an array of points and returns the absolute center point of all points.
func Center(shapes ...Geometry) *Point {
	bBox := Extent(shapes...)
	lng := (bBox[0] + bBox[2]) / 2
	lat := (bBox[1] + bBox[3]) / 2
	return &Point{lat, lng}
}

// Destination takes a Point and calculates the location of a destination point
// given a distance in degrees, radians, miles, or kilometers; and bearing in
// degrees. This uses the Haversine formula to account for global curvature.
// Units should be one of km, mi, r or d
func Destination(start *Point, distance float64, bearing float64, unit string) (*Point, error) {
	if start == nil {
		return nil, errors.New("startPoint can't be nil")
	}

	radius, ok := R[unit]
	if !ok {
		return nil, invalidUnitError(unit)
	}

	lat, lon := tm.DegreesToRads(start.Lat, start.Lng)
	bearingRad := tm.DegreeToRad(bearing)

	destLat := math.Asin(math.Sin(lat)*math.Cos(distance/radius) +
		math.Cos(lat)*math.Sin(distance/radius)*math.Cos(bearingRad))
	destLon := lon + math.Atan2(math.Sin(bearingRad)*math.Sin(distance/radius)*math.Cos(lat),
		math.Cos(distance/radius)-math.Sin(lat)*math.Sin(destLat))

	return &Point{tm.RadToDegree(destLat), tm.RadToDegree(destLon)}, nil
}

// Distance calculates the distance between two points in degress, radians, miles, or
// kilometers. This uses the Haversine formula to account for global curvature.
func Distance(point1 *Point, point2 *Point, unit string) (float64, error) {
	radius, ok := R[unit]
	if !ok {
		return 0, fmt.Errorf(unitError, unit)
	}

	dLat := tm.DegreeToRad(point2.Lat - point1.Lat)
	dLon := tm.DegreeToRad(point2.Lng - point1.Lng)
	latRad1 := tm.DegreeToRad(point1.Lat)
	latRad2 := tm.DegreeToRad(point2.Lat)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(latRad1)*math.Cos(latRad2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return radius * c, nil
}

// Extent Takes a set of features, calculates the extent of all input features, and returns a bounding box.
// Returns []float64 the bounding box of input given as an array in WSEN order (west, south, east, north)
func Extent(shapes ...Geometry) []float64 {
	extent := []float64{infinity, infinity, -infinity, -infinity}

	for _, shape := range shapes {
		for _, point := range shape.getPoints() {
			if extent[0] > point.Lng {
				extent[0] = point.Lng
			}
			if extent[1] > point.Lat {
				extent[1] = point.Lat
			}
			if extent[2] < point.Lng {
				extent[2] = point.Lng
			}
			if extent[3] < point.Lat {
				extent[3] = point.Lat
			}
		}
	}
	return extent
}

// Overlap takes two bounding box and returns true if there is an overlap.
// The order of values in array is WSEN(west, south , east, north)
func Overlap(b1 []float64, b2 []float64) (bool, error) {
	if len(b1) != 4 || len(b2) != 4 {
		return false, errors.New("Invalid bbox")
	}
	w1, s1, e1, n1 := b1[0], b1[1], b1[2], b1[3]
	w2, s2, e2, n2 := b2[0], b2[1], b2[2], b2[3]

	// b2 is left of b1
	if w1 > e2 {
		return false, nil
	}
	// b2 is right of b1
	if e1 < w2 {
		return false, nil
	}
	// b2 is above b1
	if n1 < s2 {
		return false, nil
	}
	// b2 is below b1
	if s1 > n2 {
		return false, nil
	}

	return true, nil
}

// Surround Takes a point and a width, calculates the bounding box around the point with the given width.
// Returns []float64 the bounding box of input given as an array in WSEN order (west, south, east, north)
func Surround(point *Point, width float64) []float64 {
	bottomLeft := translate(point, -width, -width)
	topRight := translate(point, width, width)

	bbox := []float64{bottomLeft.Lng, bottomLeft.Lat, topRight.Lng, topRight.Lat}

	return bbox
}
