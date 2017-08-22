package turfgo

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"fmt"
	"github.com/kpawlik/geojson"
)

func TestDecodeLineString(t *testing.T) {
	Convey("Given geoJson feature file with linestring, should return lineString", t, func() {
		j, _ := ioutil.ReadFile("./testdata/geoJsonEncoder/linestringInFeature.geojson")
		ls, err := DecodeLineStringFromFeatureJSON(j)

		points := ls.Points

		So(err, ShouldBeNil)
		So(len(points), ShouldEqual, 3)
		So(points[0], ShouldResemble, &Point{22.466878364528448, -97.88131713867188})
		So(points[1], ShouldResemble, &Point{22.175960091218524, -97.82089233398438})
		So(points[2], ShouldResemble, &Point{21.8704201873689, -97.6190185546875})
	})

	Convey("Given invalid geoJson, should return error", t, func() {
		j, _ := ioutil.ReadFile("./testdata/geoJsonEncoder/invalidGeojson.geojson")
		ls, err := DecodeLineStringFromFeatureJSON(j)
		fmt.Println(err)
		So(ls, ShouldBeNil)
		So(err.Error(), ShouldEqual, "invalid character 'i' looking for beginning of value")
	})

	Convey("Given invalid geometry, should return error", t, func() {
		j, _ := ioutil.ReadFile("./testdata/geoJsonEncoder/invalidGeometry.geojson")
		ls, err := DecodeLineStringFromFeatureJSON(j)
		fmt.Println(err)
		So(ls, ShouldBeNil)
		So(err.Error(), ShouldEqual, "ParseError: Unknown geometry type InvalidGeometry")
	})

	Convey("Given geoJson does not have a linestring, should return error", t, func() {
		j, _ := ioutil.ReadFile("./testdata/geoJsonEncoder/featureCollection.geojson")
		ls, err := DecodeLineStringFromFeatureJSON(j)
		fmt.Println(err)
		So(ls, ShouldBeNil)
		So(err.Error(), ShouldEqual, "geometry is not of type linestring")
	})
}

func TestDecodeFeatureCollection(t *testing.T) {
	Convey("Given geoJson featureCollection file, should return a valid collection", t, func() {
		j, _ := ioutil.ReadFile("./testdata/geoJsonEncoder/featureCollection.geojson")
		f, err := DecodeFeatureCollection(j)
		So(err, ShouldBeNil)
		So(f.Type, ShouldEqual, "FeatureCollection")

		So(len(f.Features), ShouldEqual, 2)
		g1, err := f.Features[0].GetGeometry()
		So(err, ShouldBeNil)
		g2, err := f.Features[1].GetGeometry()
		So(err, ShouldBeNil)
		_, ok := g1.(*geojson.Point)
		So(ok, ShouldBeTrue)
		_, ok = g2.(*geojson.LineString)
		So(ok, ShouldBeTrue)
	})

	Convey("Given invalid geoJson, should return error", t, func() {
		j, _ := ioutil.ReadFile("./testdata/geoJsonEncoder/invalidGeojson.geojson")
		ls, err := DecodeFeatureCollection(j)
		fmt.Println(err)
		So(ls, ShouldBeNil)
		So(err.Error(), ShouldEqual, "invalid character 'i' looking for beginning of value")
	})
}
