package turfgo

const (
	infinity = 0x7FF0000000000000
)

// Unit constants
const (
	Kilometers    = "kilometers"
	Miles         = "miles"
	Meters        = "meters"
	CentiMeters   = "centimeters"
	MilliMeters   = "millimeters"
	Degrees       = "degress"
	Radians       = "radians"
	NauticalMiles = "nauticalmiles"
	Inches        = "inches"
	Yards         = "yards"
	Feet          = "feet"
	Acres         = "acres"
)

//R is radius of earth
var R = map[string]float64{
	Miles:      3960,
	Kilometers: 6373,
	Meters:     6373000,
	Degrees:    57.2957795,
	Radians:    1,
}

// Factors for unit conversion
var Factors = map[string]float64{
	Miles:         3960,
	NauticalMiles: 3441.145,
	Degrees:       57.2957795,
	Radians:       1,
	Inches:        250905600,
	Yards:         6969600,
	Meters:        6373000,
	CentiMeters:   6.373e+8,
	Kilometers:    6373,
	Feet:          20908792.65,
}

// AreaFactors for area unit conversion
var AreaFactors = map[string]float64{
	Kilometers:  0.000001,
	Meters:      1,
	CentiMeters: 10000,
	MilliMeters: 1000000,
	Acres:       0.000247105,
	Miles:       3.86e-7,
	Yards:       1.195990046,
	Feet:        10.763910417,
	Inches:      1550.003100006,
}

//Geometry is geoJson geometry
type Geometry interface {
	getPoints() []*Point
}

//PolygonI is geoJson polygon
type PolygonI interface {
	getPolygons() []*Polygon
}

//A Point on earth
type Point struct {
	Lat float64
	Lng float64
}

func (p *Point) getPoints() []*Point {
	return []*Point{p}
}

//NewPoint creates a new point for given lat, lng
func NewPoint(lat float64, lon float64) *Point {
	return &Point{lat, lon}
}

//MultiPoint geojson type
type MultiPoint struct {
	Points []*Point
}

func (p *MultiPoint) getPoints() []*Point {
	return p.Points
}

//NewMultiPoint creates a new multiPoint for given points
func NewMultiPoint(points []*Point) *MultiPoint {
	return &MultiPoint{Points: points}
}

//LineString geojson type
type LineString struct {
	Points []*Point
}

func (p *LineString) getPoints() []*Point {
	return p.Points
}

//NewLineString creates a new lineString for given points
func NewLineString(points []*Point) *LineString {
	return &LineString{Points: points}
}

//MultiLineString geojson type
type MultiLineString struct {
	LineStrings []*LineString
}

func (p *MultiLineString) getPoints() []*Point {
	points := []*Point{}
	for _, lineString := range p.LineStrings {
		points = append(points, lineString.getPoints()...)
	}
	return points
}

//NewMultiLineString creates a new multiLineString for given lineStrings
func NewMultiLineString(lineStrings []*LineString) *MultiLineString {
	return &MultiLineString{LineStrings: lineStrings}
}

//Polygon geojson type
type Polygon struct {
	LineStrings []*LineString
}

func (p *Polygon) getPoints() []*Point {
	points := []*Point{}
	for _, lineString := range p.LineStrings {
		points = append(points, lineString.getPoints()...)
	}
	return points
}

func (p *Polygon) getPolygons() []*Polygon {
	return []*Polygon{p}
}

//NewPolygon creates a new polygon for given lineStrings
func NewPolygon(lineStrings []*LineString) *Polygon {
	return &Polygon{LineStrings: lineStrings}
}

//MultiPolygon geojson type
type MultiPolygon struct {
	Polygons []*Polygon
}

func (p *MultiPolygon) getPoints() []*Point {
	points := []*Point{}
	for _, polygon := range p.Polygons {
		points = append(points, polygon.getPoints()...)
	}
	return points
}

func (p *MultiPolygon) getPolygons() []*Polygon {
	return p.Polygons
}

// NewMultiPolygon creates a new multiPolygon for given polygons
func NewMultiPolygon(polygons []*Polygon) *MultiPolygon {
	return &MultiPolygon{Polygons: polygons}
}

// BoundingBox represent a bbox
type BoundingBox struct {
	West  float64
	South float64
	East  float64
	North float64
}

// NewInfiniteBBox creates a bounding box with corners really far away
func NewInfiniteBBox() *BoundingBox {
	return &BoundingBox{infinity, infinity, -infinity, -infinity}
}

// NewBBox creates bounding box with given corners
func NewBBox(w float64, s float64, e float64, n float64) *BoundingBox {
	return &BoundingBox{w, s, e, n}
}
