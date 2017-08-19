package turfgo

import (
	"errors"
	"math"
)

// Along takes a line and returns a point at a specified distance along the line.
// Returns the last point if distance is more than the span of the line.
func Along(lineString *LineString, distance float64, unit Unit) (*Point, error) {
	if lineString == nil {
		return nil, errors.New("lineString can't be nil")
	}

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
			bearing, err := Bearing(point, points[i-1])
			if err != nil {
				return nil, err
			}
			direction := bearing - 180
			interpolated, err := Destination(point, overshot, direction, unit)
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

	lat1, lng1 := DegreesToRads(point1.Lat, point1.Lng)
	lat2, lng2 := DegreesToRads(point2.Lat, point2.Lng)
	a := math.Sin(lng2-lng1) * math.Cos(lat2)
	b := math.Cos(lat1)*math.Sin(lat2) - math.Sin(lat1)*math.Cos(lat2)*math.Cos(lng2-lng1)
	return RadsToDegree(math.Atan2(a, b)), nil
}

// Center takes an array of points and returns the absolute center point of all points.
func Center(shapes ...Geometry) *Point {
	bBox := Extent(shapes...)
	lng := (bBox.West + bBox.East) / 2
	lat := (bBox.South + bBox.North) / 2
	return NewPoint(lat, lng)
}

// Destination takes a Point and calculates the location of a destination point
// given a distance in degrees, radians, miles, or kilometers; and bearing in
// degrees. This uses the Haversine formula to account for global curvature.
func Destination(start *Point, distance float64, bearing float64, unit Unit) (*Point, error) {
	if start == nil {
		return nil, errors.New("startPoint can't be nil")
	}

	r := DistanceToRads(distance, unit)
	lat, lon := DegreesToRads(start.Lat, start.Lng)
	bearingRad := DegreeToRads(bearing)

	destLat := math.Asin(math.Sin(lat)*math.Cos(r) +
		math.Cos(lat)*math.Sin(r)*math.Cos(bearingRad))
	destLon := lon + math.Atan2(math.Sin(bearingRad)*math.Sin(r)*math.Cos(lat),
		math.Cos(r)-math.Sin(lat)*math.Sin(destLat))

	return &Point{RadsToDegree(destLat), RadsToDegree(destLon)}, nil
}

// Distance calculates the distance between two points in degress, radians, miles, or
// kilometers. This uses the Haversine formula to account for global curvature.
func Distance(point1 *Point, point2 *Point, unit Unit) (float64, error) {
	if point1 == nil || point2 == nil {
		return -1, errors.New("points can't be nil")
	}
	dLat, dLng := DegreesToRads(point2.Lat-point1.Lat, point2.Lng-point1.Lng)
	latRad1 := DegreeToRads(point1.Lat)
	latRad2 := DegreeToRads(point2.Lat)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLng/2)*math.Sin(dLng/2)*math.Cos(latRad1)*math.Cos(latRad2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return RadsToDistance(c, unit), nil
}

// Bbox is an alias for Extent
func Bbox(shapes ...Geometry) *BoundingBox {
	return Extent(shapes...)
}

// Extent Takes a set of features, calculates the extent of all input features, and returns a bounding box.
func Extent(geometries ...Geometry) *BoundingBox {
	extent := NewInfiniteBBox()
	for _, shape := range geometries {
		for _, point := range shape.getPoints() {
			if extent.West > point.Lng {
				extent.West = point.Lng
			}
			if extent.South > point.Lat {
				extent.South = point.Lat
			}
			if extent.East < point.Lng {
				extent.East = point.Lng
			}
			if extent.North < point.Lat {
				extent.North = point.Lat
			}
		}
	}
	return extent
}

// BboxToCorners return the corner points SouthWest and NorthEast from bbox
func BboxToCorners(box *BoundingBox) (*Point, *Point) {
	return NewPoint(box.South, box.West), NewPoint(box.North, box.East)
}

// Expand Takes a set of features, calculates a collective bounding box around the features
// and expand it by the given distance in all directions. It returns a bounding box.
func Expand(distance float64, unit Unit, geometries ...Geometry) (*BoundingBox, error) {
	bbox := Bbox(geometries...)
	bottomLeft, topRight := BboxToCorners(bbox)

	leftEdge, err := Destination(bottomLeft, distance, -90, unit)
	if err != nil {
		return nil, err
	}
	bottomEdge, err := Destination(bottomLeft, distance, 180, unit)
	if err != nil {
		return nil, err
	}
	rightEge, err := Destination(topRight, distance, 90, unit)
	if err != nil {
		return nil, err
	}
	topEdge, err := Destination(topRight, distance, 0, unit)
	if err != nil {
		return nil, err
	}
	return Extent(leftEdge, bottomEdge, rightEge, topEdge), nil
}
