package turfgo

import (
	"encoding/json"

	"github.com/kpawlik/geojson"
)

// DecodeLineStringFromFeatureJSON decode geojson feature type lineString into *LineString
func DecodeLineStringFromFeatureJSON(gj []byte) (*LineString, error) {
	var f *geojson.Feature
	err := json.Unmarshal(gj, &f)
	if err != nil {
		return nil, err
	}
	g, err := f.GetGeometry()
	if err != nil {
		return nil, err
	}
	ls, ok := g.(*geojson.LineString)
	if !ok {
		return nil, err
	}
	points := []*Point{}
	for _, c := range ls.Coordinates {
		points = append(points, decodePoint(c))
	}
	return NewLineString(points), nil
}

func decodePoint(coord geojson.Coordinate) *Point {
	return &Point{float64(coord[1]), float64(coord[0])}
}
