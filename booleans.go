package turfgo

import (
	"errors"
	"math"
)

// Boundary type
type Boundary int

// Boundary constants
const (
	BoundaryStart Boundary = iota
	BoundaryEnd
	BoundaryBoth
	BoundaryNone
)

// DoesBboxOverlap takes two bounding box and returns true if there is an overlap.
// The order of values in array is WSEN(west, south , east, north)
func DoesBboxOverlap(b1 *BoundingBox, b2 *BoundingBox) (bool, error) {
	if b1 == nil || b2 == nil {
		return false, errors.New("Bbox can't be nil")
	}

	// b2 is left of b1
	if b1.West > b2.East {
		return false, nil
	}
	// b2 is right of b1
	if b1.East < b2.West {
		return false, nil
	}
	// b2 is above b1
	if b1.North < b2.South {
		return false, nil
	}
	// b2 is below b1
	if b1.South > b2.North {
		return false, nil
	}

	return true, nil
}

// IsPointOnLine returns true if a point is on a line.
// Accepts a parameter to ignore the start and end vertices of the linestring.
func IsPointOnLine(point Point, lineString *LineString, ignoreEnds bool) bool {
	points := lineString.getPoints()
	for i := 0; i < len(points)-1; i++ {
		excludeBoundary := BoundaryNone
		if ignoreEnds {
			if i == 0 {
				excludeBoundary = BoundaryStart
			}
			if i == len(points)-2 {
				excludeBoundary = BoundaryEnd
			}
			if i == 0 && i+1 == len(points)-1 {
				excludeBoundary = BoundaryBoth
			}
		}
		if isPointOnLineSegment(*points[i], *points[i+1], point, excludeBoundary) {
			return true
		}
	}
	return false
}

func isPointOnLineSegment(start Point, end Point, point Point, excludeBoundary Boundary) bool {
	dxc := point.Lng - start.Lng
	dyc := point.Lat - start.Lat
	dxl := end.Lng - start.Lng
	dyl := end.Lat - start.Lat
	cross := dxc*dyl - dyc*dxl
	if cross != 0 {
		return false
	}

	switch {
	case excludeBoundary == BoundaryNone:
		if math.Abs(dxl) >= math.Abs(dyl) {
			if dxl > 0 {
				return start.Lng <= point.Lng && point.Lng <= end.Lng
			}
			return end.Lng <= point.Lng && point.Lng <= start.Lng
		}
		if dyl > 0 {
			return start.Lat <= point.Lat && point.Lat <= end.Lat
		}
		return end.Lat <= point.Lat && point.Lat <= start.Lat
	case excludeBoundary == BoundaryStart:
		if math.Abs(dxl) >= math.Abs(dyl) {
			if dxl > 0 {
				return start.Lng < point.Lng && point.Lng <= end.Lng
			}
			return end.Lng <= point.Lng && point.Lng < start.Lng
		}
		if dyl > 0 {
			return start.Lat < point.Lat && point.Lat <= end.Lat
		}
		return end.Lat <= point.Lat && point.Lat < start.Lat
	case excludeBoundary == BoundaryEnd:
		if math.Abs(dxl) >= math.Abs(dyl) {
			if dxl > 0 {
				return start.Lng <= point.Lng && point.Lng < end.Lng
			}
			return end.Lng < point.Lng && point.Lng <= start.Lng
		}
		if dyl > 0 {
			return start.Lat <= point.Lat && point.Lat < end.Lat
		}
		return end.Lat < point.Lat && point.Lat <= start.Lat
	default:
		if math.Abs(dxl) >= math.Abs(dyl) {
			if dxl > 0 {
				return start.Lng < point.Lng && point.Lng < end.Lng
			}
			return end.Lng < point.Lng && point.Lng < start.Lng
		}
		if dyl > 0 {
			return end.Lat < point.Lat && point.Lat < end.Lat
		}
		return end.Lat < point.Lat && point.Lat < end.Lat
	}
}
