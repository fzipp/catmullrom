// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package catmullrom

import (
	"math"
	"strconv"
)

// A Point represents a vector with coordinates X and Y in 2-dimensional
// euclidian space.
type Point struct {
	X, Y float64
}

// add returns the vector v+w.
func (v Point) add(w Point) Point {
	return Point{v.X + w.X, v.Y + w.Y}
}

// sub returns the vector v-w.
func (v Point) sub(w Point) Point {
	return Point{v.X - w.X, v.Y - w.Y}
}

// mul returns the vector v*s.
func (v Point) mul(s float64) Point {
	return Point{v.X * s, v.Y * s}
}

// dot returns the dot (a.k.a. scalar) product of v and w.
func (v Point) dot(w Point) float64 {
	return v.X*w.X + v.Y*w.Y
}

// dist returns the euclidian distance between two vectors.
func (v Point) dist(w Point) float64 {
	return v.sub(w).len()
}

// sqLen returns the square of the length (euclidian norm) of a vector.
func (v Point) sqLen() float64 {
	return v.dot(v)
}

// len returns the length (euclidian norm) of a vector.
func (v Point) len() float64 {
	return math.Sqrt(v.sqLen())
}

// nearEq returns whether v and w are approximately equal. This relation is not
// transitive in general. The tolerance for the floating-point components is
// ±1e-5.
func (v Point) nearEq(w Point) bool {
	return nearEq(v.X, w.X, epsilon) && nearEq(v.Y, w.Y, epsilon)
}

// String returns a string representation of v like "{X: 3.25, Y: -1.5}".
func (v Point) String() string {
	return "{X: " + str(v.X) + ", Y: " + str(v.Y) + "}"
}

const epsilon = 1e-5

// nearEq compares two floating-point numbers for equality within an
// absolute difference tolerance of epsilon.
// This relation is not transitive, except for epsilon=0.
func nearEq(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

// str converts a float64 to a string in "%g" format.
func str(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}
