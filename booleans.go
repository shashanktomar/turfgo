package turfgo

import "errors"

// IsBboxOverlap takes two bounding box and returns true if there is an overlap.
// The order of values in array is WSEN(west, south , east, north)
func IsBboxOverlap(b1 *BoundingBox, b2 *BoundingBox) (bool, error) {
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
