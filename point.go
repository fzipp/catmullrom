// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package catmullrom

import (
	"math"
	"strconv"
)

// A Point represents a vector with coordinates X and Y in 2-dimensional
// euclidean space.
type Point struct {
	X, Y float64
}

// add returns the vector p+q.
func (p Point) add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// sub returns the vector p-q.
func (p Point) sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// mul returns the vector p*s.
func (p Point) mul(s float64) Point {
	return Point{p.X * s, p.Y * s}
}

// dot returns the dot (a.k.a. scalar) product of p and q.
func (p Point) dot(q Point) float64 {
	return p.X*q.X + p.Y*q.Y
}

// dist returns the euclidean distance between two vectors.
func (p Point) dist(q Point) float64 {
	return p.sub(q).len()
}

// sqLen returns the square of the length (euclidean norm) of a vector.
func (p Point) sqLen() float64 {
	return p.dot(p)
}

// len returns the length (euclidean norm) of a vector.
func (p Point) len() float64 {
	return math.Sqrt(p.sqLen())
}

// nearEq returns whether p and q are approximately equal. This relation is not
// transitive in general. The tolerance for the floating-point components is
// ±1e-5.
func (p Point) nearEq(q Point) bool {
	return nearEq(p.X, q.X, epsilon) && nearEq(p.Y, q.Y, epsilon)
}

// String returns a string representation of p like "(3.25, -1.5)".
func (p Point) String() string {
	return "(" + str(p.X) + ", " + str(p.Y) + ")"
}

const epsilon = 1e-10

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
