package turfgo

import "github.com/twpayne/gopolyline/polyline"

//EncodePolyline encodes given coordinates into a polyline for the given dimension
func EncodePolyline(coordinates []*Point, dim int) string {
	var flatC []float64
	for i := 0; i < len(coordinates); i++ {
		flatC = append(flatC, coordinates[i].Lat)
		flatC = append(flatC, coordinates[i].Lng)
	}

	return polyline.Encode(flatC, dim)
}

//DecodePolyline decodes given polyline for given dimension and return coordinates
func DecodePolyline(line string, dim int) ([]*Point, error) {
	flatC, err := polyline.Decode(line, dim)
	if err != nil {
		return nil, err
	}
	var coordinates []*Point
	for i := 0; i < len(flatC)/dim; i++ {
		point := &Point{Lat: flatC[dim*i], Lng: flatC[dim*i+1]}
		coordinates = append(coordinates, point)
	}
	return coordinates, nil
}
