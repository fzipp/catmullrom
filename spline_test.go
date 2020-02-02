// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package catmullrom

import (
	"testing"
)

type chainTest struct {
	testName         string
	controlPoints    []Point
	pointsPerSegment int
	alpha            float64
	want             []Point
}

var chainSingleSegmentTest = chainTest{
	testName: "Catmull-Rom chain with single segment",
	controlPoints: []Point{
		{X: 0, Y: 0},
		{X: 1, Y: 3},
		{X: 2, Y: 3},
		{X: 3, Y: 0},
	},
	pointsPerSegment: 5,
	alpha:            0.5,
	want: []Point{
		{X: 1, Y: 3},
		{X: 1.2352316866996362, Y: 3.113853557808694},
		{X: 1.5, Y: 3.151804743744926},
		{X: 1.764768313300364, Y: 3.113853557808694},
		{X: 2, Y: 3},
	},
}

var chainTwoSegmentsTest = chainTest{
	testName: "Catmull-Rom chain with two segments",
	controlPoints: []Point{
		{X: 0, Y: 0},
		{X: 1, Y: 1},
		{X: 2, Y: 3},
		{X: 3, Y: 1},
		{X: 4, Y: 0},
	},
	pointsPerSegment: 5,
	alpha:            0.5,
	want: []Point{
		{X: 1, Y: 1},
		{X: 1.2701649675834827, Y: 1.5355841633506426},
		{X: 1.5179244156297622, Y: 2.19829703408946},
		{X: 1.7567216558611611, Y: 2.7618613877835476},
		{X: 2, Y: 3},
		{X: 2.243278344138839, Y: 2.761861387783547},
		{X: 2.4820755843702376, Y: 2.1982970340894603},
		{X: 2.729835032416517, Y: 1.5355841633506424},
		{X: 3, Y: 1},
	},
}

func TestChain(t *testing.T) {
	tests := []chainTest{
		chainSingleSegmentTest,
		chainTwoSegmentsTest,
	}
	for _, tt := range tests {
		if spline := chain(tt.controlPoints, tt.pointsPerSegment, tt.alpha); !pointsNearEqual(spline, tt.want) {
			t.Errorf("%s:\ncontrol points: %v\npoints per segment: %d\nalpha: %g\nexpected: %v\nactual: %v",
				tt.testName, tt.controlPoints, tt.pointsPerSegment, tt.alpha, tt.want, spline)
		}
	}
}

func pointsNearEqual(a, b []Point) bool {
	if len(a) != len(b) {
		return false
	}
	for i, p := range a {
		if !p.nearEq(b[i]) {
			return false
		}
	}
	return true
}
