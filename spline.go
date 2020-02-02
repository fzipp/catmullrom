// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package catmullrom provides an implementation of the centripetal
// Catmull-Rom spline. It calculates points of a 2D spline curve given a series
// of control points.
package catmullrom

import (
	"math"
)

// SplineChain creates a spline curve through a series of control points by
// chaining Catmull-Rom splines.
//
// Each curve segment for two control points will have pointsPerSegment
// points, including the control points. Control points are not duplicated
// at the seams of the segments.
// Each control point is included in the result exactly as given.
//
// The alpha value ranges from 0 to 1. An alpha value of 0.5 results in a
// centripetal spline, alpha=0 results in a uniform spline, and alpha=1 results
// in a chordal spline.
func SplineChain(controlPoints []Point, pointsPerSegment int, alpha float64) []Point {
	P := make([]Point, len(controlPoints)+2)
	copy(P[1:], controlPoints)

	cp0 := controlPoints[0]
	cp1 := controlPoints[1]
	// Additional extrapolated control point at the beginning
	P[0] = cp0.sub(cp1.sub(cp0))

	cpy := controlPoints[len(controlPoints)-2]
	cpz := controlPoints[len(controlPoints)-1]
	// Additional extrapolated control point at the end
	P[len(P)-1] = cpz.add(cpz.sub(cpy))

	return chain(P, pointsPerSegment, alpha)
}

func chain(controlPoints []Point, pointsPerSegment int, alpha float64) []Point {
	P := controlPoints
	nSegments := len(P) - 3
	curve := make([]Point, 0, nSegments*pointsPerSegment-(nSegments-1))
	for i := 0; i < nSegments; i++ {
		segment := Spline(P[i], P[i+1], P[i+2], P[i+3], pointsPerSegment, alpha)
		if i == 0 {
			curve = append(curve, segment...)
		} else {
			// Do not duplicate points at seams.
			curve = append(curve, segment[1:]...)
		}
	}
	return curve
}

// Spline calculates the Catmull-Rom spline curve defined by the points p0, p1,
// p2, p3. The resulting curve starts with p1 and ends with p2. Both these
// points are included exactly as given. The total number of points for the
// resulting curve is defined by nPoints.
//
// The alpha value ranges from 0 to 1. An alpha value of 0.5 results in a
// centripetal spline, alpha=0 results in a uniform spline, and alpha=1 results
// in a chordal spline.
func Spline(p0, p1, p2, p3 Point, nPoints int, alpha float64) []Point {

	tj := func(ti float64, pi, pj Point) float64 {
		return math.Pow(pj.dist(pi), alpha) + ti
	}

	t0 := float64(0)
	t1 := tj(t0, p0, p1)
	t2 := tj(t1, p1, p2)
	t3 := tj(t2, p2, p3)

	step := (t2 - t1) / float64(nPoints-1)

	spline := make([]Point, nPoints)
	spline[0] = p1
	for i := 1; i < nPoints-1; i++ {

		t := t1 + (float64(i) * step)

		a1 := p0.mul((t1 - t) / (t1 - t0)).add(p1.mul((t - t0) / (t1 - t0)))
		a2 := p1.mul((t2 - t) / (t2 - t1)).add(p2.mul((t - t1) / (t2 - t1)))
		a3 := p2.mul((t3 - t) / (t3 - t2)).add(p3.mul((t - t2) / (t3 - t2)))

		b1 := a1.mul((t2 - t) / (t2 - t0)).add(a2.mul((t - t0) / (t2 - t0)))
		b2 := a2.mul((t3 - t) / (t3 - t1)).add(a3.mul((t - t1) / (t3 - t1)))

		c := b1.mul((t2 - t) / (t2 - t1)).add(b2.mul((t - t1) / (t2 - t1)))
		spline[i] = c
	}
	spline[nPoints-1] = p2
	return spline
}
