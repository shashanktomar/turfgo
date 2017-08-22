package turfgo

import (
	"github.com/kpawlik/geojson"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

// used to avoid compiler optimization
var testResultB bool

func TestOverlap(t *testing.T) {
	Convey("Should return false if boxes doesn't overlap", t, func() {
		// b2 above b1
		b1 := NewBBox(-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783)
		b2 := NewBBox(3.4716796874999996, 32.24997445586331, 8.876953125, 35.88905007936091)
		b, err := DoesBboxOverlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeFalse)

		// b2 on left of b1
		b1 = NewBBox(-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783)
		b2 = NewBBox(-12.2197265625, 28.24997445586331, -2.876953125, 39.88905007936091)
		b, err = DoesBboxOverlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeFalse)

		// b2 below b1
		b1 = NewBBox(-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783)
		b2 = NewBBox(-12.2197265625, 2.24997445586331, 23.876953125, 15.88905007936091)
		b, err = DoesBboxOverlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeFalse)

		// b2 on right of b1
		b1 = NewBBox(-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783)
		b2 = NewBBox(15.2197265625, 18.24997445586331, 23.876953125, 29.88905007936091)
		b, err = DoesBboxOverlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeFalse)
	})

	Convey("Should return true if boxes overlap", t, func() {
		// overlap where a point is inside either of the bbox
		b1 := NewBBox(-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783)
		b2 := NewBBox(-2.4716796874999996, 15.24997445586331, 8.876953125, 21.88905007936091)
		b, err := DoesBboxOverlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeTrue)

		// overlap where no point of either bbox reside inside any bbox
		b1 = NewBBox(-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783)
		b2 = NewBBox(-4.2197265625, 21.24997445586331, 26.876953125, 24.88905007936091)
		b, err = DoesBboxOverlap(b1, b2)
		So(err, ShouldBeNil)
		So(b, ShouldBeTrue)
	})

	Convey("Given nil bbox, should return error", t, func() {
		_, err := DoesBboxOverlap(nil, nil)
		So(err.Error(), ShouldEqual, "Bbox can't be nil")
	})
}

func BenchmarkIsBboxOverlap(b *testing.B) {
	b1 := NewBBox(-0.2197265625, 19.31114335506464, 13.447265624999998, 28.304380682962783)
	b2 := NewBBox(-4.2197265625, 21.24997445586331, 26.876953125, 24.88905007936091)
	for n := 0; n < b.N; n++ {
		testResultB, _ = DoesBboxOverlap(b1, b2)
	}
}

func TestIsPointOnLine(t *testing.T) {

	type testData struct {
		point        Point
		lineString   *LineString
		ignoreFields bool
	}

	parseTestFile := func(dir string)[]testData {
		res := make([]testData, 0)
		files, _ := ioutil.ReadDir(dir)
		for _, f := range files {
			j, _ := ioutil.ReadFile(dir + "/" + f.Name())
			fc, _ := DecodeFeatureCollection(j)

			ignoreEnds, _ := fc.Features[0].Properties["ignoreEndPoint"].(bool)
			g, _ := fc.Features[0].GetGeometry()
			gjPoint, ok := g.(*geojson.Point)
			So(ok, ShouldBeTrue)
			point := decodePoint(gjPoint.Coordinates)

			g, _ = fc.Features[1].GetGeometry()
			gjLine, ok := g.(*geojson.LineString)
			So(ok, ShouldBeTrue)
			line := decodeLine(gjLine)

			res = append(res, testData{*point, line, ignoreEnds})
		}
		return res
	}

	Convey("Given false cases, should pass", t, func() {
		testData := parseTestFile("./testdata/isPointOnLine/false")
		for _, t := range testData {
			So(IsPointOnLine(t.point, t.lineString, t.ignoreFields), ShouldBeFalse)
		}
	})

	Convey("Given true cases, should pass", t, func() {
		testData := parseTestFile("./testdata/isPointOnLine/true")
		for _, t := range testData {
			So(IsPointOnLine(t.point, t.lineString, t.ignoreFields), ShouldBeTrue)
		}
	})

}
