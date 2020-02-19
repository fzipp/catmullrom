// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package catmullrom

import "testing"

func TestPointString(t *testing.T) {
	tests := []struct {
		p    Point
		want string
	}{
		{Point{X: -2.3, Y: 1.1}, "(-2.3, 1.1)"},
		{Point{X: 2, Y: 1}, "(2, 1)"},
		{Point{X: 0.5, Y: 2}, "(0.5, 2)"},
		{Point{X: 1.414213, Y: -34.0213}, "(1.414213, -34.0213)"},
	}
	for _, tt := range tests {
		if s := tt.p.String(); s != tt.want {
			t.Errorf("(%g, %g).String() = %q, want %q", tt.p.X, tt.p.Y, s, tt.want)
		}
	}
}

func TestPointNearEq(t *testing.T) {
	tests := []struct {
		p, q Point
		want bool
	}{
		{Point{X: 4, Y: 1}, Point{X: 4, Y: 1}, true},
		{Point{X: -3.2145, Y: -2.5667}, Point{X: -3.2145, Y: -2.5667}, true},
		{Point{X: 2.3456789012, Y: -9.8765432109}, Point{X: 2.3456789012345678, Y: -9.8765432109876543}, true},
		{Point{X: 4, Y: 1}, Point{X: -3, Y: 7}, false},
		{Point{X: 2.34567, Y: -9.87654}, Point{X: 2.34567, Y: -9.87653}, false},
		{Point{X: 2.34567, Y: -9.87654}, Point{X: 2.34568, Y: -9.87654}, false},
	}
	for _, tt := range tests {
		if x := tt.p.nearEq(tt.q); x != tt.want {
			t.Errorf("%s.nearEq(%s) = %v, want %v", tt.p, tt.q, x, tt.want)
		}
	}
}

func TestPointAdd(t *testing.T) {
	tests := []struct {
		p, q Point
		want Point
	}{
		{Point{X: 4, Y: 1}, Point{X: 2, Y: 5}, Point{X: 6, Y: 6}},
		{Point{X: 1.2, Y: 2.3}, Point{X: -2.1, Y: 0.5}, Point{X: -0.9, Y: 2.8}},
		{Point{X: 12.5, Y: 9.25}, Point{}, Point{X: 12.5, Y: 9.25}},
		{Point{X: 1, Y: 0}, Point{X: 0, Y: 1}, Point{X: 1, Y: 1}},
	}
	for _, tt := range tests {
		if x := tt.p.add(tt.q); !x.nearEq(tt.want) {
			t.Errorf("%s + %s = %s, want %s", tt.p, tt.q, x, tt.want)
		}
	}
}

func TestPointSub(t *testing.T) {
	tests := []struct {
		p, q Point
		want Point
	}{
		{Point{X: 4, Y: 1}, Point{X: 2, Y: 5}, Point{X: 2, Y: -4}},
		{Point{X: 1.2, Y: 2.3}, Point{X: -2.1, Y: 0.5}, Point{X: 3.3, Y: 1.8}},
		{Point{X: 12.5, Y: 9.25}, Point{}, Point{X: 12.5, Y: 9.25}},
	}
	for _, tt := range tests {
		if x := tt.p.sub(tt.q); !x.nearEq(tt.want) {
			t.Errorf("%s - %s = %s, want %s", tt.p, tt.q, x, tt.want)
		}
	}
}

func TestPointMul(t *testing.T) {
	tests := []struct {
		p    Point
		s    float64
		want Point
	}{
		{Point{X: 4, Y: 1}, 2, Point{X: 8, Y: 2}},
		{Point{X: 1.4, Y: -2.5}, 0.5, Point{X: 0.7, Y: -1.25}},
		{Point{X: 12.5, Y: 9.25}, 1, Point{X: 12.5, Y: 9.25}},
		{Point{X: 2.7, Y: 1.1}, 0, Point{}},
	}
	for _, tt := range tests {
		if x := tt.p.mul(tt.s); !x.nearEq(tt.want) {
			t.Errorf("%g * %s = %s, want %s", tt.s, tt.p, x, tt.want)
		}
	}
}

func TestPointDot(t *testing.T) {
	tests := []struct {
		p, q Point
		want float64
	}{
		{Point{X: 2, Y: -3}, Point{X: -4, Y: 2}, -14},
		{Point{X: 4, Y: 8}, Point{X: 0.5, Y: 1.25}, 12},
		{Point{X: 12.5, Y: 9.25}, Point{}, 0},
		{Point{X: 1, Y: 0},  Point{X: 0, Y: 1}, 0},
		{Point{X: 4, Y: 5}, Point{X: 1, Y: 1}, 9},
	}
	for _, tt := range tests {
		if x := tt.p.dot(tt.q); x != tt.want {
			t.Errorf("%s.dot(%s) = %g, want %g", tt.p, tt.q, x, tt.want)
		}
	}
}

func TestPointDist(t *testing.T) {
	tests := []struct {
		p, q Point
		want float64
	}{
		{Point{}, Point{}, 0.0},
		{Point{}, Point{X: 1, Y: 0}, 1.0},
		{Point{X: 1, Y: 0}, Point{X: 0, Y: 1}, 1.4142135623730951},
		{Point{X: -2.3, Y: 1.1}, Point{X: -2.3, Y: 1.1}, 0.0},
		{Point{X: 2, Y: 1}, Point{X: 2, Y: 3}, 2.0},
		{Point{X: 0.5, Y: 2}, Point{X: 1.5, Y: 2.5}, 1.118033988749895},
	}
	for _, tt := range tests {
		if x := tt.p.dist(tt.q); !nearEq(x, tt.want, epsilon) {
			t.Errorf("%s.dist(%s) = %g, want %g", tt.p, tt.q, x, tt.want)
		}
	}
}

func TestPointLen(t *testing.T) {
	tests := []struct {
		p    Point
		want float64
	}{
		{Point{}, 0.0},
		{Point{X: 1, Y: 1}, 1.4142135623730951},
		{Point{X: 1, Y: 0}, 1.0},
		{Point{X: 0, Y: 1}, 1.0},
		{Point{X: -2.3, Y: 1.1}, 2.5495097567963922},
		{Point{X: 2, Y: 1}, 2.23606797749979},
		{Point{X: 0.5, Y: 2}, 2.0615528128088303},
	}
	for _, tt := range tests {
		if x := tt.p.len(); !nearEq(x, tt.want, epsilon) {
			t.Errorf("%s.len() = %g, want %g", tt.p, x, tt.want)
		}
	}
}
