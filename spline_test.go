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
		{X: 1.23523, Y: 3.11385},
		{X: 1.5, Y: 3.15180},
		{X: 1.76477, Y: 3.11385},
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
		{X: 1.27017, Y: 1.53558},
		{X: 1.51792, Y: 2.1983},
		{X: 1.75672, Y: 2.76186},
		{X: 2, Y: 3},
		{X: 2.24328, Y: 2.76186},
		{X: 2.48208, Y: 2.1983},
		{X: 2.72984, Y: 1.53558},
		{X: 3, Y: 1},
	},
}

func TestChain(t *testing.T) {
	tests := []chainTest{
		chainSingleSegmentTest,
		chainTwoSegmentsTest,
	}
	for _, tt := range tests {
		if spline := Chain(tt.controlPoints, tt.pointsPerSegment, tt.alpha); !pointsNearEqual(spline, tt.want) {
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
