package turfgo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
)

var units = [4]Unit{Kilometers, Miles, Degrees, Radians}

// used to avoid compiler optimization
var testResultF float64
var testResultP *Point
var testResultBbox *BoundingBox

var longRoute *LineString

func init() {
	gj, _ := ioutil.ReadFile("./testdata/common/route.geojson")
	longRoute, _ = DecodeLineStringFromFeatureJSON(gj)
}

func TestAlong(t *testing.T) {
	type alongTest struct {
		distance float64
		unit     Unit
		result   *Point
	}

	testValues := []alongTest{
		{1, Miles, NewPoint(38.88533657311743, -77.02417489836314)},
		{1.2, Miles, NewPoint(38.8871105586916, -77.02436062207721)},
		{1.4, Miles, NewPoint(38.88938593771034, -77.0220637504277)},
		{1.6, Miles, NewPoint(38.891879938286934, -77.02018399074201)},
		{1.8, Miles, NewPoint(38.893500737015884, -77.0224424741873)},
		{2, Miles, NewPoint(38.89617811276868, -77.02291488647461)},
		{100, Miles, NewPoint(38.931505469602044, -77.03596115112305)},
		{0, Miles, NewPoint(38.878605901789236, -77.0316696166992)},
	}

	Convey("Should return a point along distance", t, func() {
		gj, err := ioutil.ReadFile("./testdata/along/line.geojson")
		So(err, ShouldBeNil)
		ls, err := DecodeLineStringFromFeatureJSON(gj)
		So(err, ShouldBeNil)
		for _, tt := range testValues {
			p := Along(ls, tt.distance, tt.unit)
			So(p.Lat, ShouldAlmostEqual, tt.result.Lat, 0.0000001)
			So(p.Lng, ShouldAlmostEqual, tt.result.Lng, 0.0000001)
		}

	})
}

func BenchmarkAlong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testResultP = Along(longRoute, 20.234, Miles)
	}
}

func TestBearing(t *testing.T) {

	type bearingTest struct {
		point1 *Point
		point2 *Point
		result float64
	}

	testValues := []bearingTest{
		{NewPoint(39.984, -75.343),
			NewPoint(39.123, -75.534),
			-170.23304913492177},
		{NewPoint(12.9715987, 77.59456269999998),
			NewPoint(13.22328378, 77.77448784),
			34.828578946361255},
	}

	Convey("Given two points, should calculate bearing between them", t, func() {
		for _, tt := range testValues {
			actual := Bearing(tt.point1, tt.point2)
			So(actual, ShouldAlmostEqual, tt.result, 0.0000001)
		}
	})
}

func BenchmarkBearing(b *testing.B) {
	b.StopTimer()
	p1 := NewPoint(39.984, -75.343)
	p2 := NewPoint(39.123, -75.534)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		testResultF = Bearing(p1, p2)
	}
}

func TestDestination(t *testing.T) {

	type destinationTest struct {
		point    *Point
		distance float64
		bearing  float64
		result   map[Unit]Point
	}

	testValues := []destinationTest{
		{NewPoint(38.10096062273525, -75), 100, 0,
			map[Unit]Point{
				Kilometers: {39, -75},
				Miles:      {39.54782374175248, -75},
				Degrees:    {41.8990393544318, 105},
				Radians:    {7.678911930967332, -75},
			},
		},
		{NewPoint(39, -75), 100, 180,
			map[Unit]Point{
				Kilometers: {38.10096062273525, -75},
				Miles:      {37.55313688098277, -75},
				Degrees:    {-61.00000002283296, -74.99999999999999},
				Radians:    {69.42204869176791, -75},
			},
		},
		{NewPoint(39, -75), 100, 90,
			map[Unit]Point{
				Kilometers: {38.994288534328966, -73.84321473156825},
				Miles:      {38.985208813672266, -73.13849445143401},
				Degrees:    {-6.27383195845071, 22.802746801915237},
				Radians:    {32.86591377972705, -112.07480823869463},
			},
		},
		{NewPoint(39, -75), 5000, 90,
			map[Unit]Point{
				Kilometers: {26.446988157260996, -22.898974671086123},
				Miles:      {11.00429485821584, 1.1054470055309658},
				Degrees:    {28.821822144704377, -122.19517685125443},
				Radians:    {5.58578497583862, -158.06325963430967},
			},
		},
	}

	Convey("Should return correct destination", t, func() {
		for _, tt := range testValues {
			for _, unit := range units {
				expected := tt.result[unit]
				dest := Destination(tt.point, tt.distance, tt.bearing, unit)
				So(dest.Lat, ShouldAlmostEqual, expected.Lat, 0.0000001)
				So(dest.Lng, ShouldAlmostEqual, expected.Lng, 0.0000001)
			}
		}
	})

}

func BenchmarkDestination(b *testing.B) {
	b.StopTimer()
	p := NewPoint(39.984, -75.343)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		testResultP = Destination(p, 45.34, 120.5, Miles)
	}
}

func TestDistance(t *testing.T) {

	type distanceTest struct {
		point1 *Point
		point2 *Point
		result map[Unit]float64
	}

	testValues := []distanceTest{
		{NewPoint(39.984, -75.343),
			NewPoint(39.123, -75.534),
			map[Unit]float64{
				Kilometers: 97.15957803131901,
				Miles:      60.37218405837491,
				Degrees:    0.8735028650863799,
				Radians:    0.015245501024842149,
			},
		},
		{NewPoint(72.134, -10.143),
			NewPoint(39.123, -75.534),
			map[Unit]float64{
				Kilometers: 5072.014768708954,
				Miles:      3151.604971612656,
				Degrees:    45.59940998096528,
				Radians:    0.7958598413163273,
			},
		},
	}

	Convey("Given two points, should calculate distance between them", t, func() {
		for _, tt := range testValues {
			for _, unit := range units {
				actual := Distance(tt.point1, tt.point2, unit)
				So(actual, ShouldAlmostEqual, tt.result[unit], 0.0000001)
			}
		}
	})

}

func BenchmarkDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testResultF = Distance(&Point{39.984, -75.343},
			&Point{39.123, -75.534}, Miles)
	}
}

func TestExtent(t *testing.T) {

	type extentTest struct {
		geometry Geometry
		result   *BoundingBox
	}

	point := NewPoint(0.5, 102.0)
	lineString := NewLineString([]*Point{
		{-10.0, 102.0},
		{1.0, 103.0},
		{0.0, 104.0},
		{4.0, 130.0},
	})
	polygon := NewPolygon([]*LineString{NewLineString([]*Point{
		{0.0, 101},
		{1.0, 101.0},
		{1.0, 100.0},
		{0.0, 100.0},
		{0.0, 101.0},
	})})
	multiLineString := NewMultiLineString([]*LineString{
		{[]*Point{{0, 100}, {1, 101}}},
		{[]*Point{{2, 102}, {3, 103}}},
	})
	multiPoly := NewMultiPolygon([]*Polygon{
		{[]*LineString{
			{[]*Point{
				{2, 102},
				{2, 103},
				{3, 103},
				{3, 102},
				{2, 102},
			}},
		}},
		{[]*LineString{
			{[]*Point{
				{0, 100},
				{0, 101},
				{1, 101},
				{1, 100},
				{0, 100},
			}},
			{[]*Point{
				{0.2, 100.2},
				{0.2, 100.8},
				{0.8, 100.8},
				{0.8, 100.2},
				{0.2, 100.2},
			}},
		}},
	})

	testValues := []extentTest{
		{point, NewBBox(102, 0.5, 102, 0.5)},
		{lineString, NewBBox(102, -10, 130, 4)},
		{polygon, NewBBox(100, 0, 101, 1)},
		{multiLineString, NewBBox(100, 0, 103, 3)},
		{multiPoly, NewBBox(100, 0, 103, 3)},
	}

	Convey("Given different type of shapes, should return bounding box", t, func() {
		for _, tt := range testValues {
			bBox := Extent(tt.geometry)
			So(bBox, ShouldResemble, tt.result)
		}

		bBox := Extent(testValues[0].geometry, testValues[1].geometry,
			testValues[2].geometry, testValues[3].geometry, testValues[4].geometry)
		So(bBox, ShouldResemble, NewBBox(100, -10, 130, 4))
	})
}

func BenchmarkExtent(b *testing.B) {
	b.StopTimer()
	polygon := NewPolygon([]*LineString{NewLineString([]*Point{
		{0.0, 101},
		{1.0, 101.0},
		{1.0, 100.0},
		{0.0, 100.0},
		{0.0, 101.0},
	})})
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		testResultBbox = Extent(polygon)
	}
}

func TestCenter(t *testing.T) {

	Convey("Given an array of points, should return absolute center of points", t, func() {
		point1 := &Point{35.4691, -97.522259}
		point2 := &Point{35.463455, -97.502754}
		point3 := &Point{35.463245, -97.508269}
		point4 := &Point{35.465779, -97.516809}
		point5 := &Point{35.467072, -97.515372}
		lineString := NewLineString([]*Point{point1, point2, point3, point4})

		point := Center(lineString, point5)
		So(point.Lat, ShouldEqual, 35.4661725)
		So(point.Lng, ShouldEqual, -97.5125065)
	})
}

func TestExpand(t *testing.T) {
	type expandTest struct {
		geometry Geometry
		distance float64
		unit     Unit
		result   *BoundingBox
	}

	point := NewPoint(35.4691, -97.522259)
	lineString := NewLineString([]*Point{
		{35.964669147704086, -96.96258544921875},
		{35.87792352995116, -97.39654541015625},
		{35.66622234103479, -97.6409912109375},
		{35.561277754384555, -97.22351074218749},
		{35.45619556834375, -97.54486083984375},
	})

	testValues := []expandTest{
		{point, 20, Kilometers, NewBBox(-97.74303658676054, 35.28929212454705, -97.30148141323949, 35.64890787545294)},
		{lineString, 10, Kilometers, NewBBox(-97.75136243406826, 35.36629163061727, -96.85150786206225, 36.05457308543056)},
	}
	Convey("Given different type of shapes, should return bounding box", t, func() {
		for _, tt := range testValues {
			b := Expand(tt.distance, tt.unit, tt.geometry)
			So(b, ShouldResemble, tt.result)
		}
	})

}

func BenchmarkExpand(b *testing.B) {
	b.StopTimer()
	lineString := NewLineString([]*Point{
		{35.964669147704086, -96.96258544921875},
		{35.87792352995116, -97.39654541015625},
		{35.66622234103479, -97.6409912109375},
		{35.561277754384555, -97.22351074218749},
		{35.45619556834375, -97.54486083984375},
	})
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		testResultBbox = Expand(20, Kilometers, lineString)
	}
}
